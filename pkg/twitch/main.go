package main

import (
	"fmt"
	"log"
	"time"

	twitch "github.com/KoLLlaka/__bot/pkg/twitch/v2"
)

func main() {
	opts := twitch.Opts{
		UserName: "bobcehekolliaka",
		Token:    "oauth:3j4z03urgvvjtzbx78vc0by4w29iry",
		Channel:  "kolliaka",
	}

	connTw := twitch.NewTwitchClient(opts)
	go func() {
		for {
			log.Print(connTw.Read())
		}
	}()
	go func() {
		for {
			time.Sleep(time.Second * 30)
			connTw.Write(time.Now().Format("15:04:05"))
		}
	}()

	fmt.Scanln()
	connTw.Close()
}
