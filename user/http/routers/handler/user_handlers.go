package handler

import (
	"backend-test-chenxianhao/user-management/common"
	"backend-test-chenxianhao/user-management/domains"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func UserEntityRoute(router *gin.RouterGroup, h *Handler) {
	//create
	router.POST("/", h.createUser)
	//get
	router.GET("/:id", h.getUser)
	//get users
	router.GET("/", h.getAll)
	// update user
	router.PUT("/:id", h.putUser)
	// delete user
	router.DELETE("/:id", h.deleteUser)
}

func (h *Handler) createUser(ctx *gin.Context) {
	user := domains.User{}
	err := ctx.BindJSON(&user)
	if err != nil {
		log.Error(err)
		ctx.JSON(http.StatusBadRequest, common.ParamErrorResponse())
		return
	}
	user.Id = uuid.New().String()
	user.CreatedAt = time.Now()
	err = h.UserFunctions.Create(ctx, &user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, common.SystemErrorResponse())
		return
	}
	ctx.JSON(http.StatusOK, common.SucResponse(user))

}

func (h *Handler) getUser(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, common.ParamErrorResponse())
		return
	}
	user, err := h.UserFunctions.GetByID(ctx, id)
	if err != nil {
		log.Error(err)
		ctx.JSON(http.StatusInternalServerError, common.SystemErrorResponse())
		return
	}
	ctx.JSON(http.StatusOK, common.SucResponse(user))
}

func (h *Handler) putUser(ctx *gin.Context) {
	id := ctx.Param("id")
	user := domains.User{}
	err := ctx.BindJSON(&user)
	if err != nil {
		log.Error(err)
		ctx.JSON(http.StatusBadRequest, common.ParamError(nil))
		return
	}
	user.Id = id
	err = h.UserFunctions.Update(ctx, &user)
	if err != nil {
		log.Error(err)
		ctx.JSON(http.StatusInternalServerError, common.SystemErrorResponse())
		return
	}
	ctx.JSON(http.StatusOK, common.SucResponse(user))
}

func (h *Handler) deleteUser(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, common.Param)
		return
	}
	err := h.UserFunctions.Delete(ctx, id)
	if err != nil {
		log.Error(err)
		ctx.JSON(http.StatusInternalServerError, common.SystemErrorResponse())
		return
	}
	ctx.JSON(http.StatusOK, common.SucResponse(id))
}

func (h *Handler) getAll(ctx *gin.Context) {

	limit, err := strconv.Atoi(ctx.Query("limit"))
	if err != nil {
		limit = 10
	}
	offset, err := strconv.Atoi(ctx.Query("offset"))
	if err != nil {
		offset = 0
	}
	user, err := h.UserFunctions.Get(ctx, int16(limit), int16(offset))
	if err != nil {
		log.Error(err)
		ctx.JSON(http.StatusInternalServerError, common.SystemErrorResponse())
		return
	}
	ctx.JSON(http.StatusOK, common.SucResponse(user))
}
