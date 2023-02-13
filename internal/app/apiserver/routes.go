package apiserver

func (s *APIServer) configureRoutes() {
	s.mux.HandleFunc(rMakeOrder, s.handleMakeOrder())
	s.mux.HandleFunc(rGetOrders, s.handleGetOrders())
}
