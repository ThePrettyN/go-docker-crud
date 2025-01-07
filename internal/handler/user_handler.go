package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"go-docker-crud/internal/repository"
	"go-docker-crud/internal/service"
	"github.com/go-playground/validator/v10"
)

type UserHandler struct {
	Service service.UserService
}

func NewUserHandler(service service.UserService) *UserHandler {
	return &UserHandler{Service: service}
}

// CreateUser は POST /users に対して新しいユーザーを作成するハンドラーです
func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user repository.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "無効な入力です", http.StatusBadRequest)
		return
	}

	// バリデーション
	validate := validator.New()
	if err := validate.Struct(user); err != nil {
		http.Error(w, "無効なデータです", http.StatusBadRequest)
		return
	}

	createdUser, err := h.Service.CreateUser(user)
	if err != nil {
		http.Error(w, "ユーザーの作成中にエラーが発生しました", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdUser)
}

// GetAllUsers は GET /users に対してすべてのユーザーを取得するハンドラーです
func (h *UserHandler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.Service.GetAllUsers()
	if err != nil {
		http.Error(w, "ユーザーの取得中にエラーが発生しました", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(users)
}

// GetUser は GET /users/{id} に対して指定されたIDのユーザーを取得するハンドラーです
func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	user, err := h.Service.GetUserByID(id)
	if err != nil {
		http.Error(w, "ユーザーが見つかりません", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(user)
}

// UpdateUser は PUT /users/{id} に対して指定されたIDのユーザー情報を更新するハンドラーです
func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	var user repository.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "無効な入力です", http.StatusBadRequest)
		return
	}

	updatedUser, err := h.Service.UpdateUser(id, user)
	if err != nil {
		http.Error(w, "ユーザーの更新中にエラーが発生しました", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(updatedUser)
}

// DeleteUser は DELETE /users/{id} に対して指定されたIDのユーザーを削除するハンドラーです
func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	err := h.Service.DeleteUser(id)
	if err != nil {
		http.Error(w, "ユーザーの削除中にエラーが発生しました", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
