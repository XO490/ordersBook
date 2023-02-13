package apiserver

type void struct{}

var AvailableRooms = map[string]void{}

type Order struct {
	Room      string
	UserEmail string
	From      string
	To        string
}

var ActualOrders = make([]Order, 0)
