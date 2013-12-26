package main

import "net/http"

type greetings struct {
	Intro    string
	Messages []string
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	passedObj := greetings{
		Intro:    "Hello from Go!",
		Messages: []string{"Hello!", "Hi!", "¡Hola!", "Bonjour!", "Ciao!", "<script>evilScript()</script>"},
	}
	templates.ExecuteTemplate(w, "homePage", passedObj)
}
