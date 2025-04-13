package entity

// Object represents anything in a room that can be interacted with
type Object interface {
	GetName() string
}

// Door represents a door that can be locked/unlocked
type Door struct {
	Locked bool
}

// Unlock unlocks the door
func (d *Door) Unlock() {
	d.Locked = false
}

// GetName returns the name of the door
func (d *Door) GetName() string {
	return "дверь"
}

// Use applies an item to the door
func (d *Door) Use(i Item) (string, error) {
	if i.GetName() == "ключи" {
		d.Unlock()
		return "дверь открыта", nil
	}
	return "", NewGameError("дверь открывают ключём если что")
} 