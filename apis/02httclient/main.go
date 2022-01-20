package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type HttpRequester interface {
	httpGet(string) (string, int, error)
}

type UsersService interface {
	getUsers() int
}

type HttpService struct {
	Url  string
}

func (service *HttpService)httpGet(path string) (string, int, error) {
	resp, err := http.Get(service.Url + path)
	if err != nil {
		return "", 500, fmt.Errorf("HTTP error: %s", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", resp.StatusCode, fmt.Errorf("expected status OK but got %q instead", resp.Status)
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", 400,fmt.Errorf("read error: %s\n", err)
	}

	return string(data), resp.StatusCode, nil
}

func (service *HttpService) GetUsers() int{
	payload, status_code, err := service.httpGet("/get")
	if err != nil {
		fmt.Printf("GET error: %s\n", err)
	}

	fmt.Println("GET response:")
	fmt.Println(payload)
	return status_code

}

func main() {

	mordorService := HttpService{"https://httpbin.org"}

	mordorService.GetUsers()

}