package models

import "github.com/shopspring/decimal"

type FlightTotalBooking struct {
	TotalBookings int
}

type FlightTotalSpends struct {
	TotalSpends decimal.Decimal
}

type FlightAveragePrice struct {
	AveragePrice decimal.Decimal
}

type FlightTopAirlines struct {
	Name         string
	BookingCount int
	AveragePrice decimal.Decimal
}

type FlightBookingWindowDays struct {
	BookingWindowDays int
}
