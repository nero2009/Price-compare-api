package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/nero2009/pricecompare/internal/cache"
	"github.com/nero2009/pricecompare/internal/handlers"
	log "github.com/sirupsen/logrus"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	log.SetReportCaller(true)
	db, err := sql.Open("mysql", "test:test@(127.0.0.1:3306)/pricecompare?parseTime=true")

	if err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}
	errr := db.Ping()

	if errr != nil {
		log.Error(errr.Error())
		os.Exit(1)
	}
	var r *chi.Mux = chi.NewRouter()
	cacheManager := cache.NewCache()
	handlers.Handler(r, cacheManager)

	fmt.Println("Starting GO API2 service....")

	// write fancy multiple line  graphicsprintln that says "Starting GO API2 service...."
	fmt.Println((`
		______ _   _ _____  _    _ _______ _    _
		|  ____| \ | |  __ \| |  | |__   __| |  | |
		| |__  |  \| | |__) | |  | |  | |  | |__| |
		|  ___| . ` + "`" + ` |  _  /| |  | |  | |  |  __  |
		| |____| |\  | | \ \| |__| |  | |  | |  | |
		|______|_| \_|_|  \_\\____/   |_|  |_|  |_|
	`))

	errrr := http.ListenAndServe(":8091", r)
	if errrr != nil {
		log.Error(err)
	}
}
