package main

import (
	"backend-test-chenxianhao/user-management/common"
	"backend-test-chenxianhao/user-management/user/http/routers"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

func Init() {
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	common.Init()

}

func main() {
	Init()
	port := viper.GetString("server.port")
	fmt.Println("Server Running on Port: ", viper.GetString("port"))
	http.ListenAndServe(":"+port, RouteInit())
}

// init all routes of all features
func RouteInit() *gin.Engine {
	systemRouters := gin.New()
	user := systemRouters.Group("/v1/users")
	routers.InitUserRouters(user)
	return systemRouters
}
