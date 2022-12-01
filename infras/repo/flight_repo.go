package repo

import (
	"context"
	"database/sql"

	"github.com/finnpn/overview/infras/usecases/models"
	"github.com/shopspring/decimal"
)

type FlightRepo struct {
	db *sql.DB
}

func NewFlightRepo(db *sql.DB) *FlightRepo {
	return &FlightRepo{
		db: db,
	}
}

func (r *FlightRepo) GetFlightTotalBooking(ctx context.Context) (*models.FlightTotalBooking, error) {
	var (
		count = 0
		stm   = "Select Count (booking_number) from report_booking_overview where supplier_type=$1"
	)

	row := r.db.QueryRow(stm, "AIR")
	err := row.Scan(&count)
	if err != nil {
		return nil, err
	}
	return &models.FlightTotalBooking{
		TotalBookings: count,
	}, nil
}

func (r *FlightRepo) GetFlightTotalSpends(ctx context.Context) (*models.FlightTotalSpends, error) {
	var (
		sum = float64(0)
		stm = "Select Sum (payment_total_amount) from report_booking_overview where supplier_type=$1"
	)

	row := r.db.QueryRow(stm, "AIR")
	err := row.Scan(&sum)
	if err != nil {
		return nil, err
	}
	return &models.FlightTotalSpends{
		TotalSpends: decimal.NewFromFloat(sum),
	}, nil
}

func (r *FlightRepo) GetFlightAveragePrice(ctx context.Context) (*models.FlightAveragePrice, error) {
	var (
		sum = float64(0)
		stm = "Select avg (payment_total_amount) from report_booking_overview where supplier_type=$1"
	)

	row := r.db.QueryRow(stm, "AIR")
	err := row.Scan(&sum)
	if err != nil {
		return nil, err
	}
	return &models.FlightAveragePrice{
		AveragePrice: decimal.NewFromFloat(sum),
	}, nil
}

func (r *FlightRepo) GetFlightTopAirlines(ctx context.Context) ([]*models.FlightTopAirlines, error) {
	var (
		name         string
		bookingCount int
		averagePrice float64
		records      []*models.FlightTopAirlines
		stm          = "select supplier_name ,count(*) as booking_count, avg(payment_total_amount) as average_price from report_booking_overview where supplier_type='AIR' group by supplier_name order by booking_count desc"
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
		records = append(records, &models.FlightTopAirlines{
			Name:         name,
			BookingCount: bookingCount,
			AveragePrice: decimal.NewFromFloat(averagePrice),
		})
	}
	return records, nil
}

func (r *FlightRepo) GetFlightBookingWindowDays(ctx context.Context) (*models.FlightBookingWindowDays, error) {
	var (
		days = float64(0)
		stm  = "select AVG(EXTRACT(day From departure_date- booking_date)) as days from report_booking_overview;"
	)

	row := r.db.QueryRow(stm)
	err := row.Scan(&days)
	if err != nil {
		return nil, err
	}
	return &models.FlightBookingWindowDays{
		BookingWindowDays: int(days),
	}, nil
}
