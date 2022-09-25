package server

import (
	"log"
	"net/http"

	"simple-weather/internal"
)

type Server struct {
	*http.ServeMux
}

func New(api internal.WeatherService) *Server {
	srv := &Server{
		http.NewServeMux(),
	}
	srv.AttachRoutes(api)
	return srv

}

func (s *Server) Start() {
	log.Fatal(http.ListenAndServe("localhost:8080", s))
}
