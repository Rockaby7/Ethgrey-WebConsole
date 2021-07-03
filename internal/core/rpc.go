package core

import (
	"github.com/Pow-Duck/ethgrey/internal/storage"
	"github.com/Pow-Duck/ethgrey/pkg"
	"github.com/Pow-Duck/ethgrey/utils"
	"github.com/gin-gonic/gin"

	"log"
	"strings"
	"sync"
)

func (c *core) router() {
	c.app.GET("/info", c.info)
	c.app.GET("/query", c.query)
	c.app.GET("/reboot", c.reboot)
	c.app.POST("/transfer", c.transfer)
	c.app.POST("/update_ethgrey", c.updateEthgrey)
}

type InfoOutput struct {
	GeryVersion     string `json:"gery_version"`
	NetWork         bool   `json:"net_work"`
	Error           string `json:"error"`
	NetWorkSuccess  bool   `json:"net_work_success"`
	NumberOfNodes   string `json:"number_of_nodes"`
	GRBOutput       string `json:"grb_output"`
	Harvest         string `json:"harvest"`
	Ranking         string `json:"ranking"`
	EthereumAddress string `json:"ethereum_address"`
	PrivateKey      string `json:"private_key"`
	HeightBlock     string `json:"height_block"`
}

// info 获取基础信息
func (c *core) info(ctx *gin.Context) {
	var result InfoOutput

	version, err := storage.Storage.Get(pkg.GeryVersion)
	if err == nil {
		result.GeryVersion = string(version)
	}

	_, err = storage.Storage.Get(pkg.NetWork)
	if err == nil {
		result.NetWork = true
	}

	es, err := storage.Storage.Get(pkg.Error)
	if err == nil {
		result.Error = string(es)
	}

	_, err = storage.Storage.Get(pkg.NetWorkSuccess)
	if err == nil {
		result.NetWorkSuccess = true
	}

	nON, err := storage.Storage.Get(pkg.NumberOfNodes)
	if err == nil {
		result.NumberOfNodes = string(nON)
	}

	grb, err := storage.Storage.Get(pkg.GRB)
	if err == nil {
		result.GRBOutput = string(grb)
	}

	harvest, err := storage.Storage.Get(pkg.Harvest)
	if err == nil {
		result.Harvest = string(harvest)
	}

	ranking, err := storage.Storage.Get(pkg.Ranking)
	if err == nil {
		result.Ranking = string(ranking)
	}

	ethereumAddress, err := storage.Storage.Get(pkg.EthereumAddress)
	if err == nil {
		result.EthereumAddress = string(ethereumAddress)
	}

	privateKey, err := storage.Storage.Get(pkg.PrivateKey)
	if err == nil {
		result.PrivateKey = string(privateKey)
	}

	heightBlock, err := storage.Storage.Get(pkg.HeightBlock)
	if err == nil {
		result.HeightBlock = string(heightBlock)
	}

	ctx.JSON(200, pkg.Request{Data: result})
}

var mu sync.Mutex

// reboot 重启
func (c *core) reboot(ctx *gin.Context) {
	mu.Lock()
	defer mu.Unlock()

	c.cancel()

	go c.core()

	ctx.JSON(200, pkg.Request{
		Msg: "success",
	})
}

// updateEthgrey 更新
func (c *core) updateEthgrey(ctx *gin.Context) {
	mu.Lock()
	defer mu.Unlock()

	c.cancel()

	err := utils.UpdateEthgreyCore()
	if err != nil {
		ctx.JSON(500, pkg.Request{
			Msg:       err.Error(),
			ErrorCode: 500,
		})
		return
	}

	go c.core()

	ctx.JSON(200, pkg.Request{
		Msg: "success",
	})
}

// query 基础查询
func (c *core) query(ctx *gin.Context) {
	query, err := utils.Query()
	if err != nil {
		ctx.JSON(500, pkg.Request{
			ErrorCode: 500,
			Msg:       err.Error(),
		})
		return
	}

	var respList []string
	for i, v := range query {
		if i >= 2 && i < len(query)-2 {
			respList = append(respList, v)
		}
	}

	ctx.JSON(200, pkg.Request{Data: query})
}

type TransferInput struct {
	TransferAddress string `json:"transfer_address"`
	Amount          int    `json:"amount"`
}

// transfer 转账
func (c *core) transfer(ctx *gin.Context) {
	var tr TransferInput
	err := ctx.BindJSON(&tr)
	if err != nil {
		log.Fatalln("傻叉 日你二大爷")
		return
	}

	transfer, err := utils.Transfer(tr.TransferAddress, tr.Amount)
	if err != nil {
		ctx.JSON(500, pkg.Request{
			ErrorCode: 500,
			Msg:       err.Error(),
		})
		return
	}

	var errStr string
	switch {
	case strings.Contains(transfer, "successful"):
	case strings.Contains(transfer, "Sorry"):
		errStr = transfer[strings.Index(transfer, "Sorry"):strings.Index(transfer, "Do you want to")]
	}

	if errStr != "" {
		ctx.JSON(400, pkg.Request{
			ErrorCode: 500,
			Msg:       errStr,
		})
		return
	}

	ctx.JSON(200, pkg.Request{
		Msg: "success",
	})
}
