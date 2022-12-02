package server

import (
	"log"
	"net/http"

	"github.com/finnpn/overview/config"
	"github.com/finnpn/overview/interfaces/restapi/handlers"
	"github.com/gin-gonic/gin"
)

type Server struct {
	router *gin.Engine
	config *config.Config
}

func NewServer(cfg *config.Config, handlers *handlers.Handlers) *Server {
	router := gin.Default()

	routeOverview := router.Group("/overview")
	handlers.ConfigureOverview(routeOverview)

	routerFlight := router.Group("/flight")
	handlers.ConfigureFlights(routerFlight)

	routerHotel := router.Group("/hotel")
	handlers.ConfigureHotels(routerHotel)

	return &Server{
		router: router,
		config: cfg,
	}
}
func (r *Server) Run() {
	s := &http.Server{
		Addr:    config.Addr(r.config.Server.Host, r.config.Server.Port),
		Handler: r.router,
		// ReadTimeout:    10 * time.Second,
		// WriteTimeout:   10 * time.Second,
		// MaxHeaderBytes: 1 << 20,
	}
	err := s.ListenAndServe()
	if err != nil {
		log.Fatalf("server could not run with err=%v", err)
	}
	log.Printf("running with addr : %s", s.Addr)
}
