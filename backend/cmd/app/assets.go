package main

import "net/http"

func RegisterAssets() {
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./public/assets"))))

}
