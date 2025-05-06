package main

import (
	"scoreboard/config"
	"fmt"
)

func main() {
	config.InitConfig()
	fmt.Printf("DB_USER: %v\n", config.DB_USER)
	fmt.Printf("DB_USER: %v\n", config.DB_NAME)

}