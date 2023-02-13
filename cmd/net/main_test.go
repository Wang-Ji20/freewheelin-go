package main

import (
	"fmt"
	"io"
	"net/http"
	"testing"
)

func TestMF(t *testing.T) {

	resp, err := http.Get("http://www.baidu.com")
	if err != nil {
		t.Errorf("cannot get")
	}
	content, err := io.ReadAll(resp.Body)

	if err != nil {
		t.Errorf("cannot readall")
	}

	fmt.Println(string(content))
}