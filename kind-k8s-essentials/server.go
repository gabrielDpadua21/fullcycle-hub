package main

import (
	"fmt"
	"net/http"
	"os"
	"io/ioutil"
	"log"
	"time"
)

func main() {
	fmt.Print("Server running....")
	http.HandleFunc("/", Hello)
	http.HandleFunc("/message", ConfigMap)
	http.HandleFunc("/secret", Secret)
	http.HandleFunc("/healthz", Healthz)
	http.ListenAndServe(":8000", nil)
}

func Hello(w http.ResponseWriter, r *http.Request) {
	name := os.Getenv("NAME")
	color := os.Getenv("COLOR")
	fmt.Fprintf(w, "Hello, %s. The kubernets has you, follow the %s habbit", name, color)
}

func ConfigMap(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadFile("env/matrix.txt")
	if err != nil {
		log.Fatalf("Error reading file", err)
	}
	fmt.Fprintf(w, "Hello, %s, i am morpheus, the matrix has you", string(data))
}

func Secret(w http.ResponseWriter, r *http.Request) {
	user := os.Getenv("USER")
	password := os.Getenv("PASSWORD")
	fmt.Fprintf(w, "User: %s. Password: %s", user, password)
}

func Healthz(w http.ResponseWriter, r *http.Request) {

	duration := time.Since(time.Now())

	if duration.Seconds() > 25 {
		w.WriteHeader(500)
		w.Write([]byte(fmt.Sprintf("Duration: %v", duration.Seconds())))
	} else {
		w.WriteHeader(200)
		w.Write([]byte("ok"))
	}

}