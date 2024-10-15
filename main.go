package main

import (
	"fmt"
	"net/http"
)

func main(){

	mux := http.NewServeMux()

	mux.HandleFunc("/hello", handleHello)
	mux.HandleFunc("/bye", handleBye)

	srv := &http.Server{
		Addr: "127.0.0.1:8080",
		Handler: mux,
	}
	
	fmt.Println("Server is running on port 8080...")

	err := srv.ListenAndServe()
	if err != nil{
		fmt.Println("Server error: ", err)
		return
	}
}

func handleHello(w http.ResponseWriter, r *http.Request){
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Hello, this message is from server!")))
}
func handleBye(w http.ResponseWriter, r *http.Request){
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Goodbye!!")))
}
