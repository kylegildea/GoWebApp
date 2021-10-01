package handlers

import (
	"fmt"
	"net/http"

	"github.com/kylegildea/go-course/pkg/config"
	"github.com/kylegildea/go-course/pkg/models"
	"github.com/kylegildea/go-course/pkg/render"
)

// Repo the repository used by the handlers
var Repo *Repository

// Repository is the reposityory type
type Repository struct {
	App *config.AppConfig
}

//Creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

//New Handlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	fmt.Println(remoteIP)
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {

	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again"

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}
