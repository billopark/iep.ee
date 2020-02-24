package webserver

import (
	"context"
	"fmt"
	"github.com/billopark/iep.ee/config"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

func route(r *mux.Router) {
	r.PathPrefix("/").HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte(`<html><body>Hello! This is IEPEE, DNS Wildcard Service for Any IP.
Please visit <a href="https://github.com/billopark/iep.ee">github</a> for detail</body></html>`))
	}).Methods("GET")
}

func Start(halt chan bool) {
	r := mux.NewRouter()
	route(r)
	srv := &http.Server{
		Addr:         fmt.Sprintf("[::]:%d", config.Get().WebPort),
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	go func() {
		<-halt
		if err := srv.Shutdown(context.Background()); err != nil {
			panic(err)
		} else {
			log.Println("Web Server halted")
		}
	}()

	return
}
