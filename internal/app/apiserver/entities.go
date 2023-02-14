package apiserver

type void struct{}

var AvailableRooms = map[string]void{}

type Order struct {
	Room      string `json:"room"`
	UserEmail string `json:"user_email"`
	From      string `json:"from"`
	To        string `json:"to"`
}

var ActualOrders = make([]Order, 0)

// Payload
type Payload struct {
	Ok     bool   `json:"ok"`
	Result Result `json:"result"`
}

type Result struct {
	HttpStatus  int     `json:"http_status"`
	Description string  `json:"description"`
	Orders      []Order `json:"orders,omitempty"`
}
