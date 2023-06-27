package handlers

import (
	"net/http"
	"web3/models"
	"web3/pkg/config"
	"web3/pkg/render"
)

type Repository struct {
	App *config.AppConfig
}

var Repo *Repository

func NewRepo(ac *config.AppConfig) *Repository {
	return &Repository{
		App: ac,
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) HomeHandler(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl", &models.PageData{})
}

func (m *Repository) AboutHandler(w http.ResponseWriter, r *http.Request) {
	strMap := make(map[string]string)
	strMap["title"] = "About Us"
	strMap["intro"] = "This page where we talk about ourself. We love talking about ourself."
	render.RenderTemplate(w, "about.page.tmpl", &models.PageData{StrMap: strMap})
}
