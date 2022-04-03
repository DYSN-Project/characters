package request

type DeleteForm struct {
	ID  string

}

func NewDeleteForm(id string) *DeleteForm {
	return &DeleteForm{id}
}
