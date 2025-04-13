package entity

// Item represents something that can be picked up and used
type Item interface {
	GetName() string
}

// Thing is a basic implementation of the Item interface
type Thing struct {
	Name string
}

// GetName returns the name of the thing
func (t *Thing) GetName() string {
	return t.Name
}

// Usable represents an object that can be used with an item
type Usable interface {
	Use(Item) (string, error)
} 