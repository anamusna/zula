package api_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	zulaAPI "github.com/reading-tribe/zula/src/api/pkg/api"
	"github.com/stretchr/testify/assert"
)

func TestGetUser(t *testing.T) {
	// Setup
	zulaAPI.SeedDB()
	server := zulaAPI.NewServer(nil)
	defer server.MongoContextCancel()

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	c := server.EchoServer.NewContext(req, rec)
	c.SetPath(server.Constants.UsersAPIEndpoint + "/:id")
	c.SetParamNames("id")
	c.SetParamValues(zulaAPI.DefaultUserId)

	if assert.NoError(t, server.GetUser(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestListUsers(t *testing.T) {
	// Setup
	zulaAPI.SeedDB()
	server := zulaAPI.NewServer(nil)
	defer server.MongoContextCancel()

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	c := server.EchoServer.NewContext(req, rec)
	c.SetPath(server.Constants.UsersAPIEndpoint)

	if assert.NoError(t, server.ListUsers(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}
