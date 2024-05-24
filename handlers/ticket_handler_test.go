package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/osag1e/go-and-mongo/db"
	"github.com/osag1e/go-and-mongo/model"
)

func TestPostMovie(t *testing.T) {
	tdb := db.Setup(t)
	defer tdb.Teardown(t)

	r := chi.NewRouter()
	movieHandler := NewMovieTicketHandler(tdb.Movie)
	r.Post("/movie/ticket", movieHandler.HandlePostMovie)

	params := model.MovieTicket{
		Title: "Deadpool",
		Price: 27.99,
	}
	b, _ := json.Marshal(params)
	req := httptest.NewRequest("POST", "/movie/ticket", bytes.NewReader(b))
	req.Header.Add("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status code %d but got %d", http.StatusOK, w.Code)
	}

	var ticket model.MovieTicket
	err := json.NewDecoder(w.Body).Decode(&ticket)
	if err != nil {
		t.Error(err)
	}

	if ticket.Title != params.Title || ticket.Price != params.Price {
		t.Errorf("expected ticket %+v but got %+v", params, ticket)
	}
}
