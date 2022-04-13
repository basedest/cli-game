package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/basedest/cli-game/application/handler"
	domainrepo "github.com/basedest/cli-game/domain/repository"
	"github.com/basedest/cli-game/domain/usecase"
	repo "github.com/basedest/cli-game/infrastructure/repository"
)

// Global variables for test compatibility
var (
	gameRepo       domainrepo.GameRepository
	gameUseCase    usecase.GameUseCase
	commandHandler *handler.CommandHandler
)

func main() {
	fmt.Println("Добро пожаловать в игру \"Очередной день\"!")
	
	// Initialize dependencies
	initGame()
	
	// Start the game loop
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		command := scanner.Text()
		result := commandHandler.HandleCommand(command)
		fmt.Println(result)
	}
}

// For test compatibility
func handleCommand(command string) string {
	// Use the existing initialized dependencies
	if commandHandler == nil {
		initGame()
	}
	return commandHandler.HandleCommand(command)
}

func initGame() {
	// Initialize dependencies
	if gameRepo == nil {
		gameRepo = repo.NewGameRepository()
	} else {
		// For tests - completely reset the game state
		gameRepo.ResetGame()
	}
	
	gameUseCase = usecase.NewGameUseCase(gameRepo)
	commandHandler = handler.NewCommandHandler(gameUseCase)
}
