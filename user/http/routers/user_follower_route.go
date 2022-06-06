package routers

import (
	"backend-test-chenxianhao/user-management/common"
	"backend-test-chenxianhao/user-management/domains"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func FollowerRoute(router *gin.RouterGroup) {
	userMethodImpl = GetImpl()
	// following someone
	router.POST("/:id/followings/:following_id", following)
	// get followers or followings
	router.GET("/:id/followings", followers)
	// get the nearest from following list
	router.GET("/nearest-following/:name", nearest)
}

func following(ctx *gin.Context) {
	id, followingId, following := ctx.Param("id"), ctx.Param("following_id"), ctx.Query("following")
	var err error
	if following == "1" {
		err = userMethodImpl.Following(ctx, id, followingId)
	} else if following == "0" {
		err = userMethodImpl.UnFollowing(ctx, id, followingId)
	}
	if err != nil {
		log.Error(err)
		ctx.JSON(http.StatusInternalServerError, common.SystemErrorResponse())
		return
	}
	ctx.JSON(http.StatusOK, common.SucResponse(nil))
}

func followers(ctx *gin.Context) {
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
		followers, err = userMethodImpl.GetFollowings(ctx, id, int16(limit), int16(offset))
	} else {
		followers, err = userMethodImpl.GetFollowers(ctx, id, int16(limit), int16(offset))
	}
	if err != nil {
		log.Error(err)
		ctx.JSON(http.StatusInternalServerError, common.SystemErrorResponse())
		return
	}
	ctx.JSON(http.StatusOK, common.SucResponse(followers))
}

func nearest(ctx *gin.Context) {
	name := ctx.Param("name")
	follower, err := userMethodImpl.GetMinimumDistanceForFollowing(ctx, name)
	if err != nil {
		log.Error(err)
		ctx.JSON(http.StatusInternalServerError, common.SystemErrorResponse())
		return
	}
	ctx.JSON(http.StatusOK, common.SucResponse(follower))
}
