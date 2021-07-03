package core

import (
	"context"
	"log"
	"os"

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
			log.Println(r)
		}
	}
}
