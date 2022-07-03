package web

import (
	"fmt"
	"net/http"
	"otus/crud/v1/internal/entity"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

type IHandler interface {
	CreateUser(w http.ResponseWriter, r *http.Request) 
	DeleteUser(w http.ResponseWriter, r *http.Request) 
	FindUserById(w http.ResponseWriter, r *http.Request)
	UpdateUser(w http.ResponseWriter, r *http.Request)
	GetRoutes() *mux.Router
}

type handler struct {
	usecase *entity.IUsecase
}

func New(usecase *entity.IUsecase) IHandler {
	return &handler{usecase: usecase}
}
func (h *handler) GetRoutes() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	routes := h.getRoutes()

	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)
		fmt.Println("r: ", route.Pattern)
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
		url, _ := router.Get(route.Name).URL("userId", "42")
		fmt.Println("url: ", url)
	}

	return router
}

func (h *handler) getRoutes() Routes{
 return Routes{
			Route{
				"CreateUser",
				http.MethodPost,
				"/api/v1/user",
				h.CreateUser,
			},

			Route{
				"DeleteUser",
				http.MethodDelete,
				"/api/v1/user/{userId}",
				h.DeleteUser,
			},

			Route{
				Name: "FindUserById",
				Method: http.MethodGet,
				Pattern: "/api/v1/user/{userId}",
				HandlerFunc: h.FindUserById,
			},

			Route{
				"UpdateUser",
				http.MethodPut,
				"/api/v1/user/{userId}",
				h.UpdateUser,
			},
		}
}