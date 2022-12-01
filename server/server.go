package server

import (
	"log"
	"net/http"

	"github.com/finnpn/overview/config"
	"github.com/finnpn/overview/interfaces/restapi/handlers"
	"github.com/gin-gonic/gin"
)

type Router struct {
	router *gin.Engine
	config *config.Config
}

func NewRouter(cfg *config.Config, handlers *handlers.Handlers) *Router {
	router := gin.Default()

	routeOverview := router.Group("/overview")
	handlers.ConfigureOverview(routeOverview)

	routerFlight := router.Group("/flight")
	handlers.ConfigureFlights(routerFlight)

	routerHotel := router.Group("/hotel")
	handlers.ConfigureHotels(routerHotel)

	return &Router{
		router: router,
		config: cfg,
	}
}
func (r *Router) Run() {
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
