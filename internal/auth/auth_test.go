package auth

import (
  "reflect"
  "testing"
  "net/http"
)

func TestGetAPIKey(t *testing.T) {
  header := http.Header{}
  header.Add("Authorization", "ApiKey hello-world")
  got, err := GetAPIKey(header)
  if err != nil {
    t.Fatalf("unexpected error %v", err)
  }
  want := "hello-world"
  if !reflect.DeepEqual(want, got) {
    t.Fatalf("expected: %v, got: %v", want, got)
  }
}
