package entity

// Player represents the player character in the game
type Player struct {
	Inventory *Backpack
	Equipment []Equippable
	Location  *Room
	Goals     []*GoalData
}

// GoalData represents a player goal
type GoalData struct {
	Goal      string
	Completed bool
}

// FindGoalIndex returns the index of a goal by name, or -1 if not found
func FindGoalIndex(array []*GoalData, value string) int {
	for index, item := range array {
		if item.Goal == value {
			return index
		}
	}
	return -1
}

// GetGoalsString returns a string describing the player's uncompleted goals
func (p *Player) GetGoalsString() (res string) {
	res += "надо "
	var goals []string
	for _, data := range p.Goals {
		if !data.Completed {
			goals = append(goals, data.Goal)
		}
	}
	if len(goals) == 0 {
		return "ты не знаешь, что дальше делать со своей жизнью"
	}
	res += joinStrings(goals, " и ")
	return
} 