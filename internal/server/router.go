package server

import (
	"github.com/kynetiv/simple-weather/internal"
)

func (s *Server) AttachRoutes(api internal.WeatherService) {
	// define routes here
	s.Handle("/api/v1/conditions", api.GetCondition)
}
