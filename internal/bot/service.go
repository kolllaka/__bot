package bot

import (
	"github.com/KoLLlaka/__bot/internal/configs"
	"github.com/KoLLlaka/__bot/internal/model"
)

type AppBot interface {
	ReadMsg() (*model.Message, error)
	WriteMsg(*model.Message) error
	WriteCmd(*model.Message) error
}

type appBot struct {
	cfg           configs.Config
	channelFrom   chan *model.Message
	channelOutMsg chan *model.Message
	channelOutCmd chan *model.Message
}

func (ap *appBot) ReadMsg() (*model.Message, error) {
	readMsg := <-ap.channelFrom

	if !readMsg.IsCommand {
		ap.WriteMsg(readMsg)

		return readMsg, nil
	}

	if ap.cfg.WriteCmd {
		ap.WriteMsg(readMsg)
	}

	ap.WriteCmd(readMsg)

	return readMsg, nil
}

func (ap *appBot) WriteMsg(msg *model.Message) error {
	if msg != nil {
		ap.channelOutMsg <- msg
	}

	return nil
}

func (ap *appBot) WriteCmd(cmdMsg *model.Message) error {
	if cmdMsg != nil {
		ap.channelOutCmd <- cmdMsg

		if ap.cfg.WriteCmd {
			ap.WriteMsg(cmdMsg)
		}
	}

	return nil
}
