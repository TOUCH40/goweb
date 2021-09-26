package main

import (
	"net/http"
	// "strings"
)

const (
	mutexLocked   = 1 << iota // 1
	mutexWoken                // 2
	mutexStarving             // 4
)

func main() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static", fs))
}

