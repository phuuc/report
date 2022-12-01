package handlers

import (
	"log"
	"net/http"

	"github.com/finnpn/overview/interfaces/restapi/models"
	"github.com/finnpn/overview/pkg/errors"
	"github.com/finnpn/overview/pkg/ginout"
	"github.com/finnpn/overview/pkg/helpers"
	"github.com/gin-gonic/gin"
)

func (h *Handlers) FlightTotalBooking(c *gin.Context) {
	g := ginout.NewGinOut(c)
	ctx := c.Request.Context()

	data, err := h.flightUc.GetFlightTotalBookings(ctx)
	if err != nil {
		log.Printf(errors.ErrRequestFail, err)
		g.BadRequest(err.Error())
		return
	}

	g.JSONResponse(http.StatusOK, &models.FlightTotalBooking{
		TotalBookings: data.TotalBookings,
	})
}

func (h *Handlers) FlightTotalSpends(c *gin.Context) {
	g := ginout.NewGinOut(c)
	ctx := c.Request.Context()

	data, err := h.flightUc.GetFlightTotalSpends(ctx)
	if err != nil {
		log.Printf(errors.ErrRequestFail, err)
		g.BadRequest(err.Error())
		return
	}

	g.JSONResponse(http.StatusOK, &models.FlightTotalSpends{
		TotalSpends: helpers.ToVND(data.TotalSpends),
	})
}

func (h *Handlers) FlightAveragePrice(c *gin.Context) {
	g := ginout.NewGinOut(c)
	ctx := c.Request.Context()

	data, err := h.flightUc.GetFlightAveragePrice(ctx)
	if err != nil {
		log.Printf(errors.ErrRequestFail, err)
		g.BadRequest(err.Error())
		return
	}

	g.JSONResponse(http.StatusOK, &models.FlightAveragePrice{
		AveragePrice: helpers.ToVND(data.AveragePrice),
	})
}

func (h *Handlers) FlightTopAirlines(c *gin.Context) {
	g := ginout.NewGinOut(c)
	ctx := c.Request.Context()

	data, err := h.flightUc.GetFlightTopAirlines(ctx)
	if err != nil {
		log.Printf(errors.ErrRequestFail, err)
		g.BadRequest(err.Error())
		return
	}
	res := []*models.FlightTopAirlines{}
	for _, t := range data {
		res = append(res, &models.FlightTopAirlines{
			Name:         t.Name,
			BookingCount: t.BookingCount,
			AveragePrice: helpers.ToVND(t.AveragePrice),
		})
	}

	g.JSONResponse(http.StatusOK, res)
}
func (h *Handlers) FlightBookingWindowDays(c *gin.Context) {
	g := ginout.NewGinOut(c)
	ctx := c.Request.Context()

	data, err := h.flightUc.GetFlightBookingWindowDays(ctx)
	if err != nil {
		log.Printf(errors.ErrRequestFail, err)
		g.BadRequest(err.Error())
		return
	}

	g.JSONResponse(http.StatusOK, &models.FlightBookingWindowDays{
		BookingWindowDays: data.BookingWindowDays,
	})
}
