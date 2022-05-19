package controller_test

import (
	"fmt"
	"net/http"
	"strings"
	"testing"
)

func TestGetAllOK(t *testing.T) {
	response, err := http.Get("http://localhost:8080/favourites")
	if response.StatusCode != http.StatusOK {
		t.Errorf("Expected response code %d. Got %d\n", http.StatusOK, response.StatusCode)
	}

	if err != nil {
		t.Errorf("Encountered an error:%s", err)
	}
}

func TestGetByIDOK(t *testing.T) {
	response, err := http.Get("http://localhost:8080/favourites/1")
	if response.StatusCode != http.StatusOK {
		t.Errorf("Expected response code %d. Got %d\n", http.StatusOK, response.StatusCode)
	}

	if err != nil {
		t.Errorf("Encountered an error:%s", err)
	}
}

func TestGetByIDInvalidParam(t *testing.T) {
	response, err := http.Get("http://localhost:8080/favourites/hello")
	if response.StatusCode != http.StatusBadRequest {
		t.Errorf("Expected response code %d. Got %d\n", http.StatusBadRequest, response.StatusCode)
	}

	if err != nil {
		t.Errorf("Encountered an error:%s", err)
	}
}

func TestGetByIDNotFound(t *testing.T) {
	response, err := http.Get("http://localhost:8080/favourites/10000")
	if response.StatusCode != http.StatusNotFound {
		t.Errorf("Expected response code %d. Got %d\n", http.StatusNotFound, response.StatusCode)
	}

	if err != nil {
		t.Errorf("Encountered an error:%s", err)
	}
}

func TestCreateCreated(t *testing.T) {
	// that does not exists, duplicate key value does not violate unique constraint
	favouriteRequest := `
	{
        "user_id": 1,
        "asset_id": 1 
    }`

	response, err := http.Post("http://localhost:8080/favourites", "application/json", strings.NewReader(favouriteRequest))
	if response.StatusCode != http.StatusCreated {
		t.Errorf("Expected response code %d. Got %d\n", http.StatusCreated, response.StatusCode)
	}

	if err != nil {
		t.Errorf("Encountered an error:%s", err)
	}
}

func TestCreateBadRequestDublicateKeyValue(t *testing.T) {
	// that exists, duplicate key value violates unique constraint
	favouriteRequest := `
	{
        "user_id": 1,
        "asset_id": 1 
    }`

	response, err := http.Post("http://localhost:8080/favourites", "application/json", strings.NewReader(favouriteRequest))
	if response.StatusCode != http.StatusBadRequest {
		t.Errorf("Expected response code %d. Got %d\n", http.StatusBadRequest, response.StatusCode)
	}

	if err != nil {
		t.Errorf("Encountered an error:%s", err)
	}
}

func TestCreateBadRequestValidationErr(t *testing.T) {
	//
	favouriteRequest := `
	{
        "user_id": 1
    }`

	response, err := http.Post("http://localhost:8080/favourites", "application/json", strings.NewReader(favouriteRequest))
	if response.StatusCode != http.StatusBadRequest {
		t.Errorf("Expected response code %d. Got %d\n", http.StatusBadRequest, response.StatusCode)
	}

	if err != nil {
		t.Errorf("Encountered an error:%s", err)
	}
}

func TestRemoveFavFromUserNoContent(t *testing.T) {
	var client http.Client

	req, err := http.NewRequest(http.MethodDelete, "http://localhost:8080/users/1/favourites/1", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNoContent {
		t.Errorf("Expected response code %d. Got %d\n", http.StatusNoContent, resp.StatusCode)
	}

	if err != nil {
		t.Errorf("Encountered an error:%s", err)
	}
}

func TestRemoveFavFromUserUserNotFound(t *testing.T) {
	var client http.Client

	req, err := http.NewRequest(http.MethodDelete, "http://localhost:8080/users/10000/favourites/1", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNotFound {
		t.Errorf("Expected response code %d. Got %d\n", http.StatusNotFound, resp.StatusCode)
	}

	if err != nil {
		t.Errorf("Encountered an error:%s", err)
	}
}

func TestRemoveFavFromUserAssetNotFound(t *testing.T) {
	var client http.Client

	req, err := http.NewRequest(http.MethodDelete, "http://localhost:8080/users/1/favourites/10000", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNotFound {
		t.Errorf("Expected response code %d. Got %d\n", http.StatusNotFound, resp.StatusCode)
	}

	if err != nil {
		t.Errorf("Encountered an error:%s", err)
	}
}
