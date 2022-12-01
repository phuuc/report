package models

import "github.com/shopspring/decimal"

type HotelTotalBookings struct {
	TotalBookings int
}
type HotelTotalSpends struct {
	TotalSpends decimal.Decimal
}
type HotelAveragePrice struct {
	AveragePrice decimal.Decimal
}
type HotelTopBooked struct {
	Name         string
	BookingCount int
	AveragePrice decimal.Decimal
}
