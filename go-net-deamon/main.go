package main

import (
	"fmt"
	"net/http"
	"path"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	//中间件
	router.Use(LoggingMiddleware)

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "welcome to the homepage")
	})

	router.HandleFunc("/image/", func(w http.ResponseWriter, r *http.Request) {
		image_dir := "C:/Users/Gin/Desktop/image"
		http.ServeFile(w, r, image_dir)
	})

	router.HandleFunc("/image/{id}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]
		image_dir := "C:/Users/Gin/Desktop/image"
		http.ServeFile(w, r, path.Join(image_dir, id))
	})

	router.HandleFunc("/user/{name}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		name := vars["name"]
		fmt.Fprintf(w, "hello, %s!", name)
	}).Methods("GET")

	router.HandleFunc("/user/{name:[a-z]+}/{id:[0-9]+}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		name := vars["name"]
		id := vars["id"]

		fmt.Fprintf(w, "Name:%s, ID:%s", name, id)
	}).Methods("GET")

	//子路由器
	apiRouter := router.PathPrefix("/api").Subrouter()
	apiRouter.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "this is api")
	})

	apiRouter.HandleFunc("/images/", func(w http.ResponseWriter, r *http.Request) {
		image_dir := "C:/Users/Gin/Desktop/image"
		http.ServeFile(w, r, image_dir)
	}).Methods("GET")

	apiRouter.HandleFunc("/images/{id}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]
		image_dir := "C:/Users/Gin/Desktop/image"
		http.ServeFile(w, r, path.Join(image_dir, id))
	}).Methods("GET")

	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	name := r.URL.Query().Get("name")
	// 	if name == "" {
	// 		name = "guest"
	// 	}
	// 	fmt.Fprintf(w, "Hello %s!", name)
	// })

	// image_dir := http.Dir("C:/Users/Gin/Desktop/image")
	// image_file := http.FileServer(image_dir)
	// http.Handle("/image/", http.StripPrefix("/image/", image_file))

	http.ListenAndServe(":8080", router)
	// http.ListenAndServe(":8080", apiRouter)
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("request received:", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}
