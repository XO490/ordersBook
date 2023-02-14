package apiserver

import (
	"encoding/json"
	"net/http"
	"time"
)

func loggingHandle(s *APIServer, r *http.Request) {
	s.logInfo(r.Host, r.Method, r.URL.Query(), r.RemoteAddr, r.UserAgent())
}

func (s *APIServer) handleMakeOrder() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		loggingHandle(s, r)
		w.Header().Set("Content-Type", "application/json")

		var payload Payload
		var err error
		var jsonData []byte
		var fromTime, toTime time.Time

		params := r.URL.Query()
		email := params.Get("email")
		room := params.Get("room")
		from := params.Get("from")
		to := params.Get("to")

		sendResult := func(ok bool, httpStatus int, desc string) {
			w.WriteHeader(httpStatus)
			payload = Payload{
				Ok: ok,
				Result: Result{
					HttpStatus:  httpStatus,
					Description: desc,
				},
			}
			if jsonData, err = json.Marshal(&payload); err != nil {
				s.logPanic(err, "handleMakeOrder")
			}
			if _, err = w.Write(jsonData); err != nil {
				s.logPanic(err, "handleMakeOrder > send")
			}
		}

		if params.Encode() == "" || email == "" || room == "" || from == "" || to == "" {
			sendResult(false, http.StatusBadRequest, textErrorNoRequiredParams)
			return
		}

		if _, isOK := AvailableRooms[room]; !isOK {
			sendResult(false, http.StatusBadRequest, textErrorNoAvailableRooms)
			return
		}

		if fromTime, err = time.Parse("2006-01-02", from); err != nil {
			sendResult(false, http.StatusBadRequest, textErrorInvalidDateFormatFrom)
			return
		}

		if toTime, err = time.Parse("2006-01-02", from); err != nil {
			sendResult(false, http.StatusBadRequest, textErrorInvalidDateFormatTo)
			return
		}

		// when all fields are verified
		newOrder := Order{
			Room:      room,
			UserEmail: email,
			From:      from,
			To:        to,
		}

		for _, order := range ActualOrders {
			currentOrderFromTime, _ := time.Parse("2006-01-02", order.From)
			currentOrderToTime, _ := time.Parse("2006-01-02", order.To)
			if !(currentOrderToTime.Before(fromTime) || currentOrderFromTime.After(toTime)) {
				sendResult(false, http.StatusConflict, textErrorDateBookingConflict)
				return
			}
		}

		// when everything is fine :)
		ActualOrders = append(ActualOrders, newOrder)
		sendResult(true, http.StatusCreated, textOkOrderCreated)
	}
}

func (s *APIServer) handleGetOrders() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")

		w.WriteHeader(http.StatusCreated)
		loggingHandle(s, r)
	}
}
