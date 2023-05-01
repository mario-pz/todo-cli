package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"regexp"
)

func checkRegex(email string) bool {
	regex := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	return regex.MatchString(email)
}

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Task struct {
	ID          int    `json:"id"`
	UserID      int    `json:"user_id"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

func Login(user string, password string) string {

	credentials := User{
		Username: user,
		Password: password,
	}

	requestBody, err := json.Marshal(credentials)
	if err != nil {
		return err.Error()
	}

	req, err := http.NewRequest(
		"POST",
		"http://localhost:8080/api/login",
		bytes.NewBuffer(requestBody),
	)
	if err != nil {
		return err.Error()
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err.Error()
	}
	defer resp.Body.Close()

	return "Login Successful"
}

func Register(user string, password string) string {

	credentials := User{
		Username: user,
		Password: password,
	}

	requestBody, err := json.Marshal(credentials)
	if err != nil {
		return err.Error()
	}

	req, err := http.NewRequest(
		"POST",
		"http://localhost:8080/api/register",
		bytes.NewBuffer(requestBody),
	)
	if err != nil {
		return err.Error()
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err.Error()
	}
	defer resp.Body.Close()

	return "Register Successful"
}
