package models

type HotelTotalBookings struct {
	TotalBookings int `json:"total_bookings"`
}

type HotelTotalSpends struct {
	TotalSpends string `json:"total_spends"`
}

type HotelAveragePrice struct {
	AveragePrice string `json:"average_price"`
}
type HotelTopBooked struct {
	Name         string `json:"name"`
	BookingCount int    `json:"booking_count"`
	AveragePrice string `json:"average_price"`
}
