package routers

import (
	"backend-test-chenxianhao/user-management/common"
	"backend-test-chenxianhao/user-management/domains"
	"backend-test-chenxianhao/user-management/user/method"
	"backend-test-chenxianhao/user-management/user/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func UserEntityRoute(router *gin.RouterGroup) {
	userMethodImpl = method.UserMethodImpl(repository.NewUserRepositoryImpl(common.DB))
	router.POST("/", createUser)
	router.GET("/:id", getUser)
	router.GET("/", getAll)
	router.PUT("/:id", putUser)
	router.DELETE("/:id", deleteUser)
}

func createUser(ctx *gin.Context) {
	user := domains.User{}
	err := ctx.BindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, common.Param)
		return
	}
	userMethodImpl.Create(ctx, &user)
}

func getUser(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, common.Param)
		return
	}
	user, _ := userMethodImpl.GetByID(ctx, id)
	ctx.JSON(http.StatusOK, user)
}

func putUser(ctx *gin.Context) {
	id := ctx.Param("id")
	user := domains.User{}
	err := ctx.BindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}
	user.Id = id
	userMethodImpl.Update(ctx, &user)
}

func deleteUser(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, common.Param)
		return
	}
	userMethodImpl.Delete(ctx, id)
}

func getAll(ctx *gin.Context) {

	limit, err := strconv.Atoi(ctx.Query("limit"))
	if err != nil {
		limit = 10
	}
	offset, err := strconv.Atoi(ctx.Query("offset"))
	if err != nil {
		offset = 0
	}
	user, err := userMethodImpl.Get(ctx, int16(limit), int16(offset))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}
	ctx.JSON(http.StatusOK, user)
}

var userMethodImpl domains.UseMethod
