package main

import (
	"log"

	"github.com/Pow-Duck/ethgrey/internal/core"
	"github.com/Pow-Duck/ethgrey/utils"
)

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	go utils.SignalFn()
	//utils.Command(ctx,`echo -e "query\rexit\n"|java -cp grey.jar MainApp`)

	//command, err := utils.Command(`echo -e "query\rexit\n"|java -cp grey.jar MainApp`)
	//if err != nil {
	//	log.Fatalln(err)
	//	return
	//}
	//
	//fmt.Println(command)

	//time.Sleep(time.Second * 66666)
	//
	//getenv := os.Getenv("ADDRESS")

	c := core.New()
	if err := c.Run(); err != nil {
		log.Fatalln(err)
	}
}
