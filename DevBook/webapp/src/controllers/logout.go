package controllers

import (
	"net/http"
	cookies "webapp/src/coockies"
)


func FazerLogout(w http.ResponseWriter, r *http.Request) {
	cookies.Deletar(w)
	http.Redirect(w,r, "/login", 302)
}