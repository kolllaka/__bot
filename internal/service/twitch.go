package service

import "github.com/KoLLlaka/__bot/internal/model"

type twitchService interface {
	ReadMsg(chan model.Message) error
	WriteMsg(model.Message,chan model.Message) error
}
