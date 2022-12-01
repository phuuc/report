package repos

import (
	"context"

	"github.com/finnpn/overview/infras/usecases/models"
)

type HotelRepos interface {
	GetHotelTotalBookings(ctx context.Context) (*models.HotelTotalBookings, error)
	GetHotelTotalSpends(ctx context.Context) (*models.HotelTotalSpends, error)
	GetHotelAveragePrice(ctx context.Context) (*models.HotelAveragePrice, error)
	GetHotelTopBooked(ctx context.Context) ([]*models.HotelTopBooked, error)
}
