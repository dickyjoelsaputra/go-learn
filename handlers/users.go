package handlers

import (
	"encoding/json"
	dto "go_learn/dto/result"
	usersdto "go_learn/dto/users"
	"go_learn/models"
	"go_learn/repositories"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type handler struct {
  UserRepository repositories.UserRepository
}

func HandlerUser(UserRepository repositories.UserRepository) *handler {
  return &handler{UserRepository}
} 

func (h *handler) FindUsers(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")

  users, err := h.UserRepository.FindUsers()
  if err != nil {
    w.WriteHeader(http.StatusInternalServerError)
    response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
    json.NewEncoder(w).Encode(response)
  }

  w.WriteHeader(http.StatusOK)
  response := dto.SuccessResult{Code: http.StatusOK, Data: users}
  json.NewEncoder(w).Encode(response)
}

func (h *handler) GetUser(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")

  id, _ := strconv.Atoi(mux.Vars(r)["id"])

  user, err := h.UserRepository.GetUser(id)
  if err != nil {
    w.WriteHeader(http.StatusBadRequest)
    response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
    json.NewEncoder(w).Encode(response)
    return
  }

  w.WriteHeader(http.StatusOK)
  response := dto.SuccessResult{Code: http.StatusOK, Data: convertResponse(user)}
  json.NewEncoder(w).Encode(response)
}

func (h *handler) CreateUser(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")

  request := new(usersdto.CreateUserRequest)
  if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
    w.WriteHeader(http.StatusBadRequest)
    response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
    json.NewEncoder(w).Encode(response)
    return
  }

  validation := validator.New()
  err := validation.Struct(request)
  if err != nil {
    w.WriteHeader(http.StatusBadRequest)
    response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
    json.NewEncoder(w).Encode(response)
    return
  }

  // data form pattern submit to pattern entity db user
  user := models.User{
    Name:     request.Name,
    Email:    request.Email,
    Password: request.Password,
  }

  data, err := h.UserRepository.CreateUser(user)
  if err != nil {
    w.WriteHeader(http.StatusInternalServerError)
    json.NewEncoder(w).Encode(err.Error())
  }

  w.WriteHeader(http.StatusOK)
  response := dto.SuccessResult{Code: http.StatusOK, Data: convertResponse(data)}
  json.NewEncoder(w).Encode(response)
}

// Write this code
func (h *handler) UpdateUser(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")

  request := new(usersdto.UpdateUserRequest) //take pattern data submission
  if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
    w.WriteHeader(http.StatusBadRequest)
    response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
    json.NewEncoder(w).Encode(response)
    return
  }

  id, _ := strconv.Atoi(mux.Vars(r)["id"])

	user, err := h.UserRepository.GetUser(int(id))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

  if request.Name != "" {
    user.Name = request.Name
  }

  if request.Email != "" {
    user.Email = request.Email
  }

  if request.Password != "" {
    user.Password = request.Password
  }

  data, err := h.UserRepository.UpdateUser(user,id)
  if err != nil {
    w.WriteHeader(http.StatusInternalServerError)
    response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
    json.NewEncoder(w).Encode(response)
    return
  }

  w.WriteHeader(http.StatusOK)
  response := dto.SuccessResult{Code: http.StatusOK, Data: convertResponse(data)}
  json.NewEncoder(w).Encode(response)
}

func (h *handler) DeleteUser(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")

  id, _ := strconv.Atoi(mux.Vars(r)["id"])

  user, err := h.UserRepository.GetUser(id)
  if err != nil {
    w.WriteHeader(http.StatusBadRequest)
    response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
    json.NewEncoder(w).Encode(response)
    return
  }

  data, err := h.UserRepository.DeleteUser(user,id)
  if err != nil {
    w.WriteHeader(http.StatusInternalServerError)
    response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
    json.NewEncoder(w).Encode(response)
    return
  }

  w.WriteHeader(http.StatusOK)
  response := dto.SuccessResult{Code: http.StatusOK, Data: convertResponse(data)}
  json.NewEncoder(w).Encode(response)
}

  func convertResponse(u models.User) usersdto.UserResponse {
  return usersdto.UserResponse{
    ID:       u.ID,
    Name:     u.Name,
    Email:    u.Email,
    Password: u.Password,
  }
}