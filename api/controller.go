package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	. "github.com/kcwebapply/examination/api/resource"
	service "github.com/kcwebapply/examination/domain/service"
)

var request_path_root = "/memo/"

func SetRouter(r *gin.Engine) *gin.Engine {
	//r := gin.Default()
	r.GET(request_path_root+"list", getList)
	r.GET(request_path_root+"property/:memo_id", getMemo)
	r.POST(request_path_root+"list", postMemo)
	//r.Run(":8888")
	return r
}

func getList(c *gin.Context) {
	fmt.Println("request", c)
}

func getMemo(c *gin.Context) {
	memo := service.GetMemo(c.Param("memo_id"))
	c.JSON(http.StatusOK, memo)
}

func postMemo(c *gin.Context) {
	memorequest := MemoRequest{}
	c.BindJSON(&memorequest)
	c.JSON(http.StatusOK, memorequest)
}
