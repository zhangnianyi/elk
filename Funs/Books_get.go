package Funs

import (
	"elk/AppInit"
	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic/v7"
	"strings"
)

func Loadbook( ctx *gin.Context){
	//go操作es 查看索引的内容
	rsp,err :=AppInit.Esclient().Search().Index("books").Do(ctx)
	if err !=nil{
		ctx.JSON(500,gin.H{"error":err})
	}else{
			ctx.JSON(200,gin.H{
				"error":rsp.Hits.Hits,

		})
	}

}
func Loadbookbypress( ctx *gin.Context){
	//go 操作es进行精准匹配
	press,_ :=ctx.Params.Get("press")
	list :=strings.Split(press,",")
	preList :=[]interface{}{}
	for _,p:=range list{
		preList = append(preList, p)
	}
	//这句话是重点
	//es查询两个同时匹配的对象 
	newquert:=elastic.NewTermsQuery("BookPress",preList...)
	rsp,err :=AppInit.Esclient().Search().Query(newquert).Index("books").Do(ctx)
	if err !=nil{
		ctx.JSON(500,gin.H{"erro r":err})
	}else{
		ctx.JSON(200,gin.H{
			"error":rsp.Hits.Hits,

		})
	}

}