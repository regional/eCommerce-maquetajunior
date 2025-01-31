package handlers

import (
	"net/http"
)


func Healt(rw http.ResponseWriter, r *http.Request) {
	
	sendData(rw, "Majestic API Running", http.StatusOK)
}