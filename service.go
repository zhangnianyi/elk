package main

import (
	"elk/Funs"
	"github.com/gin-gonic/gin"
)
func main(){
	roter :=gin.Default()
	g:=roter.Group("/books")
	//{
	//	g.GET("/book", Funs.Loadbook)
	//}
	{
		g.Handle("GET","",Funs.Loadbook)
		g.Handle("GET","/press/:press",Funs.Loadbookbypress)

	}

	roter.Run()
}