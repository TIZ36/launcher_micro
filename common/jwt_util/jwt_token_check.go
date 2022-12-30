package jwt_util

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
)

type JwtPK struct {
	Kty string `json:"kty"`
	E   string `json:"e"`
	Use string `json:"use"`
	Kid string `json:"kid"`
	N   string `json:"n"`
}

type IdaasPK struct {
	Keys []JwtPK `json:"keys"`
}

func HodoniIdassPk() IdaasPK {
	return IdaasPK{
		Keys: []JwtPK{
			{
				Kty: "RSA",
				E:   "AQAB",
				Use: "sig",
				Kid: "KEY9LHDLNgaLnf71EmyeVqzBN78jMjp1WnwN",
				N: "xZ" +
					"-iQh1ZH0_kQTsDwSMQczhm6ZLnJiFYu4MUWAsbGBUBJ3THKtaDibrCplHeZe1IbppSrRBX3WAPZn7r5BiwWQo5KHYIj1NVW4Cbm1QOV8AefXVnhnkVy10jzhuur-Rc_K-bMFdqbfEL7XldVPTOFXaRDd7SVwo4Zspc04g3pK42wXXyZD49HxL8gFrSwHzCzBOzEIswkjyRXdBTPWE5sdeIT1R5mpmgIzNA1bL8sXjX5Ag-2P6b6onMPIP7KlZPdlru7gUMo1xCCExDmh7CQ-6k3SlSNe3h3YHV9XJpW3B8vi1Jot2ore3Tr4pAbbYaKK30Gp0-iH-3lXi2BuJzkw",
			},
		},
	}
}

func ParseJwtToken(tokenStr string) {

	claim := jwt.MapClaims{}

	token, _ := jwt.ParseWithClaims(tokenStr, claim, nil)

	fmt.Println(token)
	fmt.Println(claim)

}
