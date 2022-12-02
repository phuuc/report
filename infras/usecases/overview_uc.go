package usecases

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/finnpn/overview/infras/usecases/models"
	"github.com/finnpn/overview/interfaces/repos"
	"github.com/finnpn/overview/pkg/helpers"
	"github.com/go-redis/redis/v9"
)

type OverviewUc struct {
	overviewRepos repos.OverviewRepos
	redisClient   *redis.Client
}

func NewOverviewUc(overviewRepos repos.OverviewRepos, redisClient *redis.Client) *OverviewUc {
	return &OverviewUc{
		overviewRepos: overviewRepos,
		redisClient:   redisClient,
	}
}

func (u *OverviewUc) GetTotalTrips(ctx context.Context) (int, error) {
	redisTotal, err := u.redisClient.Get(ctx, "o-total-trips").Result()
	var totalTrips int
	if err == redis.Nil {
		totalTrips, err := u.overviewRepos.GetTotalTrips(ctx)
		if err != nil {
			log.Printf("fail getting total trips with err =%v", err)
			return 0, fmt.Errorf("fail getting total trips")
		}
		err = u.redisClient.Set(ctx, "o-total-trips", totalTrips, 0).Err()
		if err != nil {
			return 0, err
		}
	} else {
		totalTrips, err = strconv.Atoi(redisTotal)
		if err != nil {
			return 0, err
		}
	}
	return totalTrips, nil
}

func (u *OverviewUc) GetTotalBookings(ctx context.Context) (int, error) {
	totalBookings, err := u.overviewRepos.GetTotalBookings(ctx)
	if err != nil {
		log.Printf("fail getting total bookings with err =%v", err)
		return 0, fmt.Errorf("fail getting total bookings")
	}
	return totalBookings, nil
}

func (u *OverviewUc) GetTotalSpends(ctx context.Context) (string, error) {
	totalNumSpends, err := u.overviewRepos.GetTotalSpends(ctx)

	if err != nil {
		log.Printf("fail getting total bookings with err =%v", err)
		return "", fmt.Errorf("fail getting total bookings")
	}

	totalSpends := helpers.ToVND(totalNumSpends)

	return totalSpends, nil
}
func (u *OverviewUc) GetAverageSpend(ctx context.Context) (*models.AverageSpend, error) {
	averageSpend, err := u.overviewRepos.GetAverageSpend(ctx)

	if err != nil {
		log.Printf("fail getting GetAverageSpend with err =%v", err)
		return nil, fmt.Errorf("fail getting GetAverageSpend")
	}

	return averageSpend, nil
}
func (u *OverviewUc) GetTotalSpendsByService(ctx context.Context) (*models.TotalSpendsByService, error) {
	totalSpends, err := u.overviewRepos.GetTotalSpendsByService(ctx)
	if err != nil {
		log.Printf("fail getting GetTotalSpendsByService with err =%v", err)
		return nil, fmt.Errorf("fail gettingGetTotalSpendsByService")
	}

	return totalSpends, nil
}
