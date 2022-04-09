package main

import (
	"fmt"
	"net/http"
	"testing"
)

func TestNoSurve(t *testing.T) {
	var myH myHandler
	h := NoSurf(&myH)

	switch v := h.(type) {
	case http.Handler:
	// 지정안함
	default:
		t.Error(fmt.Sprintf("type is not Handler, but is %T", v))
	}
}

func TestSessionLoad(t *testing.T) {
	var myH myHandler

	h := SessionLoad(&myH)

	switch v := h.(type) {
	case http.Handler:
	// 지정안함
	default:
		t.Error(fmt.Sprintf("type is not Handler, but is %T", v))
	}
}
