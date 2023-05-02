package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

const port string = ":8080"
const defaultTimeout int = 35
const routeTimeoutKey = "ROUTE_TIMEOUT"

var timeout int = defaultTimeout

func main() {
	log.Println("Starting Server")
	timeoutStr, ok := os.LookupEnv(routeTimeoutKey)

	if ok {
		var err error
		timeout, err = strconv.Atoi(timeoutStr)
		log.Println("ENV key available: ", timeout)

		if err != nil {
			timeout = defaultTimeout
			log.Println("ENV key not a integer")
		}
	}
	r := http.NewServeMux()

	r.HandleFunc("/", testHandler)

	err := http.ListenAndServe(port, r)

	if err != nil {
		log.Println("Error on server start: ", err.Error())
	}
}

func testHandler(w http.ResponseWriter, r *http.Request) {

	time.Sleep(time.Second * time.Duration(timeout))
	fmt.Fprintf(w, "Timeout is %v", timeout)
}
