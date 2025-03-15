package forms

import (
	"fmt"
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

	postData := url.Values{}
	postData.Add("b", "test")

	r.PostForm = postData

	length := r.PostForm.Get("b")

	actualLength := form.MinLength("b", 1, r)
	fmt.Println("actual length", actualLength, length)

}

func TestForm_Hass(t *testing.T) {
	r := httptest.NewRequest("POST", "/routename", nil)
	form := New(r.PostForm)

	postData := url.Values{}
	postData.Add("c", "test")

	r.PostForm = postData

	hasField := form.Has("c", r)
	fmt.Println("has the field", hasField)
}
