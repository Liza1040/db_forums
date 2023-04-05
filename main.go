package main

import (
	"fmt"
	"log"
	"net/http"

	"DB_forums/models"
	"DB_forums/server/handlers"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var requestsTotal = prometheus.NewCounter(
	prometheus.CounterOpts{
		Name: "total_requests",
	})

func middlewareFunc(_ *mux.Router) mux.MiddlewareFunc {
	return func(handler http.Handler) http.Handler {
		return http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
			requestsTotal.Inc()
			response.Header().Set("Content-Type", "application/json")
			handler.ServeHTTP(response, request)
		})
	}
}

var hitsCounter = prometheus.NewCounter(prometheus.CounterOpts{
	Name: "hits_counter",
	Help: "Number of hits to the server",
})

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
		hitsCounter.Inc()
	})
}

func main() {
	prometheus.MustRegister(requestsTotal)

	parsedConnection, err := pgx.ParseConnectionString("host=localhost user=postgres password=root dbname=DB_forums sslmode=disable")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	parsedConnection.PreferSimpleProtocol = true

	connectionConfig := pgx.ConnPoolConfig{
		ConnConfig:     parsedConnection,
		MaxConnections: 69,
		AfterConnect:   nil,
		AcquireTimeout: 0,
	}

	models.DB, err = pgx.NewConnPool(connectionConfig)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	/*---------TEST--------
	handlers.TestDropDB()
	handlers.TestInitDB()
	//---------TEST--------*/

	router := mux.NewRouter()
	prometheus.MustRegister(hitsCounter)
	router.Use(loggingMiddleware)
	router.Path("/metrics").Handler(promhttp.Handler())
	router.Use(middlewareFunc(router))

	// router.Handle("/metrics", promhttp.Handler()).Methods(http.MethodGet)
	router.HandleFunc("/api/forum/create", handlers.ForumCreate).Methods(http.MethodPost)
	router.HandleFunc("/api/forum/{slug}/details", handlers.ForumDetails).Methods(http.MethodGet)
	router.HandleFunc("/api/forum/{slug}/create", handlers.ThreadCreate).Methods(http.MethodPost)
	router.HandleFunc("/api/forum/{slug}/users", handlers.ForumUsers).Methods(http.MethodGet)
	router.HandleFunc("/api/forum/{slug}/threads", handlers.ForumThreads).Methods(http.MethodGet)

	router.HandleFunc("/api/post/{id:[0-9]+}/details", handlers.PostDetails).Methods(http.MethodGet, http.MethodPost)

	//router.HandleFunc("/api/service/init", handlers.InitDB)
	//router.HandleFunc("/api/service/drop", handlers.DropDB)
	router.HandleFunc("/api/service/clear", handlers.ClearDB).Methods(http.MethodPost)
	router.HandleFunc("/api/service/status", handlers.StatusDB).Methods(http.MethodGet)

	router.HandleFunc("/api/thread/{slug_or_id}/create", handlers.PostsCreate).Methods(http.MethodPost)
	router.HandleFunc("/api/thread/{slug_or_id}/details", handlers.ThreadDetails).Methods(http.MethodGet, http.MethodPost)
	router.HandleFunc("/api/thread/{slug_or_id}/posts", handlers.ThreadPosts).Methods(http.MethodGet)
	router.HandleFunc("/api/thread/{slug_or_id}/vote", handlers.VoteCreate).Methods(http.MethodPost)

	router.HandleFunc("/api/user/{nickname}/create", handlers.UserCreate).Methods(http.MethodPost)
	router.HandleFunc("/api/user/{nickname}/profile", handlers.UserProfile).Methods(http.MethodGet, http.MethodPost)

	http.Handle("/", router)

	port := "5000"
	fmt.Println("Server listen at: ", port)
	err = http.ListenAndServe(":"+port, nil)
}
