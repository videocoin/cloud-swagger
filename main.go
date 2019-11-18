package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"
)

type Page struct {
	Url string
}

func main() {

	var urls = map[string]string{
		"swagger/streams":  "/swagger/openapi/streams/v1/streams_service.swagger.json",
		"swagger/users":    "/swagger/openapi/users/v1/user_service.swagger.json",
		"swagger/accounts": "/swagger/openapi/accounts/v1/account_service.swagger.json",
		"swagger/profiles": "/swagger/openapi/profiles/v1/profiles_service.swagger.json",
		"swagger/miners":   "/swagger/openapi/miners/v1/miner_service.swagger.json",
	}

	http.HandleFunc("/swagger", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("./service/templates/index.html"))
		tmpl.Execute(w, Page{})
	})
	http.HandleFunc("/swagger/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("hit page")

		tmpl := template.Must(template.ParseFiles("./service/templates/base.html"))
		data := Page{
			Url: urls[strings.ToLower(r.URL.Path[1:])],
		}
		tmpl.Execute(w, data)
	})

	fsOpenapi := http.FileServer(http.Dir("./proto/openapi"))
	http.Handle("/swagger/openapi/", http.StripPrefix("/swagger/openapi/", fsOpenapi))

	fsAssets := http.FileServer(http.Dir("./service/assets"))
	http.Handle("/swagger/assets/", http.StripPrefix("/swagger/assets/", fsAssets))
	fmt.Printf("Starting server...")
	http.ListenAndServe(":8080", nil)
}
