package entity

// GameError represents an error in the game
type GameError struct {
	Message string
}

// Error returns the error message
func (e GameError) Error() string {
	return e.Message
}

// NewGameError creates a new game error with the given message
func NewGameError(message string) error {
	return GameError{Message: message}
} 