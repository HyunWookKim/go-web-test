package myapp

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

// func TestIndexPathHanler(t *testing.T) {
// 	assert.New(t)

// 	res := httptest.NewRecorder()
// 	req := httptest.NewRequest("GET", "/", nil)

// 	indexHandler(res, req)

// 	assert.Equal(http.StatusOK, res.Code)
// 	data, _ := ioutil.ReadAll(res.Body)

// 	assert.Equal("Hello World", string(data))

// 	// if res.Code != http.StatusOK {
// 	// 	t.Fatal("Failed!!!", res.Code)
// 	// }
// }

func TestBarHanler_WithoutName(t *testing.T) {
	assert.New(t)

	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/bar", nil)

	mux := NewHttpHandler()
	mux.ServeHTTP(res, req)

	assert.Equal(http.StatusOK, res.Code)
	data, _ := ioutil.ReadAll(res.Body)

	fmt.Println(string(data))

	assert.Equal("Hello World!!!", string(data))

	// if res.Code != http.StatusOK {
	// 	t.Fatal("Failed!!!", res.Code)
	// }
}
