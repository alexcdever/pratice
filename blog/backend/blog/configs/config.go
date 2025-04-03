package configs

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
)

var Con Config

func init() {
	readFromConfigFIle("./configs/config.yaml")
}

func readFromConfigFIle(absoluteFilePath string) {
	viper.SetConfigFile(absoluteFilePath)
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("failed to read Config from file: %s", err)
	}

	if err := viper.Unmarshal(&Con); err != nil {
		log.Fatalf("failed to unmarshal Config data from file: %s", err)
	}
}

type Config struct {
	Name     string
	Host     string
	Port     int
	Source   string
	Mysql    mysqlConfig
	Postgres pgConfig
}

type DbConfig interface {
	ConnString() (connString string)
}

type mysqlConfig struct {
	Database  string
	Host      string
	Port      int
	User      string
	Password  string
	Charset   string
	ParseTime bool
	Loc       string
}

func (m mysqlConfig) MysqlConnString() (connString string) {
	// root:admin@tcp(127.0.0.1:3306)/blog?charset=utf8&parseTime=true&loc=Asia%2FShanghai
	connString = fmt.Sprintf("%s:%s@tcp(%s:%v)/%s?charset=%s&parseTime=%v&loc=%s", m.User, m.Password, m.Host, m.Port, m.Database, m.Charset, m.ParseTime, m.Loc)
	return connString
}

type pgConfig struct {
	Database string
	Host     string
	Port     int
	User     string
	Password string
	SslMode  string
	TimeZone string
}

func (p pgConfig) PgConnString() (connString string) {
	// host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai
	connString = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%v sslmode=%s TimeZone=%s", p.Host, p.User, p.Password, p.Database, p.Port, p.SslMode, p.TimeZone)
	return connString
}
