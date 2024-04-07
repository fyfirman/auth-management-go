package pkg_test

import (
	"reflect"
	"testing"

	. "github.com/fyfirman/auth-management-go/pkg"
	"github.com/go-playground/validator/v10"

	ut "github.com/go-playground/universal-translator"
)

type mockFieldError struct {
	field, tag, param string
}

func (m mockFieldError) Tag() string {
	return m.tag
}

func (m mockFieldError) ActualTag() string {
	return m.tag
}

func (m mockFieldError) Namespace() string {
	return ""
}

func (m mockFieldError) Field() string {
	return m.field
}

func (m mockFieldError) StructNamespace() string {
	return ""
}

func (m mockFieldError) StructField() string {
	return ""
}

func (m mockFieldError) Value() interface{} {
	return nil
}

func (m mockFieldError) Param() string {
	return m.param
}

func (m mockFieldError) Kind() reflect.Kind {
	return reflect.Kind(0)
}

func (m mockFieldError) Type() reflect.Type {
	return nil
}

func (m mockFieldError) Translate(trans ut.Translator) string {
	return ""
}

func (m mockFieldError) Error() string {
	return "mock error"
}

func TestPrepareValidationErrors(t *testing.T) {
	type User struct {
		Email    string `validate:"required,email"`
		Password string `validate:"required,min=8"`
	}

	user := &User{Email: "invalid", Password: "short"} // Deliberately fail validation
	validate := validator.New()
	err := validate.Struct(user)
	if err != nil {
		validationErrors := PrepareValidationErrors(err)

		// Assuming you're expecting specific error messages for the provided input
		expected := []map[string]string{
			{"field": "Email", "message": "Email is not a valid email address"},
			{"field": "Password", "message": "Password must be at least 8 characters long"},
		}

		if !reflect.DeepEqual(validationErrors, expected) {
			t.Errorf("Expected %v, got %v", expected, validationErrors)
		}
	} else {
		t.Error("Expected validation errors, but got none")
	}
}

func TestFieldErrorMessage(t *testing.T) {
	cases := []struct {
		tag      string
		field    string
		param    string
		expected string
	}{
		{"required", "Name", "", "Name is required"},
		{"email", "Email", "", "Email is not a valid email address"},
		{"min", "Password", "8", "Password must be at least 8 characters long"},
	}

	for _, c := range cases {
		e := mockFieldError{
			tag:   c.tag,
			field: c.field,
			param: c.param,
		}

		message := FieldErrorMessage(e)
		if message != c.expected {
			t.Errorf("Expected %s, got %s", c.expected, message)
		}
	}
}
