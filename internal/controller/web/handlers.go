package web

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"otus/crud/v1/internal/entity"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

func (h *handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var status int = 400
	var timeStarted = time.Now()
	defer func() {
		h.reqCounter.WithLabelValues(fmt.Sprintf("%d",status), "create_user").Inc()
		h.reqDuration.WithLabelValues("create_user").Observe(float64(time.Since(timeStarted))/1e+9)
	}()

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		h.WriteResponse(w, status, "Bad Request")
		return
	}
	defer r.Body.Close()
	var user entity.User 
	err = json.Unmarshal(bodyBytes, &user)
	if err != nil {
		log.Println(err)
		h.WriteResponse(w, status, "Bad Request")
		return
	}
	newUser, err := (*h.usecase).CreateUser(user)
	if err != nil {
		log.Println(err)
		status = 500
		h.WriteResponse(w, status, "Bad Request")
		return
	}
	respBytes, err := json.Marshal(newUser)
	if err != nil {
		log.Println(err)
		status = 500
		h.WriteResponse(w, status, "Internal error")
		return
	}
	status = http.StatusOK
	w.WriteHeader(status)
	w.Write(respBytes)
	//h.WriteResponse(w, 200, "New user created")
}

func (h *handler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	var status int = 400
	var timeStarted = time.Now()
	defer func() {
		h.reqCounter.WithLabelValues(fmt.Sprintf("%d",status), "delete_user").Inc()
		h.reqDuration.WithLabelValues("delete_user").Observe(float64(time.Since(timeStarted))/1e+9)
	}()
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	userIdStr := mux.Vars(r)["userId"]
	userId, err := strconv.ParseInt(userIdStr, 10, 64)
	if err != nil {
		log.Println(err)
		h.WriteResponse(w, status, "Bad Request")
		return
	}
	err = (*h.usecase).DeleteUser(userId)
	if err != nil {
		log.Println(err)
		status = 500
		h.WriteResponse(w, status, "Internal error")
		return
	}
	status = http.StatusNoContent
	w.WriteHeader(status)
}

func (h *handler) FindUserById(w http.ResponseWriter, r *http.Request) {
	var status int = 400
	var timeStarted = time.Now()
	defer func() {
		h.reqCounter.WithLabelValues(fmt.Sprintf("%d",status), "find_user").Inc()
		h.reqDuration.WithLabelValues("find_user").Observe(float64(time.Since(timeStarted))/1e+9)
	}()
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	userIdStr := mux.Vars(r)["userId"]
	userId, err := strconv.ParseInt(userIdStr, 10, 64)
	if err != nil {
		log.Println(err)
		h.WriteResponse(w, status, "Bad Request")
		return
	}
	u, err := (*h.usecase).FindUser(userId)
	if err != nil {
		log.Println(err)
		status = 500
		h.WriteResponse(w, status, "Internal error")
		return
	}

	if u == nil {
		status = 404
		h.WriteResponse(w, status, "User not found")
		return
	}
	respBytes, err := json.Marshal(u)
	if err != nil {
		log.Println(err)
		status = 500
		h.WriteResponse(w, status, "Internal error")
		return
	}
	status = http.StatusOK
	w.WriteHeader(status)
	w.Write(respBytes)
}

func (h *handler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	var status int = 400
	var timeStarted = time.Now()
	defer func() {
		h.reqCounter.WithLabelValues(fmt.Sprintf("%d",status), "update_user").Inc()
		h.reqDuration.WithLabelValues("update_user").Observe(float64(time.Since(timeStarted)/time.Second))
	}()
	
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	userIdStr := mux.Vars(r)["userId"]
	userId, err := strconv.ParseInt(userIdStr, 10, 64)
	if err != nil {
		log.Println(err)
		h.WriteResponse(w, status, "Bad Request")
		return
	}
	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		status = 500
		h.WriteResponse(w, status, "Internal error")
		return
	}
	defer r.Body.Close()
	var user entity.User 
	err = json.Unmarshal(bodyBytes, &user)
	if err != nil {
		log.Println(err)
		h.WriteResponse(w, status, "Bad Request")
		return
	}
	newUser, err := (*h.usecase).UpdateUser(userId, user)
	if err != nil {
		log.Println(err)
		status = 500
		h.WriteResponse(w, status, "Bad Request")
		return
	}
	respBytes, err := json.Marshal(newUser)
	if err != nil {
		log.Println(err)
		h.WriteResponse(w, status, "Internal error")
		return
	}
	status = http.StatusOK
	w.WriteHeader(status)
	w.Write(respBytes)
}

func (h *handler) WriteResponse(w http.ResponseWriter, code int, message string) {
	respBytes, err := json.Marshal(ModelError{Code: code, Message: message})
	w.WriteHeader(int(code))
	if err == nil {
		w.Write(respBytes)
	}
}
