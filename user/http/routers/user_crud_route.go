package routers

import (
	"backend-test-chenxianhao/user-management/common"
	"backend-test-chenxianhao/user-management/domains"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func UserEntityRoute(router *gin.RouterGroup) {
	userMethodImpl = GetImpl()
	//create
	router.POST("/", createUser)
	//get
	router.GET("/:id", getUser)
	//get users
	router.GET("/", getAll)
	// update user
	router.PUT("/:id", putUser)
	// delete user
	router.DELETE("/:id", deleteUser)
}

func createUser(ctx *gin.Context) {
	user := domains.User{}
	err := ctx.BindJSON(&user)
	if err != nil {
		log.Error(err)
		ctx.JSON(http.StatusBadRequest, common.ParamErrorResponse())
		return
	}
	user.Id = uuid.New().String()
	user.CreatedAt = time.Now()
	err = userMethodImpl.Create(ctx, &user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, common.SystemErrorResponse())
		return
	}
	ctx.JSON(http.StatusOK, common.SucResponse(user))

}

func getUser(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, common.ParamErrorResponse())
		return
	}
	user, err := userMethodImpl.GetByID(ctx, id)
	if err != nil {
		log.Error(err)
		ctx.JSON(http.StatusInternalServerError, common.SystemErrorResponse())
		return
	}
	ctx.JSON(http.StatusOK, common.SucResponse(user))
}

func putUser(ctx *gin.Context) {
	id := ctx.Param("id")
	user := domains.User{}
	err := ctx.BindJSON(&user)
	if err != nil {
		log.Error(err)
		ctx.JSON(http.StatusBadRequest, common.ParamError(nil))
		return
	}
	user.Id = id
	err = userMethodImpl.Update(ctx, &user)
	if err != nil {
		log.Error(err)
		ctx.JSON(http.StatusInternalServerError, common.SystemErrorResponse())
		return
	}
	ctx.JSON(http.StatusOK, common.SucResponse(user))
}

func deleteUser(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, common.Param)
		return
	}
	err := userMethodImpl.Delete(ctx, id)
	if err != nil {
		log.Error(err)
		ctx.JSON(http.StatusInternalServerError, common.SystemErrorResponse())
		return
	}
	ctx.JSON(http.StatusOK, common.SucResponse(id))
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
		log.Error(err)
		ctx.JSON(http.StatusInternalServerError, common.SystemErrorResponse())
		return
	}
	ctx.JSON(http.StatusOK, common.SucResponse(user))
}
