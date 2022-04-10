package render

import (
	"github.com/sw90lee/go-web/internal/models"
	"net/http"
	"testing"
)

func TestAddDefaultData(t *testing.T) {
	var td models.TemplateData
	r, err := getSesstion()
	if err != nil {
		t.Error(err)
	}

	session.Put(r.Context(), "flash", "123")

	result := AddDefaultData(&td, r)

	if result.Flash != "123" {
		t.Error("flash value of 123 not found")
	}
}

func getSesstion() (*http.Request, error) {

	r, err := http.NewRequest("GET", "/some-url", nil)
	if err != nil {
		return nil, err
	}

	ctx := r.Context()
	ctx, _ = session.Load(ctx, r.Header.Get("X-session"))
	r = r.WithContext(ctx)

	return r, nil
}
