package model

type Character struct {
	ID          string `bson:"_id"`
	Name        string `bson:"name,omitempty",json:"name"`
	Description string `bson:"description,omitempty",json:"description"`
	Image       string `bson:"image"`
	Status      int    `bson:"status,omitempty",json:"status"`
}

func NewCharacter(name, description, image string, status int) *Character {
	return &Character{
		Name:        name,
		Description: description,
		Image:       image,
		Status:      status,
	}
}

func NewCharacterIngot() *Character {
	return &Character{}
}
