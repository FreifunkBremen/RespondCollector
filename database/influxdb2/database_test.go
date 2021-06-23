package influxdb

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConnect(t *testing.T) {
	assert := assert.New(t)

	conn, err := Connect(map[string]interface{}{
		"address":              "",
		"username":             "",
		"password":             "",
		"insecure_skip_verify": true,
	})
	assert.Nil(conn)
	assert.Error(err)

	conn, err = Connect(map[string]interface{}{
		"address":  "http://localhost",
		"database": "",
		"username": "",
		"password": "",
	})
	assert.Nil(conn)
	assert.Error(err)

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	}))
	defer srv.Close()

	conn, err = Connect(map[string]interface{}{
		"address":  srv.URL,
		"database": "",
		"username": "",
		"password": "",
	})

	assert.NotNil(conn)
	assert.NoError(err)
}
