package handler

import (
	"backend-test-chenxianhao/user-management/domains"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestFollowing(t *testing.T) {
	// Setup
	gin.SetMode(gin.TestMode)

	t.Run("Expect 404 when following id is not found", func(t *testing.T) {
		id := ":following_id"

		mockUserResp := domains.User{
			Id:   "",
			Name: "",
		}

		userFunctions := new(MockUserFunctions)
		userFunctions.On("GetByID", mock.AnythingOfType("*gin.Context"), id).Return(mockUserResp, nil)

		// a response recorder for getting written http response
		rr := httptest.NewRecorder()

		// use a middleware to set context for test
		// the only claims we care about in this test
		// is the UID
		router := gin.Default()

		NewHandler(&Config{
			R:             router,
			UserFunctions: userFunctions,
		})

		request, err := http.NewRequest(http.MethodPost, "/v1/users/:id/followings/:following_id", nil)
		assert.NoError(t, err)

		router.ServeHTTP(rr, request)

		assert.NoError(t, err)

		assert.Equal(t, http.StatusNotFound, rr.Code)
		userFunctions.AssertExpectations(t) // assert that GetByID was called
	})
}
