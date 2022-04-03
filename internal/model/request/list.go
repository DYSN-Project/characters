package request

type ListForm struct {
	Limit  int
	Offset int
}

func NewListForm(limit, offset int) *ListForm {
	return &ListForm{limit, offset}
}
