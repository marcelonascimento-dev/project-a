package handlers

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"

	"project-a/types"
	"project-a/util"
)

var appcontext_auth *types.ApplicationContext

func CreateAuthHandler(ctx *types.ApplicationContext) {
	appcontext_auth = ctx
}

func WithAuthenticatedUser(ctx *fiber.Ctx) error {
	authHeader := ctx.Cookies("session_id")
	if authHeader == "" {
		return fiber.NewError(fiber.StatusInternalServerError, "authorization header is missing")
	}
	sess, err := appcontext_auth.SessionStore.Get(ctx)

	if err != nil {
		panic(err)
	}

	if sess == nil {
		return fiber.NewError(fiber.StatusBadRequest, "session nil")
	}

	value, err := appcontext.SessionStore.Storage.Get(sess.ID())

	if value == nil || err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "authorization is missing")
	}

	return ctx.Next()
}

func GoogleLoginHandler(ctx *fiber.Ctx) error {
	conf, err := getGoogleOAuthConfig()

	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "server environment variables not properly configured")
	}

	url := conf.AuthCodeURL("state", oauth2.AccessTypeOffline)

	ctx.Redirect(url, fiber.StatusTemporaryRedirect)

	return nil
}

func GoogleCallbackHandler(ctx *fiber.Ctx) error {

	conf, err := getGoogleOAuthConfig()
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "server environment variables not properly configured")
	}

	code := ctx.FormValue("code")
	token, err := conf.Exchange(context.Background(), code)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "server environment variables not properly configured")
	}

	client := conf.Client(context.Background(), token)
	resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "failed to get user info")
	}

	defer resp.Body.Close()

	sess, err := appcontext_auth.SessionStore.Get(ctx)

	if err != nil {
		panic(err)
	}

	if sess == nil {
		return fiber.NewError(fiber.StatusBadRequest, "session nil")
	}

	token_session, err := createToken(sess.ID())

	if err != nil {
		panic(err)
	}

	sess.Set(sess.ID(), token_session)

	// Save session
	if err := sess.Save(); err != nil {
		panic(err)
	}

	return ctx.SendString("Usuário Logado com sucesso!")
}

func getGoogleOAuthConfig() (*oauth2.Config, error) {
	credentials, ok := util.GetGoogleOAuthCredentials()
	if !ok {
		return nil, errors.New("google oauth credentials not properly configured")
	}

	return &oauth2.Config{
		ClientID:     credentials.ClientID,
		ClientSecret: credentials.ClientSecret,
		RedirectURL:  credentials.RedirectURL,
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile"},
		Endpoint:     google.Endpoint,
	}, nil
}

func createToken(userID string) (string, error) {
	credentials, ok := util.GetGoogleOAuthCredentials()
	if !ok {
		return "", errors.New("google oauth secret key not found")
	}

	atClaims := jwt.MapClaims{
		"authorized": true,
		"user_id":    userID,
		"exp":        time.Now().Add(time.Hour * 24).Unix(),
	}

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(credentials.ClientSecret))
	if err != nil {
		return "", err
	}

	return token, nil
}

func tokenVerifyMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	})
}
