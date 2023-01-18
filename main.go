package main

import (
	"net/http"

	"github.com/nullorm/polkv/app"
	"go.uber.org/zap"

	_ "github.com/nullorm/polkv/configs"
)

func main() {
	// todo- replace with grpc
	s := &http.Server{
		Addr:    ":8000",
		Handler: handler{},
	}
	if err := s.ListenAndServe(); err != nil {
		app.Log.Fatal("server shutdown with err", zap.Error(err))
	}
}
