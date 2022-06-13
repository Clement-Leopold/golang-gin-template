package handler

import (
	"backend-test-chenxianhao/user-management/common"
	"backend-test-chenxianhao/user-management/domains"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func FollowerRoute(router *gin.RouterGroup, h *Handler) {
	// following someone
	router.POST("/:id/followings/:following_id", h.following)
	// get followers or followings
	router.GET("/:id/followings", h.followers)
	// get the nearest from following list
	router.GET("/nearest-following/:name", h.nearest)
}

func (h *Handler) following(ctx *gin.Context) {
	id, followingId, following := ctx.Param("id"), ctx.Param("following_id"), ctx.Query("following")
	var err error
	followingUser, err := h.UserFunctions.GetByID(ctx, followingId)
	if followingUser.Id == "" {
		ctx.JSON(http.StatusNotFound, common.SystemErrorResponse())
		return
	}
	if following == "1" {
		err = h.UserFunctions.Following(ctx, id, followingId)
	} else if following == "0" {
		err = h.UserFunctions.UnFollowing(ctx, id, followingId)
	}
	if err != nil {
		log.Error(err)
		ctx.JSON(http.StatusInternalServerError, common.SystemErrorResponse())
		return
	}
	ctx.JSON(http.StatusOK, common.SucResponse(nil))
}

func (h *Handler) followers(ctx *gin.Context) {
	id, following := ctx.Param("id"), ctx.Query("following")
	limit, err := strconv.Atoi(ctx.Query("limit"))
	if err != nil {
		limit = 10
	}
	offset, err := strconv.Atoi(ctx.Query("offset"))
	if err != nil {
		offset = 0
	}
	var followers []domains.Follower
	if following == "1" {
		followers, err = h.UserFunctions.GetFollowings(ctx, id, int16(limit), int16(offset))
	} else {
		followers, err = h.UserFunctions.GetFollowers(ctx, id, int16(limit), int16(offset))
	}
	if err != nil {
		log.Error(err)
		ctx.JSON(http.StatusInternalServerError, common.SystemErrorResponse())
		return
	}
	ctx.JSON(http.StatusOK, common.SucResponse(followers))
}

func (h *Handler) nearest(ctx *gin.Context) {
	name := ctx.Param("name")
	follower, err := h.UserFunctions.GetMinimumDistanceForFollowing(ctx, name)
	if err != nil {
		log.Error(err)
		ctx.JSON(http.StatusInternalServerError, common.SystemErrorResponse())
		return
	}
	ctx.JSON(http.StatusOK, common.SucResponse(follower))
}
