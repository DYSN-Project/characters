package validation

import (
	"dysn/character/internal/helper"
	"dysn/character/internal/model/request"
	validation "github.com/go-ozzo/ozzo-validation"
	"regexp"
)

type ValidationInterface interface {
	ValidateList(list *request.ListForm) error
	ValidateCreate(create *request.CreateForm) error
	ValidateUpdate(create *request.UpdateForm) error
	ValidateDelete(create *request.DeleteForm) error
}

type Validation struct{}

var stringRegexp = regexp.MustCompile("[A-Za-zА-Яа-я]")

func NewValidation() *Validation {
	return &Validation{}
}

func (v *Validation) ValidateList(frm *request.ListForm) error {
	return validation.ValidateStruct(frm,
		validation.Field(&frm.Limit, validation.Min(1), validation.Max(999999)),
		validation.Field(&frm.Offset, validation.Min(1), validation.Max(999999)),
	)
}

func (v *Validation) ValidateCreate(frm *request.CreateForm) error {
	return validation.ValidateStruct(frm,
		validation.Field(&frm.Name, validation.Required, validation.Length(2, 255), validation.Match(stringRegexp)),
		validation.Field(&frm.Description, validation.Required, validation.Length(2, 1000)),
	)
}

func (v *Validation) ValidateUpdate(frm *request.UpdateForm) error {
	return validation.ValidateStruct(frm,
		validation.Field(&frm.Name, validation.Length(2, 255), validation.Match(stringRegexp)),
		validation.Field(&frm.Description, validation.Length(2, 1000)),
		validation.Field(&frm.Status, validation.In(helper.StatusActive, helper.StatusNotActive)),
	)
}

func (v *Validation) ValidateDelete(frm *request.DeleteForm) error {
	return validation.ValidateStruct(frm,
		validation.Field(&frm.ID, validation.Required, validation.Length(2, 255)),
	)
}
