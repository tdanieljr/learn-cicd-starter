package auth

import (
	"net/http"
	"reflect"
	"testing"
)

func TestGetAPI(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "/whatever", nil)
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Authorization", "ApiKey secret-key-123")

	key, err := GetAPIKey(req.Header)
	if err != nil {
		t.Fatal("returned an error on correct case")

	}
	if !reflect.DeepEqual(key, "secret-key-123") {
		t.Fatal("wrong key extracted")
	}
}
