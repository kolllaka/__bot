package ui

import (
	"fmt"

	"github.com/KoLLlaka/__bot/internal/configs"
	"github.com/KoLLlaka/__bot/internal/model"
)

type AppUI interface {
	WriteMsg() error
}

type appUI struct {
	cfg           configs.Config
	channelOutMsg chan *model.Message
}

func (aUI *appUI) WriteMsg() error {
	aUI.printMsg(<-aUI.channelOutMsg)

	return nil
}

func (aUI *appUI) printMsg(msg *model.Message) {
	dateMsg := msg.Date.Format("15:04:05")
	fmt.Printf("%s %s %s: %s\n", msg.Platform, dateMsg, msg.MsgAuthor, msg.Msg)
}
