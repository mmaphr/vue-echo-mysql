package main

import (
	// "fmt"
	// "encoding/json"
	// "bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

var id int

func TestGetAllMember(t *testing.T) {
	r := Router()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/member", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotNil(t, w.Body)
}

func TestGetAllGender(t *testing.T) {
	r := Router()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/gender", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotNil(t, w.Body)
}

// func TestCreateGender(t *testing.T) {
// 	r := Router()

// 	body := []byte(`{
// 		"gender_id": 100,
// 		"gendername": "Test"
// 	}`)

// 	w := httptest.NewRecorder()
// 	req, err := http.NewRequest(http.MethodPost, "/gender", bytes.NewBuffer(body))
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	r.ServeHTTP(w, req)
	
// 	assert.Equal(t, http.StatusOK, w.Code)
// 	assert.NotNil(t, w.Body)

// }

// func TestCreateMember(t *testing.T) {
// 	r := Router()

// 	body := []byte(`{
// 		"id": 100,
// 		"firstname": "FNTest",
// 		"lastname": "LNTest",
// 		"age": 22,
// 		"gender_id": 3
// 	}`)

// 	w := httptest.NewRecorder()
// 	req, err := http.NewRequest(http.MethodPost, "/member", bytes.NewBuffer(body))
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	r.ServeHTTP(w, req)
	
// 	assert.Equal(t, http.StatusOK, w.Code)
// 	assert.NotNil(t, w.Body)

// }

func TestGetMember(t *testing.T) {
	r := Router()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/member/1", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotNil(t, w.Body)
}

func TestGetGender(t *testing.T) {
	r := Router()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/gender/1", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotNil(t, w.Body)
}

// func TestDeleteMember(t *testing.T) {
// 	r := Router()
// 	w := httptest.NewRecorder()
// 	req, _ := http.NewRequest(http.MethodDelete, "/member/100", nil)
// 	r.ServeHTTP(w, req)

// 	assert.Equal(t, http.StatusNoContent, w.Code)
// }

