package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/websocket"
)

func main() {
	origin := flag.String("origin", "", "Origin header")
	flag.Parse()
	header := http.Header{}
	header.Set("Origin", *origin)
	conn, response, err := websocket.DefaultDialer.Dial("ws://127.0.0.1:8080/ws/live", header)
	if err != nil {
		if response != nil {
			fmt.Fprintf(os.Stderr, "websocket failed: HTTP %d: %v\n", response.StatusCode, err)
		} else {
			fmt.Fprintf(os.Stderr, "websocket failed: %v\n", err)
		}
		os.Exit(1)
	}
	defer conn.Close()
	fmt.Println("websocket connected")
}
