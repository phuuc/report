package usecases

import (
	"context"
	"fmt"
	"log"

	"github.com/finnpn/overview/infras/usecases/models"
	"github.com/finnpn/overview/interfaces/repos"
)

type HotelUc struct {
	hotelRepos repos.HotelRepos
}

func NewHotelUc(hotelRepos repos.HotelRepos) *HotelUc {
	return &HotelUc{
		hotelRepos: hotelRepos,
	}
}

func (u *HotelUc) GetHotelTotalBookings(ctx context.Context) (*models.HotelTotalBookings, error) {
	totalBookings, err := u.hotelRepos.GetHotelTotalBookings(ctx)
	if err != nil {
		log.Printf("fail getting total bookings with err =%v", err)
		return nil, fmt.Errorf("fail getting total bookings")
	}
	return totalBookings, nil
}

func (u *HotelUc) GetHotelTotalSpends(ctx context.Context) (*models.HotelTotalSpends, error) {
	totalSpends, err := u.hotelRepos.GetHotelTotalSpends(ctx)
	if err != nil {
		log.Printf("fail getting Get GetHotelTotalSpends with err =%v", err)
		return nil, fmt.Errorf("fail getting GetHotelTotalSpends")
	}
	return totalSpends, nil
}

func (u *HotelUc) GetHotelAveragePrice(ctx context.Context) (*models.HotelAveragePrice, error) {
	averagePrice, err := u.hotelRepos.GetHotelAveragePrice(ctx)
	if err != nil {
		log.Printf("fail gettingGet GetHotelAveragePrice with err =%v", err)
		return nil, fmt.Errorf("fail getting GetHotelAveragePrice")
	}
	return averagePrice, nil
}
func (u *HotelUc) GetHotelTopBooked(ctx context.Context) ([]*models.HotelTopBooked, error) {
	topBooked, err := u.hotelRepos.GetHotelTopBooked(ctx)

	if err != nil {
		log.Printf("fail gettingGet GetHotelTopBooked with err =%v", err)
		return nil, fmt.Errorf("fail getting GetHotelTopBooked")
	}
	return topBooked, nil
}
