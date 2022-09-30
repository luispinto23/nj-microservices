package handlers

import (
	"log"
	"net/http"

	"github.com/luispinto23/go-micro/data"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	lp := data.GetProducts()

	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Can't unmarshal json", http.StatusInternalServerError)
	}
}
