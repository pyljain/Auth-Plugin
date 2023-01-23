package credentials

import (
	"auth/auth"
	"encoding/json"
	"os"
	"path"
)

func SaveCredentials(tokeninfo *auth.AuthToken) error {
	// Create ~/.auth directory
	homedir, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	basePath := path.Join(homedir, ".auth")
	err = os.MkdirAll(basePath, 0700)
	if err != nil {
		return err
	}

	// Create credentials.json file
	credsBytes, err := json.Marshal(tokeninfo)
	if err != nil {
		return err
	}

	filePath := path.Join(basePath, "credentials.json")
	err = os.WriteFile(filePath, credsBytes, 0600)
	if err != nil {
		return err
	}

	return nil
}
