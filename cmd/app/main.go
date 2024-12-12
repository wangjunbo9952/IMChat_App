package main

import (
	"IMChat_App/internal/dao/pool"
	"IMChat_App/internal/router"
	"IMChat_App/internal/websocket"
	"fmt"
)

func main() {
	err := pool.InitDB()
	if err != nil {
		fmt.Println("mysql connect failed!")
		return
	}

	go websocket.MyServer.Start()

	router.NewRouter()
	// fmt.Println("mysql connect success")

}
