package main

import (
	"fmt"
	"log"
	"net/http"
)

func Route() *http.ServeMux {
	m := http.NewServeMux()
	m.HandleFunc("/greet", func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		fmt.Fprintf(w, "Hello, %s", r.FormValue("name"))
	})
	return m
}

func main() {
	m := Route()
	log.Fatal(http.ListenAndServe(":8080", m))
}

/*
このアプリケーションをテストするには、
このサーバーをgo runして立ち上げ、
curl http://localhost:8080/greet?name=gopher
を実行してレスポンスを確かめることもできる
こういった場合にはnet/http/httptestを使うと便利です
*/
