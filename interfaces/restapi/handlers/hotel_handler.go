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

func (h *Handlers) HotelTotalBooking(c *gin.Context) {
	g := ginout.NewGinOut(c)
	ctx := c.Request.Context()

	data, err := h.hotelUc.GetHotelTotalBookings(ctx)
	if err != nil {
		log.Printf(errors.ErrRequestFail, err)
		g.BadRequest(err.Error())
		return
	}

	g.JSONResponse(http.StatusOK, &models.HotelTotalBookings{
		TotalBookings: data.TotalBookings,
	})
}
func (h *Handlers) HotelTotalSpends(c *gin.Context) {
	g := ginout.NewGinOut(c)
	ctx := c.Request.Context()

	data, err := h.hotelUc.GetHotelTotalSpends(ctx)
	if err != nil {
		log.Printf(errors.ErrRequestFail, err)
		g.BadRequest(err.Error())
		return
	}

	g.JSONResponse(http.StatusOK, &models.HotelTotalSpends{
		TotalSpends: helpers.ToVND(data.TotalSpends),
	})
}

func (h *Handlers) HotelAveragePrice(c *gin.Context) {
	g := ginout.NewGinOut(c)
	ctx := c.Request.Context()

	data, err := h.hotelUc.GetHotelAveragePrice(ctx)
	if err != nil {
		log.Printf(errors.ErrRequestFail, err)
		g.BadRequest(err.Error())
		return
	}

	g.JSONResponse(http.StatusOK, &models.HotelAveragePrice{
		AveragePrice: helpers.ToVND(data.AveragePrice),
	})
}

func (h *Handlers) HotelTopBooked(c *gin.Context) {
	g := ginout.NewGinOut(c)
	ctx := c.Request.Context()

	data, err := h.hotelUc.GetHotelTopBooked(ctx)
	if err != nil {
		log.Printf(errors.ErrRequestFail, err)
		g.BadRequest(err.Error())
		return
	}
	res := []*models.HotelTopBooked{}
	for _, t := range data {
		res = append(res, &models.HotelTopBooked{
			Name:         t.Name,
			BookingCount: t.BookingCount,
			AveragePrice: helpers.ToVND(t.AveragePrice),
		})
	}

	g.JSONResponse(http.StatusOK, res)
}
