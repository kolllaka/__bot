package twitch2

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

type Opts struct {
	UserName string
	Token    string
	Channel  string
}

type Conn struct {
	conn       net.Conn
	channelIn  chan string
	channelOut chan string
}

func NewTwitchClient(opts Opts) *Conn {
	conn, err := net.Dial("tcp", "irc.chat.twitch.tv:6667")
	if err != nil {
		log.Fatalln("Failed to dial to the IRC:", err)
	}

	msg := fmt.Sprintf(
		"PASS %s\r\nNICK %s\r\nJOIN #%s\r\n",
		opts.Token,
		opts.UserName,
		opts.Channel,
	)
	conn.Write([]byte(msg))

	connToTw := &Conn{
		conn:       conn,
		channelIn:  make(chan string),
		channelOut: make(chan string),
	}
	connToTw.read()
	connToTw.write()

	return connToTw
}

func (c *Conn) Close() error {
	if err := c.conn.Close(); err != nil {
		log.Fatal(err)
	}

	log.Println("канал удачно закрыт")
	return nil
}

func (c *Conn) Read() string {
	return <-c.channelIn
}

func (c *Conn) Write(msg string) {
	c.channelOut <- msg
}

func (c *Conn) read() error {
	msgReader := bufio.NewReader(c.conn)
	var err error

	go func() {
		for {
			msg, err := msgReader.ReadString('\n')
			if err != nil {
				return
			}

			if strings.HasPrefix(msg, "PING") {
				c.conn.Write([]byte("PONG :tmi.twitch.tv\r\n"))

				continue
			}

			c.channelIn <- msg
		}
	}()

	return err
}

func (c *Conn) write() error {
	go func() {
		for {
			ms := <-c.channelOut
			log.Print(ms)

			// msg := fmt.Sprintf("PRIVMSG #%s :%s\r\n", conn.connConf.Channel, ms)
			// conn.conn.Write([]byte(msg))
		}
	}()

	return nil
}

func timestamp() string {
	return time.Now().Format("15:04:05")
}
