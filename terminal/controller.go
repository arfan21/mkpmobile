package terminal

import (
	"encoding/json"
	"log"
	"net/http"
)

type Controller struct {
	service Service
}

func NewController(service Service) *Controller {
	return &Controller{service}
}

func (c *Controller) RegisterTerminal(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	req := &RegisterTerminalRequest{}

	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"message": "error",
			"errors":  err.Error(),
		})
		return
	}

	err = c.service.RegisterTerminal(r.Context(), *req)

	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"message": err.Error(),
		})
		return
	}

	w.WriteHeader(http.StatusCreated)
}
