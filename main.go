package main

import (
	"auth/auth"
	"auth/config"
	"auth/credentials"
	"auth/pkce"
	_ "embed"
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"runtime"
)

//go:embed config.yaml
var configYaml []byte

func main() {
	config, err := config.LoadConfig(configYaml)
	if err != nil {
		fmt.Printf("Error occured reading config data: %s", err)
	}

	state := pkce.RandomString()
	codeVerifier := pkce.RandomString()
	codeChallenge := pkce.GenerateCodeChallenge(codeVerifier)
	authUrl := fmt.Sprintf(
		"%s?client_id=%s&redirect_uri=%s&response_type=code&state=%s&scope=openid profile read_api&code_challenge=%s&code_challenge_method=S256",
		config.Auth.AuthURL, config.Auth.ClientID, config.Auth.RedirectURI, state, codeChallenge)

	tokenCh := make(chan auth.AuthToken)
	defer close(tokenCh)

	go handleAuthRedirect(tokenCh, codeVerifier, &config.Auth)
	openBrowser(authUrl)

	token := <-tokenCh

	err = credentials.SaveCredentials(&token)
	if err != nil {
		log.Printf("Error while storing token %s", err)
	}
}

func openBrowser(url string) bool {
	var args []string
	switch runtime.GOOS {
	case "darwin":
		args = []string{"open"}
	case "windows":
		args = []string{"cmd", "/c", "start"}
	default:
		args = []string{"xdg-open"}
	}
	cmd := exec.Command(args[0], append(args[1:], url)...)
	return cmd.Start() == nil
}

func handleAuthRedirect(tokenCh chan auth.AuthToken, codeVerifier string, config *config.ConfigAuth) {

	http.HandleFunc("/auth/redirect", func(w http.ResponseWriter, r *http.Request) {
		code := r.URL.Query().Get("code")
		log.Printf("Code returned from Gitlab is %s", code)

		token, err := auth.RequestToken(code, codeVerifier, config)
		if err != nil {
			log.Printf("Error occured requesting access token %s", err)
		}

		w.Write([]byte("You have authenticated successfully. You can now close this browser window."))
		tokenCh <- *token
	})

	err := http.ListenAndServe(":7171", nil)
	if err != nil {
		log.Printf("Error occured while setting up server %s", err)
	}

}

// parameters = 'client_id=APP_ID&code=RETURNED_CODE&grant_type=authorization_code&redirect_uri=REDIRECT_URI&code_verifier=CODE_VERIFIER'
// RestClient.post 'https://gitlab.example.com/oauth/token', parameters
