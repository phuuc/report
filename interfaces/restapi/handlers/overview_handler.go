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

func (h *Handlers) TotalTrips(c *gin.Context) {
	g := ginout.NewGinOut(c)
	ctx := c.Request.Context()

	data, err := h.overviewUc.GetTotalTrips(ctx)
	if err != nil {
		log.Printf(errors.ErrRequestFail, err)
		g.BadRequest(err.Error())
		return
	}

	g.JSONResponse(http.StatusOK, &models.TotalTrips{
		TotalTrips: data,
	})
}

func (h *Handlers) TotalBookings(c *gin.Context) {
	g := ginout.NewGinOut(c)
	ctx := c.Request.Context()

	data, err := h.overviewUc.GetTotalBookings(ctx)
	if err != nil {
		log.Printf(errors.ErrRequestFail, err)
		g.BadRequest(err.Error())
		return
	}

	g.JSONResponse(http.StatusOK, &models.TotalBookings{
		TotalBookings: data,
	})
}

func (h *Handlers) TotalSpends(c *gin.Context) {
	g := ginout.NewGinOut(c)
	ctx := c.Request.Context()

	data, err := h.overviewUc.GetTotalSpends(ctx)
	if err != nil {
		log.Printf(errors.ErrRequestFail, err)
		g.BadRequest(err.Error())
		return
	}

	g.JSONResponse(http.StatusOK, &models.TotalSpends{
		TotalSpends: data,
	})
}

func (h *Handlers) AverageSpend(c *gin.Context) {
	g := ginout.NewGinOut(c)
	ctx := c.Request.Context()

	data, err := h.overviewUc.GetAverageSpend(ctx)
	if err != nil {
		log.Printf(errors.ErrRequestFail, err)
		g.BadRequest(err.Error())
		return
	}

	g.JSONResponse(http.StatusOK, &models.AverageSpend{
		Flight:     helpers.ToIntVND(data.Flight),
		Hotel:      helpers.ToIntVND(data.Hotel),
		Combo:      helpers.ToIntVND(data.Combo),
		TourTicket: helpers.ToIntVND(data.TourTicket),
	})
}

func (h *Handlers) TotalSpendsByService(c *gin.Context) {
	g := ginout.NewGinOut(c)
	ctx := c.Request.Context()

	data, err := h.overviewUc.GetTotalSpendsByService(ctx)
	if err != nil {
		log.Printf(errors.ErrRequestFail, err)
		g.BadRequest(err.Error())
		return
	}

	g.JSONResponse(http.StatusOK, &models.TotalSpendsByService{
		Flight:     helpers.ToIntVND(data.Flight),
		Hotel:      helpers.ToIntVND(data.Hotel),
		Combo:      helpers.ToIntVND(data.Combo),
		TourTicket: helpers.ToIntVND(data.TourTicket),
	})
}
