package main

import (
	"fmt"
	"html/template"
	"net/http"
	"path"
	app "tester_app/app"
	"time"
)

func handle(w http.ResponseWriter, r *http.Request) {
	// You might want to move ParseGlob outside of handle so it doesn't
	// re-parse on every http request.
	tmpl, err := template.ParseGlob("templates/*")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	name := ""
	if r.URL.Path == "/" {
		name = "index.html"
	} else {
		name = path.Base(r.URL.Path)
	}

	data := struct {
		Time time.Time
	}{
		Time: time.Now(),
	}

	if err := tmpl.ExecuteTemplate(w, name, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Println("error", err)
	}
}

func handlehome(w http.ResponseWriter, r *http.Request) {
	// You might want to move ParseGlob outside of handle so it doesn't
	// re-parse on every http request.
	tmpl, err := template.ParseGlob("templates/*")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	name := ""
	if r.URL.Path == "/home" {
		name = "home.html"
	} else {
		name = path.Base(r.URL.Path)
	}

	data := struct {
		Data  string
		Data2 string
	}{
		Data:  "testing out vars in go",
		Data2: app.HiThere(),
	}

	if err := tmpl.ExecuteTemplate(w, name, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Println("error", err)
	}
}

func handleSignup(w http.ResponseWriter, r *http.Request) {
	fmt.Println("sign up")
	tmpl, err := template.ParseGlob("templates/*")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	name := ""
	if r.URL.Path == "/signup" {
		name = "signup.html"
	} else {
		name = path.Base(r.URL.Path)
	}

	data := struct {
		Data  string
		Data2 app.Users
	}{
		Data:  "testing out vars in go",
		Data2: app.CreateUser("Chris", "test", "Active"),
	}

	if err := tmpl.ExecuteTemplate(w, name, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Println("error", err)
	}
}

func main() {
	fmt.Println("http server up!")
	http.Handle(
		"/static/",
		http.StripPrefix(
			"/static/",
			http.FileServer(http.Dir("static")),
		),
	)
	http.HandleFunc("/", handle)
	http.HandleFunc("/home", handlehome)
	http.HandleFunc("/signup", handleSignup)
	http.ListenAndServe(":8080", nil)
}
