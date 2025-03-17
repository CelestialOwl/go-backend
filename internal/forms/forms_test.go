package forms

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestForm_Valid(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.Form)

	isValid := form.Valid()
	if !isValid {
		t.Error("got invalid when shouldve been valid")
	}
}

func TestForm_Required(t *testing.T) {
	r := httptest.NewRequest("POST", "/routename", nil)
	form := New(r.PostForm)

	form.Required("a", "b", "c")
	if form.Valid() {
		t.Error("forms shows valid when required field is missing")
	}

	postedData := url.Values{}
	postedData.Add("a", "a")
	postedData.Add("b", "b")
	postedData.Add("c", "c")

	r, _ = http.NewRequest("POST", "/routename", nil)

	r.PostForm = postedData

	form = New(r.PostForm)
	form.Required("a", "b", "c")
	if !form.Valid() {
		t.Error("shows does not have required field when it does")
	}
}

func TestForm_MinLength(t *testing.T) {
	r := httptest.NewRequest("POST", "/routename", nil)
	form := New(r.PostForm)

	form.MinLength("x", 10)
	if form.Valid() {
		t.Error("form shows minlength for nonexistent field")
	}

	isError := form.Errors.Get("x")
	if isError == "" {
		t.Error("should throw an error but didn't get one")
	}

	postedValues := url.Values{}
	postedValues.Add("some", "milength")
	form = New(postedValues)

	form.MinLength("some", 100)
	if form.Valid() {
		t.Error("shows minlength of 100 met when data is short")
	}

	postedValues = url.Values{}
	postedValues.Add("c", "twenty")
	form = New(postedValues)
	form.MinLength("c", 5)
	if !form.Valid() {
		t.Error("shows error when minLength conditions are met, expected")
	}

	isError = form.Errors.Get("c")
	if isError != "" {
		t.Error("getting error for valid field type")
	}
}

func TestForm_Has(t *testing.T) {
	r := httptest.NewRequest("POST", "/routename", nil)
	form := New(r.PostForm)

	hasField := form.Has("somevalue")

	if hasField {
		t.Error("forms shows has field when it shouldn't")
	}

	postData := url.Values{}
	postData.Add("c", "test")
	form = New(postData)

	hasField = form.Has("c")

	if !hasField {
		t.Error("shows form doesn't have a field when it doesn't")
	}
}

func TestForm_IsEmail(t *testing.T) {
	postData := url.Values{}
	form := New(postData)

	form.IsEmail("x")
	if form.Valid() {
		t.Error("no email provided but accepted")
	}

	postData = url.Values{}
	postData.Add("email", "user@gmail.co")
	form = New(postData)
	form.IsEmail("email")
	if !form.Valid() {
		t.Error("wrong email provided but accepted")
	}

	postData = url.Values{}
	postData.Add("email", "user@co")
	form = New(postData)
	form.IsEmail("email")
	if form.Valid() {
		t.Error("wrong email provided but accepted")
	}
}
