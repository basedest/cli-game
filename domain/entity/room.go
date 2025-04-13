package entity

// Room represents a location in the game
type Room struct {
	Name      string
	Objects   []Object
	Entrances []*Entrance
}

// Entrance represents a connection to another room
type Entrance struct {
	Room *Room
	Door *Door
}

// GetEntrancesString returns a comma-separated list of accessible room names
func (r *Room) GetEntrancesString() string {
	var entrances []string
	for _, entrance := range r.Entrances {
		entrances = append(entrances, entrance.Room.Name)
	}
	return joinStrings(entrances, ", ")
} 