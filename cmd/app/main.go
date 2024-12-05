package main

import (
	"IMChat_App/internal/dao/pool"
	"IMChat_App/internal/router"
	"fmt"
)

func main() {
	err := pool.InitDB()
	if err != nil {
		fmt.Println("mysql connect failed!")
		return
	}
	router.NewRouter()
	// fmt.Println("mysql connect success")

}
