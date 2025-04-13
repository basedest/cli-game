package repository

import (
	"../../domain/entity"
	"../../domain/repository"
)

// gameRepositoryImpl implements the GameRepository interface
type gameRepositoryImpl struct {
	rooms  map[string]*entity.Room
	player *entity.Player
}

// NewGameRepository creates a new game repository
func NewGameRepository() repository.GameRepository {
	return &gameRepositoryImpl{
		rooms:  make(map[string]*entity.Room),
		player: &entity.Player{},
	}
}

// GetRoom returns a room by name
func (g *gameRepositoryImpl) GetRoom(name string) (*entity.Room, error) {
	room, ok := g.rooms[name]
	if !ok {
		return nil, entity.NewGameError("комната не найдена")
	}
	return room, nil
}

// GetAllRooms returns all rooms
func (g *gameRepositoryImpl) GetAllRooms() map[string]*entity.Room {
	return g.rooms
}

// GetPlayer returns the player
func (g *gameRepositoryImpl) GetPlayer() *entity.Player {
	return g.player
}

// ResetPlayer resets the player state
func (g *gameRepositoryImpl) ResetPlayer() {
	g.player = &entity.Player{}
}

// InitializeWorld initializes the game world
func (g *gameRepositoryImpl) InitializeWorld() {
	// Create rooms
	kitchen := &entity.Room{Name: "кухня"}
	corridor := &entity.Room{Name: "коридор"}
	bedroom := &entity.Room{Name: "комната"}
	street := &entity.Room{Name: "улица"}
	
	// Store rooms in the map
	g.rooms = map[string]*entity.Room{
		"кухня":   kitchen,
		"коридор": corridor,
		"комната": bedroom,
		"улица":   street,
	}
	
	// Initialize player
	g.player.Location = kitchen
	g.player.Equipment = nil
	g.player.Inventory = nil
	g.player.Goals = []*entity.GoalData{
		{Goal: "собрать рюкзак", Completed: false},
		{Goal: "идти в универ", Completed: false},
	}
	
	// Connect rooms and add objects
	
	// Kitchen
	kitchenTable := &entity.Storage{Name: "стол", Items: []entity.Item{&entity.Thing{Name: "чай"}}}
	kitchen.Objects = []entity.Object{kitchenTable}
	kitchen.Entrances = []*entity.Entrance{{Room: corridor}}
	
	// Corridor
	corridorDoor := &entity.Door{Locked: true}
	corridor.Objects = []entity.Object{corridorDoor}
	corridor.Entrances = []*entity.Entrance{
		{Room: kitchen},
		{Room: bedroom},
		{Room: street, Door: corridorDoor},
	}
	
	// Bedroom
	roomTable := &entity.Storage{Name: "стол", Items: []entity.Item{&entity.Thing{Name: "ключи"}, &entity.Thing{Name: "конспекты"}}}
	roomChair := &entity.Storage{Name: "стул", Items: []entity.Item{&entity.Backpack{Thing: entity.Thing{Name: "рюкзак"}}}}
	bedroom.Objects = []entity.Object{roomTable, roomChair}
	bedroom.Entrances = []*entity.Entrance{{Room: corridor}}
	
	// Street
	street.Entrances = []*entity.Entrance{{Room: corridor, Door: corridorDoor}}
} 