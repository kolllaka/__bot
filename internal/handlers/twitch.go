package handlers

import "github.com/KoLLlaka/__bot/internal/model"

type TwitchHandler interface {
	ReadMsg(chan model.Message) error
	WriteMsg(model.Message, chan model.Message) error
}

type twitchHandler struct{}

func (twh *twitchHandler) ReadMsg(channelFrom chan model.Message) error {
	return nil
}

func (twh *twitchHandler) WriteMsg(msg model.Message, channelOutMsg chan model.Message) error {
	return nil
}
