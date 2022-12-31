package configs

type Config struct {
	BasePath string
}

func NewConfig() *Config {
	c := &Config{}
	c.BasePath = "/api"

	return c
}
