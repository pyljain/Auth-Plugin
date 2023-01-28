package auth

import (
	"auth/config"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
)

func RefreshToken(token *AuthToken, config *config.ConfigAuth) (*AuthToken, error) {

	form := url.Values{
		"client_id":     []string{config.ClientID},
		"grant_type":    []string{"refresh_token"},
		"refresh_token": []string{token.RefreshToken},
		"redirect_uri":  []string{config.RedirectURI},
		"code_verifier": []string{token.CodeVerifier},
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
	at.CodeVerifier = token.CodeVerifier

	return &at, nil

}

// parameters = 'client_id=APP_ID&refresh_token=REFRESH_TOKEN&grant_type=refresh_token&redirect_uri=REDIRECT_URI&code_verifier=CODE_VERIFIER'
// RestClient.post 'https://gitlab.example.com/oauth/token', parameters
