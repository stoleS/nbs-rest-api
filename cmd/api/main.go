package main

import (
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"
	"github.com/stoleS/nbs-rest/internal/handlers"
	"github.com/stoleS/nbs-rest/internal/middleware"
	"github.com/stoleS/nbs-rest/internal/tools"
)

func main() {
	tools.InitEnv()
	tools.InitSoap()

	router := http.NewServeMux()
	handlers.Handler(router)

	stack := middleware.CreateStack(
		middleware.Logging,
	)

	server := http.Server{
		Addr:    tools.GetHostUrl(),
		Handler: stack(router),
	}

	fmt.Printf("Starting GO API service at %s ...\n", tools.GetHostUrl())

	err := server.ListenAndServe()

	if err != nil {
		log.Error(err)
	}
}
