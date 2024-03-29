package model

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"

)

type dBConf struct {
	UserName   string
	PassWord   string
	DBName     string
	Port       string
	DriverName string
	Host       string
	Err        error
}

func newDBConf(
	userName string,
	pass string,
	dbName string,
	port string,
	driverName string,
	host string,
) *dBConf {
	return &dBConf{
		UserName:   userName,
		PassWord:   pass,
		DBName:     dbName,
		Port:       port,
		DriverName: driverName,
		Host:       host,
	}
}

func (d *dBConf) CreateEndpoint() string {
	if d.Err != nil {
		log.Fatal("missing create db connection endpoint\nerror is %v", d.Err)
	}
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", d.UserName, d.PassWord, d.Host, d.Port, d.DBName)
}

func new() (*sql.DB, error)  {
	conf := newDBConf(
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_DATABASE"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_DRIVER_NAME"),
		os.Getenv("DB_HOST"),
	)

	db, err := sql.Open("mysql", conf.CreateEndpoint())
	if err != nil {
		conf.Err = fmt.Errorf("mysql connection error\nerror message:\n %v", err)
		return db, conf.Err
	}

	// DBへの接続確認を行う
    err = db.Ping()
    if err != nil {
		conf.Err = fmt.Errorf("mysql connection error\nerror message:\n %v", err)
		return db, conf.Err
    }

	return db, nil
}
