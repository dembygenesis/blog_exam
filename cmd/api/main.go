package main

import (
	"github.com/dembygenesis/blog_exam/pkg/api"
	"github.com/dembygenesis/blog_exam/pkg/config"
	"github.com/dembygenesis/blog_exam/pkg/logic"
	"github.com/dembygenesis/blog_exam/pkg/store"
	"github.com/dembygenesis/blog_exam/pkg/store/mongo"
	"github.com/dembygenesis/blog_exam/pkg/store/mysql"
	_ "github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
	"os"
)

func main() {
	// Read env
	cfg, err := config.ReadConfig(".env")
	if err != nil {
		log.Fatalf("Error attempting to read config: %v", err.Error())
		os.Exit(1)
	}

	// Set config
	err = cfg.SetConfig()
	if err != nil {
		log.Fatalf("error setting config: %v", err.Error())
	}

	// Setup store
	var _store store.Store
	if cfg.Database.Driver == "mysql" {
		// We use the MYSQL implementation for use
		_store, err = mysql.NewStore(&cfg.Database)
		if err != nil {
			log.Fatalf("error establishing MYSQL store: %v", err.Error())
		}
	} else if cfg.Database.Driver == "mongo" {
		// We can add a mongo implementation here too, if desired.
		_store, _ = mongo.NewStore()
	} else {
		log.Fatal("store is undefined")
	}

	// Setup logic handler
	_logic := logic.NewLogicHandler(_store)

	// Setup app server
	server := api.NewServer(_logic, cfg.AppPort)

	// Start server
	err = server.Start()
	if err != nil {
		log.Fatalf("error starting server: %v", err.Error())
	}
}