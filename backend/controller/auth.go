package controller

import (
	"bytes"
	b64 "encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/labstack/echo"
)

func TestPrint(c echo.Context) error {
	fmt.Printf("ttttest")
	return c.String(http.StatusOK, "Hello, World!")
}

func TestAuth(c echo.Context) error {
	// 認証情報をbase64エンコード
	apiKey := os.Getenv("TWITTER_API_KEY")
	secKey := os.Getenv("TWITTER_API_SECRET")
	keys := apiKey + ":" + secKey

	encKeys := b64.StdEncoding.EncodeToString([]byte(keys))
	fmt.Println(encKeys)

	req, _ := http.NewRequest(
		"POST",
		"https://api.twitter.com/oauth2/token",
		bytes.NewBuffer([]byte("grant_type=client_credentials")),
	)
	req.Header.Set("Authorization", "Basic "+encKeys)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded;charset=UTF-8.")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	fmt.Println("[body] " + string(body))

	return c.String(http.StatusOK, "Hello, World!4")
}
