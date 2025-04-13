package repository

import "../entity"

// GameRepository defines methods for accessing game state
type GameRepository interface {
	// Room operations
	GetRoom(name string) (*entity.Room, error)
	GetAllRooms() map[string]*entity.Room
	
	// Player operations
	GetPlayer() *entity.Player
	ResetPlayer()
	
	// World initialization
	InitializeWorld()
} 