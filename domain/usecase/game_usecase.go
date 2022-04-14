package usecase

import (
	"github.com/basedest/cli-game/domain/entity"
	"github.com/basedest/cli-game/domain/repository"
)

// GameUseCase defines the game logic operations
type GameUseCase interface {
	Go(room string) (string, error)
	LookAround() string
	TakeItem(item string) (string, error)
	Equip(item string) (string, error)
	Use(item, obj string) (string, error)
	Restart() string
}

// gameUseCase implements the GameUseCase interface
type gameUseCase struct {
	repo repository.GameRepository
}

// NewGameUseCase creates a new game use case
func NewGameUseCase(repo repository.GameRepository) GameUseCase {
	return &gameUseCase{
		repo: repo,
	}
}

// Go moves the player to a different room
func (g *gameUseCase) Go(roomName string) (string, error) {
	player := g.repo.GetPlayer()
	
	for _, entrance := range player.Location.Entrances {
		if entrance.Room.Name == roomName {
			if entrance.Door != nil && entrance.Door.Locked {
				return "", entity.NewGameError("дверь закрыта")
			}
			player.Location = entrance.Room
			
			// Если игрок попал на улицу, отмечаем цель как выполненную
			if player.Location.Name == "улица" {
				idx := entity.FindGoalIndex(player.Goals, "идти в универ")
				if idx != -1 {
					player.Goals[idx].Completed = true
					
					// Проверяем, все ли цели выполнены
					allGoalsCompleted := true
					for _, goal := range player.Goals {
						if !goal.Completed {
							allGoalsCompleted = false
							break
						}
					}
					
					// Если все цели выполнены, выводим торжественное сообщение
					if allGoalsCompleted {
						return onGoMessage(player.Location) + "\n🎉 Поздравляем! Вы успешно собрали рюкзак и добрались до университета! Миссия выполнена! 🎉", nil
					}
				}
			}
			
			return onGoMessage(player.Location), nil
		}
	}
	
	return "", entity.NewGameError("нет пути в " + roomName)
}

// LookAround returns a description of the current room
func (g *gameUseCase) LookAround() string {
	player := g.repo.GetPlayer()
	var res string
	
	// Special case for kitchen
	if player.Location.Name == "кухня" {
		res += "ты находишься на кухне, "
	}

	// Special case for street
	if player.Location.Name == "улица" {
		res += "ты находишься на улице, "
	}
	
	itemsFound := false
	for _, object := range player.Location.Objects {
		if container, ok := object.(entity.Container); ok && !container.IsEmpty() {
			res += container.String() + ", "
			itemsFound = true
		}
	}
	
	// Special case for empty room that's not kitchen and street
	if !itemsFound && player.Location.Name != "кухня" && player.Location.Name != "улица" {
		res += "пустая комната"
	} else if itemsFound {
		// Remove trailing comma and space
		res = res[:len(res)-2]
	}
	
	// Add goals except in the room
	if player.Location.Name != "комната" {
		goalsString := player.GetGoalsString()
		if goalsString != "" {
			// Добавляем запятую только если есть предыдущий текст и мы добавляем цели
			if res != "" {
				res += ", "
			}
			res += goalsString
		}
	}
	
	res += ". можно пройти - " + player.Location.GetEntrancesString()
	return res
}

// TakeItem allows the player to pick up an item
func (g *gameUseCase) TakeItem(itemName string) (string, error) {
	player := g.repo.GetPlayer()
	
	if player.Inventory == nil {
		return "", entity.NewGameError("некуда класть")
	}
	
	for _, object := range player.Location.Objects {
		if container, ok := object.(entity.Container); ok {
			item, err := container.Remove(itemName)
			if err == nil {
				player.Inventory.Add(item)
				return "предмет добавлен в инвентарь: " + item.GetName(), nil
			}
		}
	}
	
	return "", entity.NewGameError("нет такого")
}

// Equip allows the player to equip an item
func (g *gameUseCase) Equip(itemName string) (string, error) {
	player := g.repo.GetPlayer()
	
	for _, object := range player.Location.Objects {
		if container, ok := object.(entity.Container); ok {
			item, err := container.Remove(itemName)
			if err == nil {
				if equip, ok := item.(entity.Equippable); ok {
					return "вы надели: " + item.GetName(), equip.Equip(player)
				}
				return "", entity.NewGameError("ты как " + item.GetName() + " надеть собрался?")
			}
		}
	}
	
	return "", entity.NewGameError("нет такого")
}

// Use allows the player to use an item on an object
func (g *gameUseCase) Use(itemName, objName string) (string, error) {
	player := g.repo.GetPlayer()
	
	if player.Inventory == nil {
		return "", entity.NewGameError("нет предмета в инвентаре - " + itemName)
	}
	
	item, err := player.Inventory.GetItemByName(itemName)
	if err != nil {
		return "", entity.NewGameError("нет предмета в инвентаре - " + itemName)
	}
	
	for _, obj := range player.Location.Objects {
		if obj.GetName() == objName {
			if usable, ok := obj.(entity.Usable); ok {
				return usable.Use(item)
			}
			return "", entity.NewGameError("не к чему применить")
		}
	}
	
	return "", entity.NewGameError("не к чему применить")
}

// Restart resets the game state
func (g *gameUseCase) Restart() string {
	g.repo.InitializeWorld()
	return "игра перезапущена"
}

// onGoMessage returns a message when entering a room
func onGoMessage(room *entity.Room) (res string) {
	switch room.Name {
	case "кухня":
		res += "кухня, ничего интересного."
	case "коридор":
		res += "ничего интересного."
	case "комната":
		res += "ты в своей комнате."
	case "улица":
		return "на улице весна. можно пройти - домой"
	default:
		return "ты оказался в бескрайней бездне. тут ничего нет. ты обречён."
	}
	res += " можно пройти - " + room.GetEntrancesString()
	return
} 