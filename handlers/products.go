package handlers

import (
	"encoding/json"
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

	encodedProducts, err := json.Marshal(lp)
	if err != nil {
		http.Error(rw, "Can't unmarshal json", http.StatusInternalServerError)
	}

	rw.Write(encodedProducts)
}
