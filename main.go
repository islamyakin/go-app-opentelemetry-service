package main

import (
	"log"
	"os"
	"strconv"
)

const name = "order-service"

var port = "8080"
var otel_host = "127.0.0.1:4317"
var sampler = float64(1)

func init() {

	e_port, exist := os.LookupEnv("PORT")
	if exist {
		port = e_port
	}
	e_otel_host, exist := os.LookupEnv("OTEL_HOST")
	if exist {
		otel_host = e_otel_host
	}
	e_sampler, exist := os.LookupEnv("OTEL_SAMPLER_RATIO")
	if exist {
		e_sampler_float, err := strconv.ParseFloat(e_sampler, 64)
		if err != nil {
			log.Panic(err)
		}
		sampler = e_sampler_float
	}
}
