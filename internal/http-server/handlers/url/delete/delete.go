package delete

import (
	"log/slog"
	"url_shortener/internal/storage/sqlite"
)

type URLDeleter interface {
	DeleteURL(id string) error
}
type Handler struct {
	log     *slog.Logger
	storage sqlite.Storage
}

var urlDeleter URLDeleter

// TODO: implement delete handler
