package auth

import (
 "net/http"
 "testing"
 "reflect"
)

func TestGetAPIKey(t *testing.T) {
 headers := http.Header{}
 headers.Set("Authorization", "ApiKey 123456789")

 got, err := GetAPIKey(headers)
 want := "123456789"

 if err != nil {
  t.Fatalf("expected no error, got %v", err)
 }

 if !reflect.DeepEqual(got, want) {
  t.Errorf("got %v, want %v", got, want)
 }
}
