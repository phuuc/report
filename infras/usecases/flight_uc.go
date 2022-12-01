package usecases

import (
	"context"
	"fmt"
	"log"

	"github.com/finnpn/overview/infras/usecases/models"
	"github.com/finnpn/overview/interfaces/repos"
)

type FlightUc struct {
	flightRepos repos.FlightRepos
}

func NewFlightUc(flightRepos repos.FlightRepos) *FlightUc {
	return &FlightUc{
		flightRepos: flightRepos,
	}
}

func (u *FlightUc) GetFlightTotalBookings(ctx context.Context) (*models.FlightTotalBooking, error) {
	totalBookings, err := u.flightRepos.GetFlightTotalBooking(ctx)
	if err != nil {
		log.Printf("fail getting total bookings with err =%v", err)
		return nil, fmt.Errorf("fail getting total bookings")
	}
	return totalBookings, nil
}

func (u *FlightUc) GetFlightTotalSpends(ctx context.Context) (*models.FlightTotalSpends, error) {
	totalSpends, err := u.flightRepos.GetFlightTotalSpends(ctx)
	if err != nil {
		log.Printf("fail getting Get FlightTotalSpends with err =%v", err)
		return nil, fmt.Errorf("fail getting GetFlightTotalSpends")
	}
	return totalSpends, nil
}

func (u *FlightUc) GetFlightAveragePrice(ctx context.Context) (*models.FlightAveragePrice, error) {
	averagePrice, err := u.flightRepos.GetFlightAveragePrice(ctx)
	if err != nil {
		log.Printf("fail gettingGet GetFlightAveragePrice with err =%v", err)
		return nil, fmt.Errorf("fail getting GetFlightAveragePrice")
	}
	return averagePrice, nil
}

func (u *FlightUc) GetFlightTopAirlines(ctx context.Context) ([]*models.FlightTopAirlines, error) {
	topAirlines, err := u.flightRepos.GetFlightTopAirlines(ctx)

	if err != nil {
		log.Printf("fail gettingGet GetFlightTopAirlines with err =%v", err)
		return nil, fmt.Errorf("fail getting GetFlightTopAirlines")
	}
	return topAirlines, nil
}

func (u *FlightUc) GetFlightBookingWindowDays(ctx context.Context) (*models.FlightBookingWindowDays, error) {
	windowsDays, err := u.flightRepos.GetFlightBookingWindowDays(ctx)

	if err != nil {
		log.Printf("fail gettingGet GetFlightBookingWindowDays with err =%v", err)
		return nil, fmt.Errorf("fail getting GetFlightBookingWindowDays")
	}
	return windowsDays, nil
}
