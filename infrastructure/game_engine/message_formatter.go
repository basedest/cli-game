package game_engine

// MessageFormatter provides consistent formatting for game messages
type MessageFormatter struct{}

// NewMessageFormatter creates a new message formatter
func NewMessageFormatter() *MessageFormatter {
	return &MessageFormatter{}
}

// FormatRoomDescription formats a room description message
func (m *MessageFormatter) FormatRoomDescription(roomName, description, exits string) string {
	return description + " можно пройти - " + exits
}

// FormatInventoryItem formats an inventory item message
func (m *MessageFormatter) FormatInventoryItem(itemName string) string {
	return "предмет добавлен в инвентарь: " + itemName
}

// FormatEquippedItem formats an equipped item message
func (m *MessageFormatter) FormatEquippedItem(itemName string) string {
	return "вы надели: " + itemName
}

// FormatErrorMessage formats an error message
func (m *MessageFormatter) FormatErrorMessage(message string) string {
	return message
} 