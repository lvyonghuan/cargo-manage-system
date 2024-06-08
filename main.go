package main

import (
	"cargo-manage-system/api"
	"cargo-manage-system/config"
	"cargo-manage-system/dao"
)

func main() {
	cfg := config.ReadConfig()
	dao.InitDB(cfg.Database.Username, cfg.Database.Password)
	api.InitRouter()
}
