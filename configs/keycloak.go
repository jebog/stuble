package configs

import (
	"github.com/Nerzal/gocloak/v12"
	"os"
)

type Keycloak struct {
	gocloak      *gocloak.GoCloak
	clientId     string
	clientSecret string
	realm        string
}

func NewKeycloak() *Keycloak {
	return &Keycloak{
		gocloak:      gocloak.NewClient(os.Getenv("KEYCLOAK_URL")),
		clientId:     os.Getenv("KEYCLOAK_CLIENT_ID"),
		clientSecret: os.Getenv("KEYCLOAK_CLIENT_SECRET"),
		realm:        os.Getenv("KEYCLOAK_REALM"),
	}
}
