package Funs

import (
	"elk/AppInit"
	"github.com/gin-gonic/gin"
)

func Loadbook( ctx *gin.Context){
	rsp,err :=AppInit.Esclient().Search().Index("books").Do(ctx)
	if err !=nil{
		ctx.JSON(500,gin.H{"error":err})
	}else{
			ctx.JSON(200,gin.H{
				"error":rsp.Hits.Hits,

		})
	}

}