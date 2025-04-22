package main

import(
	"log"
	"time"
	"net/http"
	"github.com/gorilla/mux"
)
func main () {
	InitDB();
	InitUtils();
	
	//Setup Router
	router := mux.NewRouter();

	//Define Routes
	router.HandleFunc("/signup",  SignupHandler).Methods(http.MethodPost);
	router.HandleFunc("/signin" , SigninHandler).Methods(http.MethodPost);

	//Define Server
	server := &http.Server {
		Handler: 		router,
		Addr:			":8080",
		WriteTimeout: 15 * time.Second,
        ReadTimeout:  15 * time.Second,
        IdleTimeout:  60 * time.Second,
	}

	log.Println("Starting server on port 8080")

	if err := server.ListenAndServe(); err != nil {
        log.Fatal("ListenAndServe Error: ", err)
    }

}