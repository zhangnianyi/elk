package main

import (
	"context"
	"elk/AppInit"
	"elk/Models"
	"fmt"
	"log"
	"strconv"
)

func main() {
	pasge :=1
	pagezise :=50
	for{
		//读取50条到booklist里面
		book_list :=Models.BookList{}
		db :=AppInit.GetDB().Order("book_id desc").Limit(pagezise).Offset((pasge-1)*pagezise).Find(&book_list)
		if db.Error !=nil  ||len(book_list) ==0{
			break
		}
		//fmt.Println(book_list[2])
		//遍历book list插入到es
		for _,book :=range book_list{
			ctx :=context.Background()
			rsp,err :=AppInit.Esclient().Index().Index("books").
				Id(strconv.Itoa(book.BookID)).BodyJson(book).Do(ctx)
			if err !=nil{
				log.Fatal(err)
			}
			fmt.Println(rsp)
		}

		break
	}
}
