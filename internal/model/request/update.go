package request

type UpdateForm struct {
	Name        string
	Description string
	Status      int
}

func NewUpdateForm(name, description string,status int) *UpdateForm {
	return &UpdateForm{name,description,status}
}
