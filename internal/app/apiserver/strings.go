package apiserver

import "fmt"

const (
	versionSever = "v0.0.1"
	versionAPI   = "api-v1"
)

var (
	// Messages
	textAboutServerInfo = fmt.Sprintf("Orders Book | Server: %s, %s", versionSever, versionAPI)
	textStartServer     = "Server started at "

	// Routes name
	rGetOrders = "/orders"
	rMakeOrder = "/order"
)
