package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"unicode/utf8"

	"oscar.dev/github.com/lets_go/internal/models"
)

type snippetCreateForm struct {
	Title       string
	Content     string
	Expires     int
	FieldErrors map[string]string
}

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	snippets, err := app.snippets.Latest()
	if err != nil {
		app.serverError(w, r, err)
		return
	}
	data := app.newTemplateData(r)
	data.Snippets = snippets
	app.render(w, r, http.StatusOK, "home.tmpl.html", data)
}

func (app *application) getSnippetCreate(w http.ResponseWriter, r *http.Request) {
	data := app.newTemplateData(r)
	app.render(w, r, http.StatusOK, "create.tmpl.html", data)
}
func (app *application) postSnippetCreate(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	expires, err := strconv.Atoi(r.PostForm.Get("expires"))
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}
	form := snippetCreateForm{
		Title:       r.PostForm.Get("title"),
		Content:     r.PostForm.Get("content"),
		Expires:     expires,
		FieldErrors: map[string]string{},
	}

	fieldErrors := make(map[string]string)
	if strings.TrimSpace(form.Title) == "" {
		fieldErrors["title"] = "This can not be blank"
	} else if utf8.RuneCountInString(form.Title) > 100 {
		fieldErrors["title"] = "This field cannot be of more than 100 characters long"
	}

	if strings.TrimSpace(form.Content) == "" {
		fieldErrors["content"] = "This field cannot be blank"
	}

	if form.Expires != 1 && form.Expires != 7 && form.Expires != 365 {
		fieldErrors["expires"] = "This field must equal 1, 7 or 365"
	}
	if len(form.FieldErrors) > 0 {
		data := app.newTemplateData(r)
		data.Form = form
		app.render(w, r, http.StatusUnprocessableEntity, "create.tmpl.html", data)
		return
	}
	id, err := app.snippets.Insert(form.Title, form.Content, expires)
	if err != nil {
		app.serverError(w, r, err)
	}
	http.Redirect(w, r, fmt.Sprintf("/snippet/view/%d", id), http.StatusSeeOther)
}
func (app *application) snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	snippet, err := app.snippets.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			http.NotFound(w, r)
		} else {
			app.serverError(w, r, err)
		}
		return
	}
	data := app.newTemplateData(r)
	data.Snippet = snippet
	app.render(w, r, http.StatusOK, "view.tmpl.html", data)
}
