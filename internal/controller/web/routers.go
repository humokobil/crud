package web

import (
	"fmt"
	"net/http"
	"otus/crud/v1/internal/entity"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
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
	GetPromCollectors() []prometheus.Collector
}

type handler struct {
	usecase *entity.IUsecase
	reqCounter *prometheus.CounterVec
	reqDuration *prometheus.HistogramVec
}

func New(usecase *entity.IUsecase) IHandler {
	h := &handler{usecase: usecase}
	h.initMetrics()
	return h
}

func (h *handler) initMetrics() {
	h.reqCounter = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "http_request_count",
		Help: "Number of http requests",
	},
	[]string{"status", "method"})

	h.reqDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			//Namespace:   "otus-crud",
			//Subsystem:   "otus_subsystem",
			Name:        "http_request_duration_histogram_seconds",
			Help:        "Request time duration.",
			Buckets:      []float64{.5, .95, .99},//prometheus.DefBuckets,
			//ConstLabels: []string{"constLabels"},
		},
		[]string{"method"},
	)
}

func (h *handler) GetPromCollectors() []prometheus.Collector {
	return []prometheus.Collector{h.reqCounter, h.reqDuration}
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