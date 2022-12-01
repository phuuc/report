package repos

import (
	"context"

	"github.com/finnpn/overview/infras/usecases/models"
	"github.com/shopspring/decimal"
)

//go:generate mockgen -destination=overview_repos_mock.go -package=repos -source=overview_repos.go

type OverviewRepos interface {
	GetTotalTrips(ctx context.Context) (int, error)
	GetTotalBookings(ctx context.Context) (int, error)
	GetTotalSpends(ctx context.Context) (decimal.Decimal, error)
	GetAverageSpend(ctx context.Context) (*models.AverageSpend, error)
	GetTotalSpendsByService(ctx context.Context) (*models.TotalSpendsByService, error)
}
