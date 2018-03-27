package main

import (
    "fmt"
    "log"
    "net/http"

    "github.com/gorilla/websocket"
)

var clients = make(map[*websocket.Conn]bool)
var clients = make(map[*websocket.Conn]bool)

var upgrader = websocket.Upgrader{}

type Message struct {
        Username string `json:"username"`
        Message  string `json:"message"`
}


