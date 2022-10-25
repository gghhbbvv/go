//Запускаем сервер. http://localhost:8000/?name=Vlad&age=99
package main

import (
	"fmt"
	"log"
	"net/http"
)

func greetings(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	age := r.URL.Query().Get("age")
	fmt.Fprintf(w, "<h1>Hello! My name is %s. My age is %s. </h1>", name, age)
}

func main() {
	http.HandleFunc("/", greetings)
	log.Println("http://localhost:8000")
	http.ListenAndServe(":8000", nil)
}
