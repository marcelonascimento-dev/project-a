package util

import (
	"os"
)

const (
	GoogleClientIdKey          = "google-project-a-oauth-client-id"
	GoogleClientSecretKey      = "google-project-a-oauth-client-secret"
	GoogleRedirectUrlKey       = "google-project-a-oauth-redirect-url"
	PostgresqlConnectionString = "postgresql-connectionstring"
)

type GoogleOAuthCredentials struct {
	ClientID     string
	ClientSecret string
	RedirectURL  string
}

func GetGoogleOAuthCredentials() (GoogleOAuthCredentials, bool) {
	v1, ok1 := os.LookupEnv(GoogleClientIdKey)
	v2, ok2 := os.LookupEnv(GoogleClientSecretKey)
	v3, ok3 := os.LookupEnv(GoogleRedirectUrlKey)

	return GoogleOAuthCredentials{
		ClientID:     v1,
		ClientSecret: v2,
		RedirectURL:  v3,
	}, ok1 && ok2 && ok3
}

func GetPostresqlConnectionString() (string, bool) {
	v, ok := os.LookupEnv(PostgresqlConnectionString)

	return v, ok
}
