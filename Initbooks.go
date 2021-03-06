package main

import (
	"context"
	"elk/AppInit"
	"github.com/olivere/elastic/v7"
	"elk/Models"
	"fmt"
	"log"
	"strconv"
	"sync"
)

func main() {
	page :=1
	pagesize :=500
	wg :=sync.WaitGroup{}
	for{
		book_list:=Models.BookList{}
		//读取50条到booklist里面
		db:=AppInit.GetDB().Select("book_id,book_name,book_intr,book_price1,book_price2,book_author,book_press,book_kind " +
			",if(book_date='','1970-01-01',ltrim(book_date)) as book_date").
			Order("book_id desc").Limit(pagesize).Offset((page-1)*pagesize).Find(&book_list)
		if db.Error!=nil || len(book_list)==0{
			break
		}
		wg.Add(1)   //开启携程
		//client :=AppInit.Esclient()
		//bulk :=client.Bulk()
		////fmt.Println(book_list[2])
		////遍历book list插入到es
		//for _,books :=range book_list{
		//	rqe :=elastic.NewBulkIndexRequest()
		//	rqe.Index("books").Id(strconv.Itoa(books.BookID).Doc(books))
		//	bulk.Add(rqe)
		//	//ctx :=context.Background()
		//	//rsp,err :=AppInit.Esclient().Index().Index("books").
		//	//	Id(strconv.Itoa(book.BookID)).BodyJson(book).Do(ctx)
		//	//if err !=nil{
		//	//	log.Fatal(err)
		//	//}
		//	//fmt.Println(rsp)
		//}
		//rsp,err :=bulk.Do(context.Background())
		go func() {
			defer  wg.Done()  //很重要
			client:=AppInit.Esclient()   //打开一个es的链接
			bulk:=client.Bulk()  //采用bulk的方式插入
			for _,book:=range book_list {

				req:=elastic.NewBulkIndexRequest().Index("books").
					Id(strconv.Itoa(book.BookID)).Doc(book)
				bulk.Add(req)

			}
			rsp,err:=bulk.Do(context.Background())  //执行这个bulk
			if err !=nil{
				log.Fatal(err)
			}else {
				fmt.Println(rsp)
			}
		}()
		page=page+1   //这个必须加 不要跳循环了跳转不到第二页
	}
		wg.Wait()
}
