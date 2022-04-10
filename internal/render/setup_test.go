package render

import (
	"encoding/gob"
	"github.com/alexedwards/scs/v2"
	"github.com/sw90lee/go-web/internal/config"
	"github.com/sw90lee/go-web/internal/handlers"
	"github.com/sw90lee/go-web/internal/models"
	"log"
	"net/http"
	"os"
	"testing"
	"time"
)

var session *scs.SessionManager
var testApp config.AppConfig

func TestMain(m *testing.M) {
	gob.Register(models.Reservation{})

	// change this to true when in production
	testApp.InProduction = false

	// set up the session
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = false

	testApp.Session = session

	app = &testApp

	tc, err := CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(app)
	handlers.NewHandlers(repo)
	NewTemplates(app)

	os.Exit(m.Run())
}
