package repo

import (
	"context"
	"database/sql"

	"github.com/finnpn/overview/infras/usecases/models"
	"github.com/shopspring/decimal"
)

type OverviewRepo struct {
	db *sql.DB
}

func NewOverviewRepo(db *sql.DB) *OverviewRepo {
	return &OverviewRepo{
		db: db,
	}
}

func (r *OverviewRepo) GetTotalTrips(ctx context.Context) (int, error) {
	var (
		count = 0
		stm   = "Select Sum(case when round_type =$1 then 2 else 1 end) as totaltrips from report_booking_overview"
	)

	row := r.db.QueryRow(stm, "RoundTrip")
	err := row.Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (r *OverviewRepo) GetTotalBookings(ctx context.Context) (int, error) {
	var (
		count = 0
		stm   = "Select Count (booking_number) from report_booking_overview"
	)

	row := r.db.QueryRow(stm)
	err := row.Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (r *OverviewRepo) GetTotalSpends(ctx context.Context) (decimal.Decimal, error) {
	var (
		sum        = float64(0)
		decimalSum decimal.Decimal
		stm        = "Select sum(payment_total_amount) from report_booking_overview"
	)

	row := r.db.QueryRow(stm)
	err := row.Scan(&sum)
	decimalSum = decimal.NewFromFloat(sum)
	if err != nil {
		return decimal.Zero, err
	}
	return decimalSum, nil
}

func (r *OverviewRepo) GetAverageSpend(ctx context.Context) (*models.AverageSpend, error) {
	var (
		flight, hotel, tours, combo float64
		stm                         = "Select avg(case when supplier_type = 'AIR' then payment_total_amount else null end) as flight , avg(case when supplier_type = 'HOTEL' then payment_total_amount else null end) as hotel, avg(case when supplier_type = 'TOURS' then payment_total_amount else null end) as tours ,avg(case when supplier_type = 'COMBO' then payment_total_amount else null end) as combo from report_booking_overview"
	)
	row := r.db.QueryRow(stm)
	err := row.Scan(&flight, &hotel, &tours, &combo)
	if err != nil {
		return nil, err
	}

	averageSpend := &models.AverageSpend{
		Flight:     int64(flight),
		Hotel:      int64(hotel),
		Combo:      int64(combo),
		TourTicket: int64(tours),
	}

	return averageSpend, nil
}

func (r *OverviewRepo) GetTotalSpendsByService(ctx context.Context) (*models.TotalSpendsByService, error) {
	var (
		flight, hotel, tours, combo float64
		stm                         = "Select sum(case when supplier_type = 'AIR' then payment_total_amount else null end) as flight , sum(case when supplier_type = 'HOTEL' then payment_total_amount else null end) as hotel, sum(case when supplier_type = 'TOURS' then payment_total_amount else null end) as tours , sum(case when supplier_type = 'COMBO' then payment_total_amount else null end) as combo from report_booking_overview"
	)
	row := r.db.QueryRow(stm)
	err := row.Scan(&flight, &hotel, &tours, &combo)
	if err != nil {
		return nil, err
	}

	totalSpends := &models.TotalSpendsByService{
		Flight:     int64(flight),
		Hotel:      int64(hotel),
		Combo:      int64(combo),
		TourTicket: int64(tours),
	}

	return totalSpends, nil
}
func (r *OverviewRepo) GetTotalEmployeeSpends(ctx context.Context) (decimal.Decimal, error) {
	var (
		sum        = float64(0)
		decimalSum decimal.Decimal
		stm        = "Select sum(payment_total_amount) from report_booking_overview"
	)

	row := r.db.QueryRow(stm)
	err := row.Scan(&sum)
	decimalSum = decimal.NewFromFloat(sum)
	if err != nil {
		return decimal.Zero, err
	}
	return decimalSum, nil
}
