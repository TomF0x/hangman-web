package main

import (
	"fmt"
	hw "hangman-web"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"
	"text/template"
	"time"
)

var args = os.Args[1:]
var try = 10
var randomwordhide, randomword = "", ""
var usedletter = []string{}
var state = ""
var message = ""
var tmpl *template.Template
var win = false
var loose = false
var username = ""
var diff = ""
var timer = time.Now().Second()
var chrono = 0

func main() {
	tmpl, _ = template.ParseFiles("../html/difficulties.html")
	fs := http.FileServer(http.Dir("../css/"))
	http.Handle("/css/", http.StripPrefix("/css/", fs))
	fs = http.FileServer(http.Dir("../imgs/"))
	http.Handle("/imgs/", http.StripPrefix("/imgs/", fs))
	http.HandleFunc("/", handler)
	http.HandleFunc("/reset", reste)
	http.HandleFunc("/scoreboard", hw.Scoreboard)
	http.HandleFunc("/hangman", hangman)
	http.HandleFunc("/difficulties", difficulties)
	http.ListenAndServe("localhost:8000", nil)
}

func generate(file string) (string, string) {
	rand.Seed(time.Now().UnixNano())
	wordlist := hw.Loadressource(file, 1)
	randomword = wordlist[rand.Intn(len(wordlist))]
	fmt.Println(randomword)
	return hw.GenerateWord(randomword), randomword
}

func hangman(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		w.WriteHeader(404)
	case "POST":
		letter := r.FormValue("letter")
		randomwordhide, state = hw.Finder(letter, randomword, randomwordhide, &usedletter)
		if state == "wordgood" {
			win = true
		}
		message = hw.Message(state, randomword, letter, &try)
		w.WriteHeader(200)
	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}
}

func difficulties(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		w.WriteHeader(404)
	case "POST":
		tmpl, _ = template.ParseFiles("../html/index.html")
		diff = r.FormValue("difficultie")
		randomwordhide, randomword = generate("../difficulties/" + diff + ".txt")
		username = r.FormValue("username")
		timer = time.Now().Second()
	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}
}

func reste(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		w.WriteHeader(404)
	case "POST":
		replay := r.FormValue("reset")
		if replay == "true" {
			try = 10
			randomwordhide, randomword = "", ""
			usedletter = []string{}
			state = ""
			message = ""
			win = false
			loose = false
			timer = time.Now().Second()
			tmpl, _ = template.ParseFiles("../html/difficulties.html")
		}
	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	if len(args) > 0 && args[0] == "--split" && randomwordhide != "" {
		tmpl, _ = template.ParseFiles("../html/bring-to-death.html")
	}
	if randomwordhide != "" && randomword != "" {
		if randomwordhide == randomword {
			win = true
		} else if try < 1 {
			loose = true
		}
	}
	if win == true {
		win = false
		tmpl, _ = template.ParseFiles("../html/win.html")
		chrono = time.Now().Second() - timer
		f, _ := os.OpenFile("scores.txt", os.O_APPEND|os.O_WRONLY, 0644)

		defer f.Close()

		if diff == "hard" {
			_, _ = f.WriteString(username + "|" + diff + "|" + strconv.Itoa(chrono) + "|" + strconv.Itoa((120-chrono)*3) + "\n")
		} else if diff == "normal" {
			_, _ = f.WriteString(username + "|" + diff + "|" + strconv.Itoa(chrono) + "|" + strconv.Itoa((120-chrono)*2) + "\n")
		} else {
			_, _ = f.WriteString(username + "|" + diff + "|" + strconv.Itoa(chrono) + "|" + strconv.Itoa(120-chrono) + "\n")
		}
		randomwordhide, randomword = "", ""
	} else if loose == true {
		loose = false
		tmpl, _ = template.ParseFiles("../html/loose.html")
	}

	var used = ""

	for _, letter := range usedletter {
		used += letter + " "
	}

	data := struct {
		Usedletter string
		Randomword string
		HideWord   string
		Try        string
		Msg        string
		Jose       string
	}{
		Usedletter: strings.ToUpper(used),
		Randomword: randomword,
		HideWord:   strings.ToUpper(randomwordhide),
		Try:        strconv.Itoa(try),
		Msg:        message,
		Jose:       strconv.Itoa(10 - try),
	}
	_ = tmpl.Execute(w, data)
}
