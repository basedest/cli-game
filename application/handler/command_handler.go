package handler

import (
	"strings"

	"../../domain/usecase"
)

// CommandHandler handles game commands
type CommandHandler struct {
	gameUseCase usecase.GameUseCase
}

// NewCommandHandler creates a new command handler
func NewCommandHandler(gameUseCase usecase.GameUseCase) *CommandHandler {
	return &CommandHandler{
		gameUseCase: gameUseCase,
	}
}

// HandleCommand processes a game command and returns the result
func (h *CommandHandler) HandleCommand(command string) string {
	words := strings.Split(command, " ")
	if len(words) == 0 {
		return "пустая команда"
	}
	
	action := words[0]
	
	switch action {
	case "идти":
		if len(words) < 2 {
			return "куда идти?"
		}
		return handleResult(h.gameUseCase.Go(words[1]))
		
	case "осмотреться":
		return h.gameUseCase.LookAround()
		
	case "взять":
		if len(words) < 2 {
			return "что взять?"
		}
		return handleResult(h.gameUseCase.TakeItem(words[1]))
		
	case "надеть":
		if len(words) < 2 {
			return "что надеть?"
		}
		return handleResult(h.gameUseCase.Equip(words[1]))
		
	case "применить":
		if len(words) < 2 {
			return "что применить?"
		}
		if len(words) < 3 {
			return "к чему применить?"
		}
		return handleResult(h.gameUseCase.Use(words[1], words[2]))
		
	case "рестарт":
		return h.gameUseCase.Restart()
		
	default:
		return "неизвестная команда"
	}
}

// handleResult processes the result and error from a use case
func handleResult(result string, err error) string {
	if err != nil {
		return err.Error()
	}
	return result
} 