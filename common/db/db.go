package db

import (
	"errors"

	"github.com/gopherty/timespace/configs"

	"github.com/go-xorm/xorm"
)

var db *xorm.Engine

// Register implement IRegister interface
type Register struct {
}

// Name module name
func (Register) Name() string {
	return "Common.DB"
}

// CheckIn register db module
func (Register) CheckIn() (err error) {
	conf := configs.Instance()

	if conf.DB.Driver == "" || conf.DB.Source == "" {
		return errors.New("driver or source not be empty")
	}

	engine, err := xorm.NewEngine(conf.DB.Driver, conf.DB.Source)
	if err != nil {
		return err
	}

	engine.SetMaxOpenConns(conf.DB.MaxOpenConn)
	engine.SetMaxIdleConns(conf.DB.MaxIdleConn)

	engine.ShowSQL(conf.DB.ShowSQL)

	if conf.DB.Cache != 0 {
		cache := xorm.NewLRUCacher(xorm.NewMemoryStore(), conf.DB.Cache)
		engine.SetDefaultCacher(cache)
	}

	err = engine.Ping()
	if err != nil {
		return
	}

	db = engine

	// // 创建用户相关的表
	// err = createTable(&Administrator{})
	// if err != nil {
	// 	return
	// }
	// // 同步表结构
	// err = db.Sync2(&Administrator{})
	// if err != nil {
	// 	return
	// }
	//
	// // generate administrator
	// ok, err := db.Get(&Administrator{
	// 	User: cnf.User.Name,
	// })
	// if err != nil {
	// 	return
	// }
	// if ok {
	// 	return
	// }
	//
	// var passwd string
	// h := sha512.New()
	// _, err = h.Write([]byte(cnf.User.Passwd))
	// if err != nil {
	// 	return
	// }
	// passwd = hex.EncodeToString(h.Sum(nil))[:18]
	// _, err = db.InsertOne(&Administrator{
	// 	User:     cnf.User.Name,
	// 	Password: passwd,
	// })
	return
}

func createTable(beans ...interface{}) (err error) {
	var exists bool
	for _, bean := range beans {
		exists, err = db.IsTableExist(bean)
		if err != nil {
			return
		}
		if !exists {
			err = db.CreateTables(bean)
			if err != nil {
				return
			}
		}
	}
	return
}

// Engine global database engine
func Engine() *xorm.Engine {
	return db
}
