package apiserver

import "fmt"

const (
	versionSever = "v0.0.1"
	versionAPI   = "api-v1"
)

var (
	// Messages
	textAboutServerInfo = fmt.Sprintf("Orders Book | Server: %s, %s", versionSever, versionAPI)
	textStartServer     = "Server started at"
	// make order
	textErrorNoRequiredParams      = "Required parameters are missing: email, room, from, to"
	textErrorNoAvailableRooms      = "No available rooms"
	textErrorInvalidDateFormatFrom = "Invalid date format: from"
	textErrorInvalidDateFormatTo   = "Invalid date format: to"
	textErrorDateBookingConflict   = "Date booking conflict"
	textOkOrderCreated             = "The order has been placed"
	// get orders
	textErrorNoRequiredEmail = "Required parameters are missing: email"
	textOkRequestSuccessful  = "Request successful"

	// Routes name
	rGetOrders = "/orders"
	rMakeOrder = "/order"
)
