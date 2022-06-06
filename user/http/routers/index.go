package routers

import (
	"backend-test-chenxianhao/user-management/common"
	"backend-test-chenxianhao/user-management/domains"
	"backend-test-chenxianhao/user-management/user/impl"
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

var userMethodImpl domains.UserFunctions
var implOnce sync.Once
var log = logging.MustGetLogger("userRouters")

func GetImpl() domains.UserFunctions {
	implOnce.Do(func() {
		userMethodImpl = impl.UserMethodImpl(repository.NewUserRepositoryImpl(common.DB))
	})
	return userMethodImpl
}
