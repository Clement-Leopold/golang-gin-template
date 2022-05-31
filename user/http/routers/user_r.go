package routers

import (
	"backend-test-chenxianhao/user-management/common"
	"backend-test-chenxianhao/user-management/domains"
	"backend-test-chenxianhao/user-management/user/method"
	"backend-test-chenxianhao/user-management/user/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/op/go-logging"
)

var log = logging.MustGetLogger("UserCRUD")

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
		log.Error(err)
		ctx.JSON(http.StatusBadRequest, common.ParamResponse())
		return
	}
	err = userMethodImpl.Create(ctx, &user)
	if err != nil {
		log.Error(err)
		ctx.JSON(http.StatusInternalServerError, common.SystemResponse())
	}
	ctx.JSON(http.StatusOK, common.SucResponse(user))

}

func getUser(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, common.ParamResponse())
		return
	}
	user, err := userMethodImpl.GetByID(ctx, id)
	if err != nil {
		log.Error(err)
		ctx.JSON(http.StatusInternalServerError, common.SystemResponse())
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
	}
	user.Id = id
	err = userMethodImpl.Update(ctx, &user)
	if err != nil {
		log.Error(err)
		ctx.JSON(http.StatusInternalServerError, common.SystemResponse())
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
		ctx.JSON(http.StatusInternalServerError, common.SystemResponse())
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
		ctx.JSON(http.StatusInternalServerError, common.SystemResponse())
	}
	ctx.JSON(http.StatusOK, user)
}

var userMethodImpl domains.UseMethod
