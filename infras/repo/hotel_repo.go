package repo

import (
	"context"
	"database/sql"

	"github.com/finnpn/overview/infras/usecases/models"
	"github.com/shopspring/decimal"
)

type HotelRepo struct {
	db *sql.DB
}

func NewHotelRepo(db *sql.DB) *HotelRepo {
	return &HotelRepo{
		db: db,
	}
}

func (r *HotelRepo) GetHotelTotalBookings(ctx context.Context) (*models.HotelTotalBookings, error) {
	var (
		count = 0
		stm   = "Select Count (booking_number) from report_booking_overview where supplier_type=$1"
	)

	row := r.db.QueryRow(stm, "HOTEL")
	err := row.Scan(&count)
	if err != nil {
		return nil, err
	}
	return &models.HotelTotalBookings{
		TotalBookings: count,
	}, nil
}
func (r *HotelRepo) GetHotelTotalSpends(ctx context.Context) (*models.HotelTotalSpends, error) {
	var (
		sum = float64(0)
		stm = "Select Sum (payment_total_amount) from report_booking_overview where supplier_type=$1"
	)

	row := r.db.QueryRow(stm, "HOTEL")
	err := row.Scan(&sum)
	if err != nil {
		return nil, err
	}
	return &models.HotelTotalSpends{
		TotalSpends: decimal.NewFromFloat(sum),
	}, nil
}

func (r *HotelRepo) GetHotelAveragePrice(ctx context.Context) (*models.HotelAveragePrice, error) {
	var (
		sum = float64(0)
		stm = "select Avg(payment_total_amount/ nights) from report_booking_overview where nights != 0"
	)

	row := r.db.QueryRow(stm)
	err := row.Scan(&sum)
	if err != nil {
		return nil, err
	}
	return &models.HotelAveragePrice{
		AveragePrice: decimal.NewFromFloat(sum),
	}, nil
}

func (r *HotelRepo) GetHotelTopBooked(ctx context.Context) ([]*models.HotelTopBooked, error) {
	var (
		name         string
		bookingCount int
		averagePrice float64
		records      []*models.HotelTopBooked
		stm          = "select supplier_name ,count(*) as booking_count, avg(payment_total_amount) as average_price from report_booking_overview where supplier_type = 'HOTEL' group by supplier_name order by booking_count desc "
	)

	rows, err := r.db.Query(stm)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&name, &bookingCount, &averagePrice)
		if err != nil {
			return nil, err
		}
		records = append(records, &models.HotelTopBooked{
			Name:         name,
			BookingCount: bookingCount,
			AveragePrice: decimal.NewFromFloat(averagePrice),
		})
	}
	return records, nil
}
