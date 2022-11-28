package accesscontrol

import (
	"encoding/base64"
	"fmt"

	httputil "github.com/akuity/api-client-go/pkg/utils/http"
)

type ClientCredential interface {
	Scheme() string
	Credential() string
}

type APIKeyCredential struct {
	ID     string
	Secret string
}

func NewAPIKeyCredential(id, secret string) ClientCredential {
	return APIKeyCredential{
		ID:     id,
		Secret: secret,
	}
}

func (a APIKeyCredential) Scheme() string {
	return httputil.AuthSchemeBasic
}

func (a APIKeyCredential) Credential() string {
	return base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", a.ID, a.Secret)))
}

type UserCredential struct {
	Token string
}

func NewUserCredential(token string) ClientCredential {
	return UserCredential{
		Token: token,
	}
}

func (u UserCredential) Scheme() string {
	return httputil.AuthSchemeBearer
}

func (u UserCredential) Credential() string {
	return u.Token
}
