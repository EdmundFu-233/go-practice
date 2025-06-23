package main

import (
	"fmt"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Server Server `yaml:"server"`
	DB     DB     `yaml:"database"`
}

type Server struct {
	Port int    `yaml:"port"`
	Host string `yaml:"host"`
}

type DB struct {
	DSN  string `yaml:"dsn"`
	Pool int    `yaml:"pool"`
}

func main() {
	configYAML := `
server:
  port: 8080
  host: 0.0.0.0
database:
  dsn: postgres://user:pass@localhost/db
  pool: 10
`
	var config Config
	yaml.Unmarshal([]byte(configYAML), &config)
	fmt.Printf("%+v\n", config)
}
