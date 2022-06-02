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
	router.POST("/:id/followings/:following_id", following)
	// get followers or followings
	router.GET("/:id/followings", followers)
}

func following(ctx *gin.Context) {
	id := ctx.Param("id")
	followingId := ctx.Param("following_id")
	following := ctx.Query("following")
	var err error
	if following == "1" {
		err = userMethodImpl.Following(ctx, id, followingId)
	} else {
		err = userMethodImpl.UnFollowing(ctx, id, followingId)
	}
	if err != nil {
		log.Error(err)
		ctx.JSON(http.StatusInternalServerError, common.SystemResponse())
		return
	}
	ctx.JSON(http.StatusOK, common.SucResponse(nil))
}

func followers(ctx *gin.Context) {
	id := ctx.Param("id")
	following := ctx.Query("following")
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
		followers, err = userMethodImpl.GetFollowers(ctx, id, int16(limit), int16(offset))
	} else {
		followers, err = userMethodImpl.GetFollowings(ctx, id, int16(limit), int16(offset))
	}
	if err != nil {
		log.Error(err)
		ctx.JSON(http.StatusInternalServerError, common.SystemResponse())
		return
	}
	ctx.JSON(http.StatusOK, common.SucResponse(followers))
}
