package config

import (
	"os"
	"strconv"
)

func NewConfig() *Config {
	listen_port := os.Getenv("API_LISTEN_PORT")
	if len(listen_port) != 4 {
		panic("API_LISTEN_PORT must be 4 digits of number")
	}

	db_port := os.Getenv("DATABASE_PORT")
	if len(listen_port) < 3 {
		panic("DATABASE_PORT must atleast 3 digits")
	}

	listen_port_int, err := strconv.Atoi(listen_port)
	if err != nil {
		panic(err)
	}

	db_port_int, err := strconv.Atoi(db_port)
	if err != nil {
		panic(err)
	}

	return &Config{
		Api: &ConfigApi{
			ListenHost: os.Getenv("API_LISTEN_HOST"),
			ListenPort: listen_port_int,
		},
		Database: &ConfigDatabase{
			Host:         os.Getenv("DATABASE_HOST"),
			Port:         db_port_int,
			User:         os.Getenv("DATABASE_USER"),
			Password:     os.Getenv("DATABASE_PASSWORD"),
			DatabaseName: os.Getenv("DATABASE_DBNAME"),
		},
		Jwt: &ConfigJwt{
			Key: os.Getenv("JWT_KEY"),
		},
	}
}
