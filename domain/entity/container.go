package entity

// Container represents anything that can store items
type Container interface {
	Object
	Add(Item)
	Remove(string) (Item, error)
	GetAll() []Item
	IsEmpty() bool
	GetItemByName(string) (Item, error)
	String() string
}

// Storage is a basic implementation of the Container interface
type Storage struct {
	Name  string
	Items []Item
}

// GetName returns the name of the storage
func (s Storage) GetName() string {
	return s.Name
}

// Add adds an item to the storage
func (s *Storage) Add(i Item) {
	s.Items = append(s.Items, i)
}

// Remove removes an item from the storage by name
func (s *Storage) Remove(i string) (Item, error) {
	for index, item := range s.Items {
		if item.GetName() == i {
			s.Items = append(s.Items[:index], s.Items[index+1:]...)
			return item, nil
		}
	}
	return nil, NewGameError("нет такого")
}

// GetAll returns all items in the storage
func (s Storage) GetAll() []Item {
	return s.Items
}

// IsEmpty returns true if the storage is empty
func (s *Storage) IsEmpty() bool {
	return len(s.Items) == 0
}

// String returns a string representation of the storage contents
func (s Storage) String() string {
	var items []string
	for _, item := range s.Items {
		items = append(items, item.GetName())
	}
	return "на " + s.GetName() + "е: " + joinStrings(items, ", ")
}

// GetItemByName returns an item by name
func (s Storage) GetItemByName(i string) (Item, error) {
	for _, item := range s.Items {
		if item.GetName() == i {
			return item, nil
		}
	}
	return nil, NewGameError("нет такого")
} 