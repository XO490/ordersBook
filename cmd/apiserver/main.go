package main

import (
	"encoding/json"
	"flag"
	"log"
	"ordersBook/internal/app/apiserver"
	"os"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "conf", "config.json", "path to config file")
	flag.Parse()
}

func main() {
	var err error
	var file []byte
	config := apiserver.NewConfig()

	// Open JSON config
	if file, err = os.ReadFile(configPath); err != nil {
		log.Printf("[Error]: read config> %v", err)
	}

	if err = json.Unmarshal(file, &config); err != nil {
		log.Printf("[Error]: open config> %v", err)
	}

	apiserver.AvailableRoomsInit(config)

	s := apiserver.New(config)
	log.Printf("[Error]: start server> %v", s.Start())

}
