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

// Payload
type Payload struct {
	Ok     bool   `json:"ok"`
	Result Result `json:"result"`
}

type Result struct {
	HttpStatus  int    `json:"http_status"`
	Description string `json:"description"`
}
