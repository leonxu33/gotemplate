package http

import (
	"context"
	"fmt"
	"net/http"
	"path"

	log "github.com/cihub/seelog"
	"github.com/lyokalita/gotemplate/src/config"
	"github.com/rs/cors"
)

var server *http.Server

func StartServer() {
	sm := constructServerMux()
	addr := getServerAddr()

	server = &http.Server{
		Addr:    addr,
		Handler: sm,
		// IdleTimeout:  time.Duration(120) * time.Second,
		// ReadTimeout:  5 * time.Second,
		// WriteTimeout: 5 * time.Second,
	}

	go func() {
		err := server.ListenAndServe()
		//err := server.ListenAndServeTLS(".cert/naspublic.crt", ".cert/naspublic.key")
		if err != nil {
			log.Error(err)
		}
	}()

	log.Infof("listens at %s", path.Join(addr, config.ApiPath))
}

func StopServer(ctx context.Context) {
	err := server.Shutdown(ctx)
	if err != nil {
		log.Info(err)
	}
}

func constructServerMux() *http.ServeMux {
	authCors := cors.New(cors.Options{
		AllowedOrigins: config.AllowOrigin,
		AllowedMethods: []string{http.MethodPost, http.MethodGet},
	})
	handler := authCors.Handler(NewTestHandler())

	sm := http.NewServeMux()
	sm.Handle(path.Join(config.ApiPath, "test"), handler)

	return sm
}

func getServerAddr() string {
	return fmt.Sprintf("%s:%d", config.ServerHost, config.ServerPort)
}
