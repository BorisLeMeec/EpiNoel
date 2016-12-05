package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
	"strings"
)

func logToIntra() (*http.Cookie, error) {
	resp, err := http.Post("https://intra.epitech.eu/login", "application/x-www-form-urlencoded", bytes.NewBufferString("login="+os.Getenv("LOGIN_INTRA")+"&password="+os.Getenv("PASSWD_INTRA")))
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 401 {
		return nil, errors.New(resp.Status)
	}
	log.Printf("%s\n", resp.Header)
	cookieValue := strings.Split(strings.Split(resp.Header["Set-Cookie"][0], ";")[0], "=")[1]
	return &http.Cookie{Name: "PHPSESSID", Value: cookieValue}, nil
}

func getEpiURL(login string) (url string) {
	url = "https://intra.epitech.eu/user/" + login + "/?format=json"
	return
}

type pic struct {
	URL string `json:"picture"`
}

func getPicProfile(login string) (url string) {
	var picture pic
	client := &http.Client{}

	profile := getEpiURL(login)
	log.Printf("Fetch %s\n", profile)
	req, err := http.NewRequest("GET", profile, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.AddCookie(cookieSessionIntra)
	body, _ := httputil.DumpRequestOut(req, true)
	fmt.Printf("%s\n", body)
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	html, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", html)
	res.Body.Close()
	if err := json.Unmarshal(html, &picture); err != nil {
		log.Fatal(err)
	}

	return picture.URL
}
