package main

import (
	"fmt"
	template "html/template"
	"net/http"
	packageA "redisPlayground/PackageA"
	"time"
)

var globalSessions *packageA.Manager

func main() {
	// packageA.CreateSession(globalSessions)
	// http.HandleFunc("/index", login)
	// http.HandleFunc("/nyan", responseNyan)
	// http.ListenAndServe(":5000", nil)
	messages := make(chan []int)

	go f(messages)
	msg := <-messages
	fmt.Println(msg)
}

func responseNyan(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Nyan")
}

func f(messages chan []int) {
	hoge := []int{1, 2, 3}
	hoge = append(hoge, 4)
	messages <- hoge
}

type Page struct{}

func parseHtml(filePath string) template.Template {
	tmpl, err := template.ParseFiles("index.html")
	if err != nil {
		panic(err)
	}
	return *tmpl
}

func login(w http.ResponseWriter, r *http.Request) {
	// sess := globalSessions.SessionStart(w, r)
	if r.Method == "GET" {
		tmpl := parseHtml("index.html")
		w.Header().Set("Content-Type", "text/html")
		tmpl.Execute(w, Page{})
		// t.Execute(w, sess.Get("username"))
	} else {
		// sess.Set("username", r.Form["username"])
		http.Redirect(w, r, "/", 302)
	}
}

func count(w http.ResponseWriter, r *http.Request) {
	sess := globalSessions.SessionStart(w, r)
	createtime := sess.Get("createtime")
	if createtime == nil {
		sess.Set("createtime", time.Now().Unix())
	} else if (createtime.(int64) + 360) < (time.Now().Unix()) {
		globalSessions.SessionDestroy(w, r)
		sess = globalSessions.SessionStart(w, r)
	}
	ct := sess.Get("countnum")
	if ct == nil {
		sess.Set("countnum", 1)
	} else {
		sess.Set("countnum", (ct.(int) + 1))
	}
	t, _ := template.ParseFiles("count.gtpl")
	w.Header().Set("Content-Type", "text/html")
	t.Execute(w, sess.Get("countnum"))
}
