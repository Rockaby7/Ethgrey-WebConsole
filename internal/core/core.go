package core

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/Pow-Duck/ethgrey/internal/storage"
	"github.com/Pow-Duck/ethgrey/pkg"
	"github.com/Pow-Duck/ethgrey/utils"
	"github.com/gin-gonic/gin"
)

type core struct {
	app    *gin.Engine
	cancel context.CancelFunc
}

func New() *core {
	return &core{}
}

func (c *core) Run() error {
	c.app = gin.Default()
	c.router()

	addr := os.Getenv("ADDR")
	if addr == "" {
		addr = "0.0.0.0:8097"
	}

	go c.core()
	return c.app.Run(addr)
}

func (c *core) core() {
	ctx, cancel := context.WithCancel(context.Background())
	outChan := make(chan string, 100)
	err := utils.CommandWriteBack(ctx, "java -jar grey.jar", outChan)
	if err != nil {
		log.Fatalln(err)
		return
	}

	c.cancel = cancel
loop:
	for {
		select {
		case r, ex := <-outChan:
			if !ex {
				break loop
			}
			// 处理
			c.parse(r)
		}
	}

	fmt.Println("Over")
}

func (c *core) parse(t string) {
	switch {
	case strings.Contains(t, "Welcome to Grey"):
		err := storage.Storage.SetNX(pkg.GeryVersion, []byte(strings.TrimSpace(t[strings.Index(t, "Grey"):])), 0)
		if err != nil {
			log.Println(err)
		}
	case strings.Contains(t, "Grey Test Network"):
		switch strings.Contains(t, "Successful") {
		case true:
			err := storage.Storage.SetNX(pkg.NetWork, []byte("true"), 0)
			if err != nil {
				log.Println(err)
			}
		case false:

		}
	case strings.Contains(t, "Sorry"):
		err := storage.Storage.SetNX(pkg.Error, []byte(strings.TrimSpace(t)), 0)
		if err != nil {
			log.Println(err)
		}
	case strings.Contains(t, "Address activated successfully"):
		err := storage.Storage.SetNX(pkg.NetWorkSuccess, []byte("true"), 0)
		if err != nil {
			log.Println(err)
		}
	case strings.Contains(t, "Number Of Nodes"):
		err := storage.Storage.SetNX(pkg.NumberOfNodes, []byte(strings.TrimSpace(t[strings.Index(t, ":")+1:])), 0)
		if err != nil {
			log.Println(err)
		}
	case strings.Contains(t, "GRB Output"):
		err := storage.Storage.SetNX(pkg.GRB, []byte(strings.TrimSpace(t[strings.Index(t, ":")+1:])), 0)
		if err != nil {
			log.Println(err)
		}
	case strings.Contains(t, "Harvest"):
		err := storage.Storage.SetNX(pkg.Harvest, []byte(strings.TrimSpace(t[strings.Index(t, ":")+1:])), 0)
		if err != nil {
			log.Println(err)
		}
	case strings.Contains(t, "Ranking:"):
		err := storage.Storage.SetNX(pkg.Ranking, []byte(strings.TrimSpace(t[strings.Index(t, ":")+1:])), 0)
		if err != nil {
			log.Println(err)
		}
	case strings.Contains(t, "Ethereum Address"):
		err := storage.Storage.SetNX(pkg.EthereumAddress, []byte(strings.TrimSpace(t[strings.Index(t, ":")+1:])), 0)
		if err != nil {
			log.Println(err)
		}
	case strings.Contains(t, "Private Key"):
		err := storage.Storage.SetNX(pkg.PrivateKey, []byte(strings.TrimSpace(t[strings.Index(t, ":")+1:])), 0)
		if err != nil {
			log.Println(err)
		}
	case strings.Contains(t, "Height confirmation of block"):
		err := storage.Storage.SetNX(pkg.HeightBlock, []byte(strings.TrimSpace(t[strings.Index(t, "block")+5:])), 0)
		if err != nil {
			log.Println(err)
		}
	}
}
