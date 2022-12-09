package main

import (
	"embed"
	_ "embed"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/patrickmn/go-cache"
	"github.com/trietmn/go-wiki"
	"log"
	"net/http"
	"os"
	"text/template"
	"time"
)

//go:embed "wrapper.html.tmpl"
var wrapper string

//go:embed static
var static embed.FS

var Cache = cache.New(5*time.Minute, 5*time.Minute)

type pageData struct {
	Title   string
	Content string
}

func getHTML(pageName string) (pageData, error) {
	var pd pageData

	data, found := Cache.Get(pageName)
	if found {
		pd = data.(pageData)
	} else {
		page, err := gowiki.GetPage(pageName, -1, false, true)
		if err != nil {
			log.Println(err)
			return pageData{}, err
		}
		content, err := page.GetHTML()
		if err != nil {
			log.Println(err)
			pd = pageData{Title: page.Title, Content: content}
			return pd, err
		}
		pd = pageData{Title: page.Title, Content: content}
		Cache.Set(pageName, pd, cache.DefaultExpiration)
	}
	return pd, nil
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func PageHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	t, err := template.New("wrapper").Parse(wrapper)
	if err != nil {
		log.Println(err)
		return
	}
	pd, err := getHTML(vars["page"])
	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
		_, err := fmt.Fprint(w, err)
		if err != nil {
			return
		}
	} else {
		w.WriteHeader(http.StatusOK)
		err = t.Execute(w, pd)
		if err != nil {
			log.Println(err)
		}
	}
}

func main() {
	port := getEnv("PORT", "8080")

	r := mux.NewRouter()
	r.HandleFunc("/wiki/{page}", PageHandler)
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.FS(static))))

	log.Printf("Listening on http://127.0.0.1:%v", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), r))
}
