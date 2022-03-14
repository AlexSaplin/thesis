package main

import (
	"context"
	"encoding/json"
	"flag"
	"os"

	"lynx/pkg/api"
	"lynx/pkg/config"
)

var fs = flag.NewFlagSet("lynx", flag.ExitOnError)
var configPath = fs.String("config", "", "lynx server config path")

func mustLoadConfig(path string) config.LynxConfig {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	var cfg config.LynxConfig
	err = decoder.Decode(&cfg)
	if err != nil {
		panic(err)
	}
	return cfg
}

func main() {
	_ = fs.Parse(os.Args[1:])

	config := mustLoadConfig(*configPath)

	app, err := api.NewLynxAPI(context.Background(), config)
	if err != nil {
		panic(err)
	}
	app.Run()
}
