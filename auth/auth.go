package auth

import (
	"auth/config"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
	"time"
)

func RequestToken(code string, codeVerifier string, config *config.ConfigAuth) (*AuthToken, error) {

	form := url.Values{
		"client_id":     []string{config.ClientID},
		"code":          []string{code},
		"grant_type":    []string{"authorization_code"},
		"redirect_uri":  []string{config.RedirectURI},
		"code_verifier": []string{codeVerifier},
	}

	log.Printf("Form is %v", form)

	resp, err := http.PostForm(config.TokenURL, form)
	if err != nil {
		return nil, err
	}

	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	log.Printf("response with token %s", string(respBytes))

	at := AuthToken{}

	err = json.Unmarshal(respBytes, &at)
	if err != nil {
		return nil, err
	}

	at.CalcExpiresDate()
	at.CodeVerifier = codeVerifier

	return &at, nil
}

type AuthToken struct {
	AccessToken  string    `json:"access_token"`
	RefreshToken string    `json:"refresh_token"`
	ExpiresIn    int       `json:"expires_in"`
	ExpiresDate  time.Time `json:"expires_date"`
	CodeVerifier string    `json:"code_verifier"`
}

func (t *AuthToken) CalcExpiresDate() {
	t.ExpiresDate = time.Now().Add(time.Second * time.Duration(t.ExpiresIn))
}
