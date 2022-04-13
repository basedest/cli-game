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

// Выводит информацию о доступных командах
func printHelp() {
	fmt.Println("Доступные команды:")
	fmt.Println("  осмотреться - осмотреть текущую комнату и её содержимое")
	fmt.Println("  идти [направление] - перейти в указанную комнату (например, 'идти коридор')")
	fmt.Println("  взять [предмет] - взять предмет (например, 'взять ключи')")
	fmt.Println("  надеть [предмет] - надеть предмет (например, 'надеть рюкзак')")
	fmt.Println("  применить [предмет] [цель] - использовать предмет на цели (например, 'применить ключи дверь')")
	fmt.Println("  рестарт - перезапустить игру")
	fmt.Println("  помощь - показать доступные команды")
	fmt.Println()
}

func main() {
	fmt.Println("Добро пожаловать в игру \"Очередной день\"!")
	fmt.Println("Ваша цель - собрать рюкзак, взять необходимые предметы и отправиться в университет.")
	fmt.Println()
	
	// Вывод списка команд при запуске
	printHelp()
	
	// Initialize dependencies
	initGame()
	
	// Start the game loop
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		command := scanner.Text()
		
		// Если команда "помощь", показываем список команд
		if command == "помощь" {
			printHelp()
			continue
		}
		
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
	
	// Обрабатываем команду помощи отдельно для тестов
	if command == "помощь" {
		return "список команд показан"
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
