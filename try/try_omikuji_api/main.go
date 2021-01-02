package main

import (
	"math/rand"
	"net/http"
	"text/template"
	"time"
)

var tmpl = template.Must(template.New("msg").
	Parse("<html><body>{{.Name}}さんの運勢は「<b>{{.Omikuji}}</b>」です</body></html>"))

type Result struct {
	Name    string
	Omikuji string
}

func omikuji() string {
	n := rand.Intn(6) // 0-5

	var result string
	switch n + 1 {
	case 6:
		result = "大吉"
	case 5, 4:
		result = "中吉"
	case 3, 2:
		result = "小吉"
	case 1:
		result = "凶"
	}
	return result
}

func handler(w http.ResponseWriter, r *http.Request) {
	var name string
	if v := r.FormValue("p"); v != "" {
		name = v
	} else {
		name = "guest"
	}

	result := Result{
		Name:    name,
		Omikuji: omikuji(),
	}
	tmpl.Execute(w, result)

}

func main() {
	// 現在時刻
	t := time.Now().UnixNano()
	rand.Seed(t)

	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
