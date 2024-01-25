package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/google/go-querystring/query"
)

type Credentials struct {
	GrantType  string `url:"grant_type"`
	Username   string `url:"username"`
	Password   string `url:"password"`
	ClientId   string `url:"client_id"`   // optional
	Version    string `url:"v"`           // optional
	DeviceId   string `url:"device_id"`   // optional
	DeviceInfo string `url:"device_info"` // optional
}

type Token struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    uint   `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	PersonID     int    `json:"personId"`
}
type DecodedToken struct {
	AUTHENTICATION_APP string `json:"AUTHENTICATION_APP"`
}

func (c Credentials) Login() Token {
	fmt.Println("Logging in as", c.Username)
	return Login(c.Username, c.Password)
}

var Login = func(id string, pwd string) Token {
	// Encode credentials as body of POST request
	values, _ := query.Values(Credentials{
		GrantType: "password",
		Username:  id,
		Password:  pwd,
		Version:   apiFrontendVersion,
	})
	// fmt.Println(values.Encode())
	reader := bytes.NewReader([]byte(values.Encode()))
	// POST Request
	contentType := "application/x-www-form-urlencoded;charset=UTF-8"
	resp, err := http.Post(apiToken, contentType, reader)
	if err != nil {
		log.Fatal(err)
	}
	bytes, err := io.ReadAll(resp.Body)
	// Unmarshal response to Token struct
	var token Token
	err = json.Unmarshal(bytes, &token)
	if err != nil {
		log.Fatal(err)
	}
	token.DecodeJWT()
	return token
}

func (t *Token) DecodeJWT() {
	// b64 decode user data from jwt token
	me := strings.Split(t.AccessToken, ".")
	decoded, err := base64.StdEncoding.DecodeString(me[1] + "=")
	if err != nil {
		log.Fatal(err)
	}

	// unmarshal decoded string
	var decodedToken DecodedToken
	err = json.Unmarshal(decoded, &decodedToken)
	if err != nil {
		log.Fatal(err)
	}
	userData := []Token{}
	err = json.Unmarshal([]byte(decodedToken.AUTHENTICATION_APP), &userData)
	if err != nil || len(userData) != 1 {
		log.Fatal(err)
	}
	t.PersonID = userData[0].PersonID
}
