package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/nero2009/pricecompare/internal/cache"
	"github.com/nero2009/pricecompare/internal/handlers"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetReportCaller(true)
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

	err := http.ListenAndServe(":8091", r)
	if err != nil {
		log.Error(err)
	}

	// fmt.Println(document.Text())
}
