package handler

import (
	"database_implementation/internal"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

type ProductDefault struct {
	sv internal.ProductService
}

func NewProductDefault(sv internal.ProductService) *ProductDefault {
	return &ProductDefault{
		sv: sv,
	}
}

type BodyResponse struct {
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func (h *ProductDefault) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		products, err := h.sv.GetAll()
		if err != nil {
			code := http.StatusInternalServerError
			body := BodyResponse{
				Message: "An error occurred",
				Data:    nil,
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(code)
			json.NewEncoder(w).Encode(body)
			return
		}

		code := http.StatusOK
		body := BodyResponse{
			Message: "Products found",
			Data:    products,
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(code)
		json.NewEncoder(w).Encode(body)
	}
}

func (h *ProductDefault) GetOne() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			code := http.StatusBadRequest
			body := BodyResponse{
				Message: "Invalid ID. ID must be an integer",
				Data:    nil,
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(code)
			json.NewEncoder(w).Encode(body)
			return
		}

		product, err := h.sv.GetOne(id)
		if err != nil {
			code := http.StatusNotFound
			body := BodyResponse{
				Message: "Product not found",
				Data:    nil,
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(code)
			json.NewEncoder(w).Encode(body)
			return
		}

		code := http.StatusOK
		body := BodyResponse{
			Message: "Product found.",
			Data:    product,
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(code)
		json.NewEncoder(w).Encode(body)
	}
}
