package bot

import (
	"fmt"

	"github.com/KoLLlaka/__bot/internal/model"
)

func (ap *appBot) ParseCommand() error {
	for {
		select {
		case cmdMsg := <-ap.channelOutCmd:
			doCmd(cmdMsg)
			ap.WriteMsg(cmdMsg)
		}
	}
}

func doCmd(cmd *model.Message) func() {
	switch cmd.NameCmd {
	case "любовь":
		cmd.Msg = fmt.Sprintf("%s любит Лулу на 100%%", cmd.MsgAuthor)
		cmd.IsCommand = false
		return nil
	}

	return nil
}
