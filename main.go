package main

import (
	"kaliwe.ru/menu/config"
)

func main() {
	router := config.Init()
	router.Run()
}
