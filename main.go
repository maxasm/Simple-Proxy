package main

import (
	"net/http"
	"log"
	"os"
)

// logger for printing debug statements + FLAGS
var dl *log.Logger = log.New(os.Stdout, "[DEBUG]", log.Ldate)

const HTTP_PORT = 4554
const HTTPS_PORT = 5445

func main() {

}
