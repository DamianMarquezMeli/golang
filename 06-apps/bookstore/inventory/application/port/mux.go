package port

import "net/http"

type MuxRouter interface {
	get(path string, handler func(w http.ResponseWriter, r *http.Request))
	post(path string, handler func(w http.ResponseWriter, r *http.Request))
	put(path string, handler func(w http.ResponseWriter, r *http.Request))
	patch(path string, handler func(w http.ResponseWriter, r *http.Request))
	delete(path string, handler func(w http.ResponseWriter, r *http.Request))
}
