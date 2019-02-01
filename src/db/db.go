package db

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
	"my-blog-server/src/utils"
	"net/http"
)

const DB_ADDR = "127.0.0.1:27017"

func GetCollect(databaseName, collectName string, args ...interface{}) (*mgo.Collection, *mgo.Session) {

	var dbAddr string = DB_ADDR
	if len(args) > 0 {
		dbAddr = args[0].(string)
	}
	session, err := mgo.Dial(dbAddr)


	utils.HandleError("connect mongo error: ", err)

	c := session.DB(databaseName).C(collectName)

	return c, session
}


func GetUserInfo(router *gin.Engine) {

	router.GET("/mongo", func(context *gin.Context) {
		 user := FindByName("mao")
		 b, _ := json.Marshal(user)
		 context.JSON(http.StatusOK, string(b))
	})

}

func main() {
	//InitUserData()
	//fmt.Println("id is: ", FindByPhone("2").Id)
	fmt.Println(FindByName("mao"))
}
