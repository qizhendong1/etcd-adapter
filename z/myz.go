package main

import (
	"fmt"
	"log"

	"github.com/mitchellh/mapstructure"
	"gopkg.in/yaml.v2"
)

// Config 结构体
type Config struct {
	Server   ServerConfig   `yaml:"server" mapstructure:"server"`
	Database DatabaseConfig `yaml:"database" mapstructure:"database"`
}

// ServerConfig 结构体
type ServerConfig struct {
	Host string `yaml:"host" mapstructure:"host"`
	Port int    `yaml:"port" mapstructure:"port"`
}

// DatabaseConfig 结构体
type DatabaseConfig struct {
	User     string `yaml:"user" mapstructure:"user"`
	Password string `yaml:"password" mapstructure:"password"`
	Name     string `yaml:"name" mapstructure:"name"`
}

func main() {
	// 1. 读取 YAML 文件内容
	yamlFile := []byte(`
server:
  host: localhost
  port: 8080
database:
  user: admin
  password: secret
  name: mydb
`)

	// 2. 解析 YAML 到 map
	var data map[string]interface{}
	err := yaml.Unmarshal(yamlFile, &data)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	// 3. 使用 mapstructure 将 map 转换为结构体
	var config Config
	err = mapstructure.Decode(data, &config)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	// 4. 输出结果
	fmt.Printf("Server Host: %s\n", config.Server.Host)
	fmt.Printf("Server Port: %d\n", config.Server.Port)
	fmt.Printf("Database User: %s\n", config.Database.User)
	fmt.Printf("Database Name: %s\n", config.Database.Name)
}
