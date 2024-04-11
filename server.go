package main

import "net/http"

func main() {
	http.HandleFunc("/", Hello)
	http.ListenAndServe(":3000", nil)
}

func Hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1><h3><i>ʕ•́ᴥ•̀ʔっ</i> Jonas esteve aqui <i>ʕ•́ᴥ•̀ʔっ</i></h3></h1>"))
}
