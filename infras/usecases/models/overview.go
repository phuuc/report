package models

type AverageSpend struct {
	Flight     int64
	Hotel      int64
	Combo      int64
	TourTicket int64
}

type TotalSpendsByService struct {
	Flight     int64
	Hotel      int64
	Combo      int64
	TourTicket int64
}
