package render

import (
	"net/http"
	"testing"

	"github.com/tsawler/bookings-app/internal/models"
)

func TestDefaultData(t *testing.T) {
	var td models.TemplateData
	r, err := getSession()

	if err != nil {
		t.Error(err)
	}

	session.Put(r.Context(), "flash", "123")

	result := AddDefaultData(&td, r)

	if result.Flash != "123" {
		t.Error("flash value not found")
	}
}

func TestTemplate(t *testing.T) {
	pathToTemplates = "./../../templates"
	tc, err := CreateTemplateCache()
	if err != nil {
		t.Error(err)
	}

	app.TemplateCache = tc

	r, err := getSession()
	if err != nil {
		t.Error(err)
	}

	var ww myWriter

	err = Template(&ww, r, "home.page.tmpl", &models.TemplateData{})
	if err != nil {
		t.Error("error writing template to brwoser")
	}

	err = Template(&ww, r, "non-existent.page.tmpl", &models.TemplateData{})
	if err == nil {
		t.Error("render template that doesn't exist")
	}
}

func getSession() (*http.Request, error) {
	r, err := http.NewRequest("GET", "/some-url", nil)

	if err != nil {
		return nil, err
	}
	ctx := r.Context()
	ctx, _ = session.Load(ctx, r.Header.Get("X-Session"))
	r = r.WithContext(ctx)

	return r, nil
}

func TestNewRenderer(t *testing.T) {
	NewRenderer(app)
}

func TestCreateTemplateCache(t *testing.T) {
	pathToTemplates = "./../../tempalates"
	_, err := CreateTemplateCache()

	if err != nil {
		t.Error(err)
	}
}
