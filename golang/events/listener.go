package events

import (
	"jaysonhelseth/gopictureframe/shared"
	"log"
)

func KeyListener(character rune) {
	switch character {
	case '*':
		TogglePlayer()
	case '0':
		log.Printf("Exit")
		shared.AppWindow.Close()
	default:
		log.Printf("%c", character)
	}
}
