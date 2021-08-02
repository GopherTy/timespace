// Package main is root directory of this project
package main

import (
	"log"

	"github.com/gopherty/timespace/common/db"

	"github.com/gopherty/timespace/common"

	_ "github.com/go-sql-driver/mysql"

	"github.com/gin-gonic/gin"
	"github.com/gopherty/timespace/configs"
	"github.com/gopherty/timespace/router"
)

func main() {
	err := configs.Init()
	if err != nil {
		log.Fatalf("config init failed. %v", err)
	}

	// common module
	registers := []common.IRegister{
		db.Register{},
	}
	for _, reg := range registers {
		if err := reg.CheckIn(); err != nil {
			log.Fatalf("%s register failed. %v", reg.Name(), err)
		}
	}

	// config instance
	conf := configs.Instance()

	r := gin.Default()
	if conf.Server.Release {
		gin.SetMode(gin.ReleaseMode)
	}

	// router address
	router.Route(r)

	// run
	if conf.Server.CertFile != "" && conf.Server.KeyFile != "" {
		err = r.RunTLS(conf.Server.Address, conf.Server.CertFile, conf.Server.KeyFile)
	} else {
		err = r.Run(conf.Server.Address)
	}

	if err != nil {
		log.Fatalf("server run failed. %v", err)
	}
}
