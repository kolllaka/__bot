package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"
	"time"

	"github.com/KoLLlaka/__bot/internal/model"
	twitch "github.com/KoLLlaka/__bot/pkg/twitch/v2"

	"github.com/gorilla/websocket"
	"github.com/joho/godotenv"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

var (
	token   string
	user    string
	channel string
)

func init() {
	if err := godotenv.Load(".env"); err != nil {
		log.Println("No .env file found")
	}

	token = os.Getenv("TWITCH_OAUTH_TOKEN")
	if len(token) == 0 {
		log.Fatalln("TWITCH_OAUTH_TOKEN is missing")
	}
	channel = os.Getenv("TWITCH_CHANNEL")
	if len(token) == 0 {
		log.Fatalln("TWITCH_CHANNEL is missing")
	}
	user = os.Getenv("TWITCH_USER")
	if len(token) == 0 {
		log.Fatalln("TWITCH_USER is missing")
	}
}

func main() {
	// канал с текстовыми сообщениями
	channelOutMsg := make(chan string)
	// // канал с командами
	// channelOutCmd := make(chan model.Message)

	// twitch
	opts := twitch.Opts{
		UserName: user,
		Token:    token,
		Channel:  channel,
	}
	connTw := twitch.NewTwitchClient(opts)
	defer connTw.Close()

	go func() {
		for {
			msg := connTw.Read()

			regex := `:(.*)!.*@.*.tmi.twitch.tv PRIVMSG #(.*) :(.*)`
			re, err := regexp.Compile(regex)
			if err != nil {
				log.Fatalln("Failed to compile the regex")
			}

			parseMsg := re.FindStringSubmatch(msg)
			if parseMsg == nil {
				continue
			}

			message := model.Message{
				Platform:  "twitch",
				MsgAuthor: parseMsg[1],
				Channel:   parseMsg[2],
				Msg:       parseMsg[3],
				Date:      time.Now(),
				IsCommand: false,
				NameCmd:   "",
			}

			//fmt.Println(printMessage(message))
			channelOutMsg <- printMessage(message)
		}
	}()

	//server
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./templates/index.html")
	})

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Fatal(err)
		}

		go func(conn *websocket.Conn) {
			for {
				msg := <-channelOutMsg
				//log.Println("in handle", msg)
				conn.WriteMessage(1, []byte(msg))
			}
		}(conn)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))

	// trovoService, _ := handler.Trovo(channelFrom)
	// twitchService, _ := handler.Twitch(channelFrom)
	fmt.Scanln()
}

func printMessage(msg model.Message) string {
	return fmt.Sprintf("%s %s %s: %s", msg.Date.Format("15:04:05"), msg.Platform, msg.MsgAuthor, msg.Msg)
}
