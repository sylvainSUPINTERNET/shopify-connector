package types

import (
	"golang.org/x/oauth2/google"
	"golang.org/x/oauth2"
	"io/ioutil"
	"errors"
	"context"
)

const firebaseScope = "https://www.googleapis.com/auth/firebase.messaging"

type tokenProvider struct {
    tokenSource oauth2.TokenSource
}

// newTokenProvider function to get token for fcm-send
func newTokenProvider(credentialsLocation string) (*tokenProvider, error) {
    jsonKey, err := ioutil.ReadFile(credentialsLocation)
    if err != nil {
        return nil, errors.New("fcm: failed to read credentials file at: " + credentialsLocation)
    }
    cfg, err := google.JWTConfigFromJSON(jsonKey, firebaseScope)
    if err != nil {
        return nil, errors.New("fcm: failed to get JWT config for the firebase.messaging scope")
    }
    ts := cfg.TokenSource(context.Background())
    return &tokenProvider{
        tokenSource: ts,
    }, nil
}

func (src *tokenProvider) token() (string, error) {
    token, err := src.tokenSource.Token()
    if err != nil {
        return "", errors.New("fcm: failed to generate Bearer token")
    }
    return token.AccessToken, nil
}
