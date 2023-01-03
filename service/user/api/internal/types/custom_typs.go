package types

type IdaasUserLoginCache struct {
	PreferredUserName string `json:"preferred_username"`
	Name              string `json:"name"`
	Email             string `json:"email"`
	PhoneNum          string `json:"phone_number"`
	AccessToken       string `json:"access_token"`
	IdToken           string `json:"id_token"`
	ExpiresIn         int64  `json:"expires_in"`
	ExpiresAt         int64  `json:"expires_at"`
}
