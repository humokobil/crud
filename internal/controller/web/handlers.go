package web

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"otus/crud/v1/internal/entity"
	"strconv"

	"github.com/gorilla/mux"
)

func (h *handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		h.WriteResponse(w, 400, "Bad Request")
		return
	}
	defer r.Body.Close()
	var user entity.User 
	err = json.Unmarshal(bodyBytes, &user)
	if err != nil {
		log.Println(err)
		h.WriteResponse(w, 400, "Bad Request")
		return
	}
	newUser, err := (*h.usecase).CreateUser(user)
	if err != nil {
		log.Println(err)
		h.WriteResponse(w, 400, "Bad Request")
		return
	}
	respBytes, err := json.Marshal(newUser)
	if err != nil {
		log.Println(err)
		h.WriteResponse(w, 400, "Internal error")
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(respBytes)
	//h.WriteResponse(w, 200, "New user created")
}

func (h *handler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	userIdStr := mux.Vars(r)["userId"]
	userId, err := strconv.ParseInt(userIdStr, 10, 64)
	if err != nil {
		log.Println(err)
		h.WriteResponse(w, 400, "Bad Request")
		return
	}
	err = (*h.usecase).DeleteUser(userId)
	if err != nil {
		log.Println(err)
		h.WriteResponse(w, 400, "Internal error")
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *handler) FindUserById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	userIdStr := mux.Vars(r)["userId"]
	userId, err := strconv.ParseInt(userIdStr, 10, 64)
	if err != nil {
		log.Println(err)
		h.WriteResponse(w, 400, "Bad Request")
		return
	}
	u, err := (*h.usecase).FindUser(userId)
	if err != nil {
		log.Println(err)
		h.WriteResponse(w, 400, "User not found")
		return
	}
	respBytes, err := json.Marshal(u)
	if err != nil {
		log.Println(err)
		h.WriteResponse(w, 400, "Internal error")
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(respBytes)
}

func (h *handler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	userIdStr := mux.Vars(r)["userId"]
	userId, err := strconv.ParseInt(userIdStr, 10, 64)
	if err != nil {
		log.Println(err)
		h.WriteResponse(w, 400, "Bad Request")
		return
	}
	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		h.WriteResponse(w, 400, "Bad Request")
		return
	}
	defer r.Body.Close()
	var user entity.User 
	err = json.Unmarshal(bodyBytes, &user)
	if err != nil {
		log.Println(err)
		h.WriteResponse(w, 400, "Bad Request")
		return
	}
	newUser, err := (*h.usecase).UpdateUser(userId, user)
	if err != nil {
		log.Println(err)
		h.WriteResponse(w, 400, "Bad Request")
		return
	}
	respBytes, err := json.Marshal(newUser)
	if err != nil {
		log.Println(err)
		h.WriteResponse(w, 400, "Internal error")
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(respBytes)
}

func (h *handler) WriteResponse(w http.ResponseWriter, code int32, message string) {
	respBytes, err := json.Marshal(ModelError{Code: code, Message: message})
	if err == nil {
		w.Write(respBytes)
	}
	w.WriteHeader(int(code))
}
