package httputil

import (
	"testing"

	"encoding/json"
	"net/http"
)

func TestNewError(t *testing.T) {
	r, err := http.NewRequest("GET", "/test", nil)
	r.Host = "localhost"
	if err != nil {
		t.Fatal(err)
	}

	marshaledErr := NewError(r, "testError").Bytes()

	var errJSON Error
	err = json.Unmarshal(marshaledErr, &errJSON)
	if err != nil {
		t.Fatal(err.Error())
	}

	if errJSON.Kind != "Error" {
		t.Errorf("errJSON.Kind = %v, expected %v", errJSON.Kind, "Error")
	}
	if errJSON.Error != "testError" {
		t.Errorf("errJSON.Error = %v, expected %v", errJSON.Error, "testError")
	}
	if errJSON.Links["self"].URL != "localhost/test" {
		t.Errorf("errJSON.Links[self].URL = %v, expected %v", errJSON.Links["test"].URL, "localhost/test")
	}
}

func TestErrorString(t *testing.T) {
	r, err := http.NewRequest("GET", "/test", nil)
	r.Host = "localhost"
	if err != nil {
		t.Fatal(err)
	}

	e := NewError(r, "testError")
	str := e.String()
	if len(str) == 0 {
		t.Error("len(str) = 0, expected > 0")
	}
}
