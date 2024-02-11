package main

import "goexpert.com/module07/configs"

func main()  {
	config , _ := configs.LoadConfig(".")
	println(config.DBDriver)
}