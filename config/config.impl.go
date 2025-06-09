package config

type ConfigApi struct {
	ListenHost string
	ListenPort int
}

type ConfigDatabase struct {
	Host         string
	Port         int
	User         string
	Password     string
	DatabaseName string
}

type ConfigJwt struct {
	Key string
}

type Config struct {
	Database    *ConfigDatabase
	Api         *ConfigApi
	Jwt         *ConfigJwt
	RabbitMQUrl string
}
