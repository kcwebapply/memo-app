package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	. "github.com/kcwebapply/memo-app/api/resource"
	service "github.com/kcwebapply/memo-app/domain/service"
)

var request_path_root = "/memo/"

func SetRouter(r *gin.Engine) *gin.Engine {
	r.GET(request_path_root+"list", getList)
	r.GET(request_path_root+"get/:memo_id", getMemo)
	r.POST(request_path_root+"post", postMemo)
	r.DELETE(request_path_root+"delete/:memo_id", deleteMemo)
	return r
}

func getList(c *gin.Context) {
	memos := service.GetAllMemo()
	c.JSON(http.StatusOK, memos)
}

func getMemo(c *gin.Context) {
	memo := service.GetMemo(c.Param("memo_id"))
	c.JSON(http.StatusOK, memo)
}

func postMemo(c *gin.Context) {
	memorequest := MemoRequest{}
	c.BindJSON(&memorequest)
	service.PostMemo(memorequest)
	c.JSON(http.StatusOK, memorequest)
}

func deleteMemo(c *gin.Context) {
	service.DeleteMemo(c.Param("memo_id"))
}
