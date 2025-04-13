package entity

// Equippable represents an item that can be equipped
type Equippable interface {
	Item
	Equip(*Player) error
}

// Backpack is an implementation of both Container and Equippable
type Backpack struct {
	Thing          // Embedding Thing to inherit GetName
	Items  []Item
	Holder *Player
}

// Add adds an item to the backpack
func (b *Backpack) Add(item Item) {
	b.Items = append(b.Items, item)
}

// Remove removes an item from the backpack by name
func (b *Backpack) Remove(i string) (Item, error) {
	for index, item := range b.Items {
		if item.GetName() == i {
			b.Items = append(b.Items[:index], b.Items[index+1:]...)
			return item, nil
		}
	}
	return nil, NewGameError("нет такого")
}

// GetAll returns all items in the backpack
func (b *Backpack) GetAll() []Item {
	return b.Items
}

// IsEmpty returns true if the backpack is empty
func (b *Backpack) IsEmpty() bool {
	return len(b.Items) == 0
}

// String returns a string representation of the backpack contents
func (b Backpack) String() string {
	var items []string
	for _, item := range b.Items {
		items = append(items, item.GetName())
	}
	return "в рюкзаке: " + joinStrings(items, ", ")
}

// GetItemByName returns an item from the backpack by name
func (b *Backpack) GetItemByName(i string) (Item, error) {
	for _, item := range b.Items {
		if item.GetName() == i {
			return item, nil
		}
	}
	return nil, NewGameError("нет такого")
}

// Equip equips the backpack to a player
func (b *Backpack) Equip(p *Player) error {
	if b.Holder != nil {
		return NewGameError("ты уже надел рюкзак")
	}
	p.Inventory = b
	p.Equipment = append(p.Equipment, b)
	idx := FindGoalIndex(p.Goals, "собрать рюкзак")
	if idx != -1 {
		p.Goals[idx].Completed = true
	}
	b.Holder = p
	return nil
} 