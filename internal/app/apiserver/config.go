package apiserver

type Config struct {
	BindAddr  string   `json:"bind_addr"`
	RoomTypes []string `json:"room_types"`
}

func NewConfig() *Config {
	return &Config{
		BindAddr:  ":8080",
		RoomTypes: []string{"econom", "standart", "lux"},
	}
}

func AvailableRoomsInit(config *Config) {
	for _, e := range config.RoomTypes {
		AvailableRooms[e] = void{}
	}
}
