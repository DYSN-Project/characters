package request

type CreateForm struct {
	Name        string
	Description string
}

func NewCreateForm(name, description string) *CreateForm {
	return &CreateForm{name, description}
}
