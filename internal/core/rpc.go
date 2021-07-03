package core

import "github.com/gin-gonic/gin"

func (c *core) router() {
	c.app.GET("/info", c.info)
	c.app.GET("/query", c.query)
	c.app.POST("/transfer", c.transfer)
}

// info 获取基础信息
func (c *core) info(ctx *gin.Context) {

}

// query 基础查询
func (c *core) query(ctx *gin.Context) {

}

// transfer 转账
func (c *core) transfer(ctx *gin.Context) {

}
