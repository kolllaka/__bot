package connect

import (
	"fmt"
	"log"
	"net"

)

func ConnectToIRC(cfg config.Config) net.Conn {
	conn, err := net.Dial("tcp", "irc.chat.twitch.tv:6667")
	if err != nil {
		log.Fatalln("Failed to dial to the IRC:", err)
	}

	// fmt.Fprintf(conn, "PASS %s\r\n", cfg.Token)
	// fmt.Fprintf(conn, "NICK %s\r\n", cfg.User)
	// fmt.Fprintf(conn, "JOIN #%s\r\n", cfg.Channel)
	msg := fmt.Sprintf("PASS %s\r\nNICK %s\r\nJOIN #%s\r\n", cfg.Token, cfg.User, cfg.Channel)
	conn.Write([]byte(msg))

	return conn
}
