package main

import (
	"bufio"
	"fmt"
	"os"

	"./application/handler"
	"./domain/usecase"
	"./infrastructure/repository"
)

func main() {
	fmt.Println("Добро пожаловать в игру \"Очередной день\"!")
	
	// Initialize dependencies
	gameRepo := repository.NewGameRepository()
	gameRepo.InitializeWorld()
	
	gameUseCase := usecase.NewGameUseCase(gameRepo)
	commandHandler := handler.NewCommandHandler(gameUseCase)
	
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
	// Initialize dependencies for tests
	gameRepo := repository.NewGameRepository()
	gameUseCase := usecase.NewGameUseCase(gameRepo)
	commandHandler := handler.NewCommandHandler(gameUseCase)
	
	return commandHandler.HandleCommand(command)
}

func initGame() {
	// This function is required by tests but its functionality
	// is now in the repository's InitializeWorld method
	gameRepo := repository.NewGameRepository()
	gameRepo.InitializeWorld()
}
