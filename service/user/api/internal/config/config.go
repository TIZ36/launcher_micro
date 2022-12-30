package config

import (
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf

	SSO struct {
		// step1: auth endpoint to get authorization code
		AuthorizationEndPoint string
		ClientId              string
		RedirectUrl           string
		ResponseType          string
		Scope                 string
		State                 string
		// step2: token endpoint to get id token
		TokenEndpoint string
		GrantType     string
		ClientSecret  string
	}
}
