package models

type FlightTotalBooking struct {
	TotalBookings int `json:"total_bookings"`
}

type FlightTotalSpends struct {
	TotalSpends string `json:"total_spends"`
}

type FlightAveragePrice struct {
	AveragePrice string `json:"average_price"`
}

type FlightTopAirlines struct {
	Name         string `json:"name"`
	BookingCount int    `json:"booking_count"`
	AveragePrice string `json:"average_price"`
}

type FlightBookingWindowDays struct {
	BookingWindowDays int `json:"booking_window_days"`
}
