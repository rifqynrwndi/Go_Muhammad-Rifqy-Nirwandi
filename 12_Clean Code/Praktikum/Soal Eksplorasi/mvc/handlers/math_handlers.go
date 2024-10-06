package handlers

import (
    "encoding/json"
    "net/http"
    "strconv"
    "go-math-api/mvc/services"
)

type MathHandler struct {
    service *services.MathService
}

func NewMathHandler() *MathHandler {
    return &MathHandler{
        service: services.NewMathService(),
    }
}

func (h *MathHandler) Add(w http.ResponseWriter, r *http.Request) {
    x, y := parseParams(r)
    result := h.service.Add(x, y)
    json.NewEncoder(w).Encode(map[string]float64{"result": result})
}

func (h *MathHandler) Subtract(w http.ResponseWriter, r *http.Request) {
    x, y := parseParams(r)
    result := h.service.Subtract(x, y)
    json.NewEncoder(w).Encode(map[string]float64{"result": result})
}

func (h *MathHandler) Multiply(w http.ResponseWriter, r *http.Request) { // Add Multiply method
    x, y := parseParams(r)
    result := h.service.Multiply(x, y)
    json.NewEncoder(w).Encode(map[string]float64{"result": result})
}

func (h *MathHandler) Divide(w http.ResponseWriter, r *http.Request) {
    x, y := parseParams(r)
    result := h.service.Divide(x, y)
    json.NewEncoder(w).Encode(map[string]float64{"result": result})
}

// Implement a helper function to parse query parameters
func parseParams(r *http.Request) (float64, float64) {
    query := r.URL.Query()
    x, _ := strconv.ParseFloat(query.Get("x"), 64)
    y, _ := strconv.ParseFloat(query.Get("y"), 64)
    return x, y
}
