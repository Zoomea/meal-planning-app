package main

import (
	"fmt"
	"github.com/Zoomea/meal-planning-app/web"
	"os"
)

const (
	serveDir  = "./public"
	servePort = 8080
)

func main() {
	err := web.Serve(serveDir, servePort)
	if err != nil {
		fmt.Printf("err: %s\n", err)
		os.Exit(1)
	}
}
