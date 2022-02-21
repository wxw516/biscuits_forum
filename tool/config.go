package tool

import (
	"bufio"
	"encoding/json"
	"os"
)

type Config struct {
	WebName string `json:"web_name"`
	WebMode string `json:"web_mode"`
	WebHost string `json:"web_host"`
	WebPort string `json:"web_port"`
}

type DatabaseConfig struct {
	Driver   string `json:"driver"`
	User     string `json:"user"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	DbName   string `json:"db_name"`
	Charset  string `json:"charset"`
}

var cfg *Config = nil

//解析config文件
func ParseConfig(path string)(*Config,error)  {
	file ,err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	decoder := json.NewDecoder(reader)
	if err = decoder.Decode(&cfg);err != nil{
		return nil, err
	}
	return cfg, nil
}