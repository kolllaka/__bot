package messages

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"time"

	"honnef.co/go/tools/config"
)

type Conn struct {
	conn net.Conn

	connConf config.Config

	ChannelFromString chan string
	channelOutMsg     chan string
	channelOutCmd     chan string
}

func NewConn(conn net.Conn, cfg config.Config, msgOutChan chan string) *Conn {
	msgChan := make(chan string)
	//msgOutChan := make(chan string)

	return &Conn{
		conn:              conn,
		connConf:          cfg,
		ChannelFromString: msgChan,
		channelOutMsg:     msgOutChan,
	}
}

func (conn *Conn) ReadMessages() error {
	msgReader := bufio.NewReader(conn.conn)
	var err error
	//msgsChan := make(chan string)

	go func() {
		for {
			msg, err := msgReader.ReadString('\n')
			if err != nil {
				return
			}

			if strings.HasPrefix(msg, "PING") {
				conn.conn.Write([]byte("PONG :tmi.twitch.tv\r\n"))

				continue
			}

			//fmt.Println(msg)
			conn.ChannelFromString <- msg
		}
	}()

	return err
}

func (conn *Conn) WriteMessages() error {
	go func() {
		for {
			ms := <-conn.ChannelFromString
			fmt.Print(timestamp(), ms)

			// msg := fmt.Sprintf("PRIVMSG #%s :%s\r\n", conn.connConf.Channel, ms)
			// conn.conn.Write([]byte(msg))
		}
	}()

	return nil
}

func timestamp() string {
	return time.Now().Format("15:04:05")
}
