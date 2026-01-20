package response

import "github.com/flipped-aurora/gin-vue-admin/server/model/app"

type LoginResponse struct {
	User      app.AppUsers `json:"user"`
	Token     string       `json:"token"`
	ExpiresAt int64        `json:"expiresAt"`
}

type CasdoorResponse struct {
	AccessToken  string `json:"accessToken"`
	IDToken      string `json:"idToken"`
	RefreshToken string `json:"refreshToken"`
	TokenType    string `json:"tokenType"`
	ExpiresIn    int64  `json:"expiresIn"`
	Scope        string `json:"scope"`
}
