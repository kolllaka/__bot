package handlers

type HandlerChannel interface {
	NewTwitchHandler() TwitchHandler
}
