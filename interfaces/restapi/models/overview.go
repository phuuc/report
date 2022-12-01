package models

type TotalTrips struct {
	TotalTrips int `json:"total_trips"`
}
type TotalBookings struct {
	TotalBookings int `json:"total_bookings"`
}

type TotalSpends struct {
	TotalSpends string `json:"total_spends"`
}

type AverageSpend struct {
	Flight     string `json:"flight"`
	Hotel      string `json:"hotel"`
	Combo      string `json:"combo"`
	TourTicket string `json:"tour_ticket"`
}

type TotalSpendsByService struct {
	Flight     string `json:"flight"`
	Hotel      string `json:"hotel"`
	Combo      string `json:"combo"`
	TourTicket string `json:"tour_ticket"`
}
