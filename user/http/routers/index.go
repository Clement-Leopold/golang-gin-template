package routers

import (
	"backend-test-chenxianhao/user-management/common"
	"backend-test-chenxianhao/user-management/domains"
	"backend-test-chenxianhao/user-management/user/method"
	"backend-test-chenxianhao/user-management/user/repository"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/op/go-logging"
)

// init route group
func InitUserRouters(g *gin.RouterGroup) {
	UserEntityRoute(g)
	FollowerRoute(g)
}

var userMethodImpl domains.UserMethod
var implOnce sync.Once
var log = logging.MustGetLogger("userRouters")

func GetImpl() domains.UserMethod {
	implOnce.Do(func() {
		userMethodImpl = method.UserMethodImpl(repository.NewUserRepositoryImpl(common.DB))
	})
	return userMethodImpl
}
