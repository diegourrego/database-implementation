package main

import (
	"database_implementation/internal/application"
	"fmt"
	"github.com/go-sql-driver/mysql"
)

func main() {
	//implementation.Init()
	addrCfg := ":8080"
	mysqlCfg := mysql.Config{
		User:      "root",
		Passwd:    "Diego123",
		Net:       "tcp",
		Addr:      "localhost:3306",
		DBName:    "storage_db",
		ParseTime: true,
	}

	cfg := application.ConfigServerChi{Addr: addrCfg, MySQLDSN: mysqlCfg.FormatDSN()}
	server := application.NewServerChi(cfg)
	if err := server.Run(); err != nil {
		fmt.Println(err)
		return
	}
}
