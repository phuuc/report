package repos

import (
	"context"

	"github.com/finnpn/overview/infras/usecases/models"
)

type FlightRepos interface {
	GetFlightTotalBooking(ctx context.Context) (*models.FlightTotalBooking, error)
	GetFlightTotalSpends(ctx context.Context) (*models.FlightTotalSpends, error)
	GetFlightAveragePrice(ctx context.Context) (*models.FlightAveragePrice, error)
	GetFlightTopAirlines(ctx context.Context) ([]*models.FlightTopAirlines, error)
	GetFlightBookingWindowDays(ctx context.Context) (*models.FlightBookingWindowDays, error)
}
