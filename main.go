package main

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
	"log"
)

func main() {
	client,err  :=elastic.NewClient(
		elastic.SetURL("http://118.24.102.88:9200/"),
		elastic.SetSniff(false),
	)
	if err !=nil{
		log.Fatal(err)
	}
	//fmt.Println(client)
	ctx :=context.Background()
	//mapping,err:= client.GetMapping().Index("news").Do(ctx)
	//if err !=nil{
	//	log.Fatal(err)
	//}
	json := `{"news_title":"test1","news_type":"php","news_status":1}`

	data,err :=client.Index().Index("news").Id("101").
		BodyString(json).Do(ctx)
	if err !=nil{
		log.Fatal(err)
	}
	fmt.Println(data)

}

