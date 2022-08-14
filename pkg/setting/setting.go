package setting

import (
	"log"

	"gopkg.in/ini.v1"
)

var (
	// server
	AppMode  string
	HttpPort int

	// database
	Db         string
	DbHost     string
	DbPort     int
	DbUser     string
	DbPassword string
	DbName     string
)

func init() {
	file, err := ini.Load("config/config.ini")
	if err != nil {
		log.Fatal(err)
	}
	LoadServer(file)
	LoadDatabase(file)
}

func LoadServer(file *ini.File) {
	AppMode = file.Section("server").Key("AppMode").MustString("debug")
	HttpPort = file.Section("server").Key("HttpPort").MustInt(3000)
}

func LoadDatabase(file *ini.File) {
	Db = file.Section("database").Key("Database").MustString("Db")
	DbHost = file.Section("database").Key("DbHost").MustString("localhost")
	DbPort = file.Section("database").Key("DbPort").MustInt(3306)
	DbUser = file.Section("database").Key("DbUser").MustString("go-blog")
	DbPassword = file.Section("database").Key("DbPassword").MustString("go-blog")
	DbName = file.Section("database").Key("DbName").MustString("go-blog")
}
