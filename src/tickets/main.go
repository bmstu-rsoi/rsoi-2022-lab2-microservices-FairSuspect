package main

import (
	"flag"
	"http-rest-api/internal/app/apiserver"
	"http-rest-api/store"
	"log"

	"github.com/BurntSushi/toml"
)

var (
	configPath string
	repo       store.TicketRepository
)

func init() {
	log.Default().Println(flag.Args())
	// if len(os.Args) == 2 {

	flag.StringVar(&configPath, "config", "configs/apiserver.toml", "path to config file")
	// }
	log.Default().Println("Config path: " + configPath)
}

func main() {

	flag.Parse()
	configFlag := flag.Lookup("config")
	config := apiserver.NewConfig()
	log.Default().Println(configFlag)
	if configFlag != nil {
		_, err := toml.DecodeFile(configPath, config)
		if err != nil {
			log.Fatal(err)
		}
	}

	st := store.New(config.Store)
	repo = *st.Ticket()

	s := apiserver.New(config)
	log.Default().Println(configFlag)

	if err := s.Start(); err != nil {
		log.Fatal(err)
	}

}
