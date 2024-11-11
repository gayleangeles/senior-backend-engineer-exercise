package main

import (
	"log"
	"sync"
	"syndio/api"
	"syndio/db"

	"github.com/alexflint/go-arg"
)

var args struct {
	DbPath   string `arg:"env:DB_PATH" help:"Path to the SQLite database file."`
	BindJson string `arg:"env:BIND_JSON" help:"Address to bind the API to."`
}

func main() {
	arg.MustParse(&args)

	if args.DbPath == "" {
		args.DbPath = "../db/employees.db"
	}

	if args.BindJson == "" {
		args.BindJson = ":8080"
	}

	database := db.DB{}

	err := database.InitDB(args.DbPath)

	if err != nil {
		log.Fatal(err)
	}

	defer database.Database.Close()

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		log.Printf("Starting the API Server...\n")
		api.Serve(args.BindJson, database)
		wg.Done()
	}()

	wg.Wait()
}
