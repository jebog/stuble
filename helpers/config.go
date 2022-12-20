package helpers

import (
	"github.com/Nerzal/gocloak/v12"
	"os"
)

type Config struct {
	BasePath string
	Keycloak *gocloak.GoCloak
}

func NewConfig() *Config {
	c := &Config{}
	c.BasePath = "/api"

	//Keycloak Config
	c.Keycloak = gocloak.NewClient(os.Getenv("KEYCLOAK_URL"))

	return c
}
