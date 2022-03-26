package app

import (
	"fmt"
	"log"

	"go-risko/config"

	"net/http"

	"github.com/rs/cors"
)

func StartApp() {
	webServerMsg := fmt.Sprintf("Web server running on port:%v", config.WEB_PORT)
	log.Println(webServerMsg)
	handler := cors.AllowAll().Handler((router()))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", config.WEB_PORT), handler))
}