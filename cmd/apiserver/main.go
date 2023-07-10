package main

import (
	"Practice3/internal/app/apiserver"
	"flag"
	"log"
	"os"

	"github.com/BurntSushi/toml"
)

var (
	configPath      string
	environmentPath = os.Getenv("GOPATH")
)

func init() {
	flag.StringVar(&configPath, "config-path", "\\Practice3\\configs\\apiserver.toml", "path to config file")
}

func main() {
	flag.Parse()

	config := apiserver.NewConfig()

	_, err := toml.DecodeFile(environmentPath+configPath, config)

	if err != nil {
		log.Fatal(err)
		return
	}

	s := apiserver.New(config)

	if err := s.Start(); err != nil {
		log.Fatal(err)
		return
	}
}
