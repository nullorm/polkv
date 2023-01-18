package main

import "net/http"

type handler struct{}

func (handler) ServeHTTP(http.ResponseWriter, *http.Request) {}
