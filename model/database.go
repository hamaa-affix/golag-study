package model

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"

)

type DBConf struct {
	UserName   string
	PassWord   string
	DBName     string
	Port       int
	DriverName string
	Host       string
	Err        error
}

func NewDBConf(
	userName string,
	pass string,
	dbName string,
	port int,
	driverName string,
	host string,
) *DBConf {
	return &DBConf{
		UserName:   userName,
		PassWord:   pass,
		DBName:     dbName,
		Port:       port,
		DriverName: driverName,
		Host:       host,
	}
}

func (d *DBConf) CreateEndpoint() string {
	if d.Err != nil {
		log.Fatal("missing create db connection endpoint\nerror is %v", d.Err)
	}
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", d.UserName, d.PassWord, d.Host, d.Port, d.DBName)
}

func New() (*sql.DB, error)  {
	conf := NewDBConf(
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_DATABASE"),
		3306,
		"mysql",
		"mysql",
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