package bot

import "github.com/Amazeful/Amazefulbot/message"

type Module interface {
	ProcessMessage(message *message.ChatMessage)
}
