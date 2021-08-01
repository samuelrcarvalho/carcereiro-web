package main

import (
	"crypto/rand"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Auth struct {
	User      string `gorm:"primaryKey"`
	Authcode  string
	Expire_in time.Time
}

func enviarCode(user string, db *gorm.DB) {
	code := geraCode(6)

	teste := &Auth{User: user, Authcode: code, Expire_in: time.Now()}

	db.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(teste)

	url := configs["rocket_url"]
	method := "POST"

	payload := strings.NewReader("channel=%40" + user + "&text=" + code)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Content-type", "application/x-www-form-urlencoded")
	req.Header.Add("X-User-Id", configs["rocket_integracao_user_id"])
	req.Header.Add("X-Auth-Token", configs["rocket_integracao_auth_token"])

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}

func geraCode(max int) string {
	table := [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}
	b := make([]byte, max)
	n, err := io.ReadAtLeast(rand.Reader, b, max)
	if n != max {
		panic(err)
	}
	for i := 0; i < len(b); i++ {
		b[i] = table[int(b[i])%len(table)]
	}
	return string(b)
}
