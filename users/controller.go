package users

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

func (c *Controller) RegisterUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	req := &RegisterUserRequest{}

	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		log.Println("error decode json", err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"message": "error",
			"errors":  err.Error(),
		})
		return
	}

	err = c.service.RegisterUser(r.Context(), *req)

	if err != nil {
		log.Println("error service", err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"message": "error",
			"errors":  err.Error(),
		})
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (c *Controller) Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	req := &LoginUserRequest{}

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

	res, err := c.service.Login(r.Context(), *req)

	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"message": "error",
			"errors":  err.Error(),
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "success",
		"data":    res,
	})
}
