package main

import (
	"awesomeProject/api"
	"awesomeProject/dao"
)

func main() {
	dao.InitDb()
	api.Init()

}
