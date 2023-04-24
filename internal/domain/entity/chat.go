package entity

type ChatConfig struct {
	Model *Model
}

type Chat struct {
	ID                   string
	UserID               string
	InitialSystemMessage *Message
	Messages             []*Message
	ErasedMessages       []*Message
	Status               string
	TokenUsage           int
	Config               *ChatConfig
}

func NewModel(name string, maxTokens int) *Model {
	return &Model{
		Name:      name,
		MaxTokens: maxTokens,
	}
}

func (m *Model) GetMaxTokens() int {
	return m.MaxTokens
}

func (m *Model) GetModelName() string {
	return m.Name
}