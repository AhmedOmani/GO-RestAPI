package main

import(
	"log"
	"time"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main () {
	InitDB();
	InitAuth();
	InitUtils();
	InitRedis();

	//Setup Router
	router := mux.NewRouter();

	//Define Routes
	router.HandleFunc("/signup",  SignupHandler).Methods(http.MethodPost);
	router.HandleFunc("/signin" , SigninHandler).Methods(http.MethodPost);

	//CORS middleware
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000"},
		AllowedMethods: []string{"GET" , "POST" , "PUT" , "DELETE" , "OPTIONS"},
		AllowedHeaders: []string{"Authorization" , "Content-Type"},
	});

	//rap router with CORS middleware
	handler := c.Handler(router);

	//Define Server
	server := &http.Server {
		Handler: 		handler,
		Addr:			":8080",
		WriteTimeout: 15 * time.Second,
        ReadTimeout:  15 * time.Second,
        IdleTimeout:  60 * time.Second,
	}

	log.Println("Starting server on port 8080.")

	if err := server.ListenAndServe(); err != nil {
        log.Fatal("ListenAndServe Error: ", err)
    }

}