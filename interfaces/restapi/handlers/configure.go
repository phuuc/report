package handlers

import (
	"github.com/finnpn/overview/infras/usecases"
	"github.com/gin-gonic/gin"
)

type Handlers struct {
	overviewUc *usecases.OverviewUc
	flightUc   *usecases.FlightUc
	hotelUc    *usecases.HotelUc
}

func NewHandlers(overviewUc *usecases.OverviewUc, flightUc *usecases.FlightUc, hotelUc *usecases.HotelUc) *Handlers {
	return &Handlers{
		overviewUc: overviewUc,
		flightUc:   flightUc,
		hotelUc:    hotelUc,
	}
}

func (h *Handlers) ConfigureOverview(router *gin.RouterGroup) {
	router.GET("/total-trips", h.TotalTrips)
	router.GET("/total-bookings", h.TotalBookings)
	router.GET("/total-spends", h.TotalSpends)
	router.GET("/average-spend", h.AverageSpend)
	router.GET("/total-spends-service", h.TotalSpendsByService)
}

func (h *Handlers) ConfigureFlights(router *gin.RouterGroup) {
	router.GET("/total-bookings", h.FlightTotalBooking)
	router.GET("/total-spends", h.FlightTotalSpends)
	router.GET("/average-price", h.FlightAveragePrice)
	router.GET("/top-airlines", h.FlightTopAirlines)
	router.GET("/booking-window-days", h.FlightBookingWindowDays)
}

func (h *Handlers) ConfigureHotels(router *gin.RouterGroup) {
	router.GET("/total-bookings", h.HotelTotalBooking)
	router.GET("/total-spends", h.HotelTotalSpends)
	router.GET("/average-price", h.HotelAveragePrice)
	router.GET("/top-booked", h.HotelTopBooked)
}
