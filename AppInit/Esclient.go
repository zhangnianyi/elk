package AppInit

import (
	"github.com/olivere/elastic/v7"

)
func Esclient() *elastic.Client {
	 client, err := elastic.NewClient(
		elastic.SetURL("http://118.24.102.88:9200/"),
		elastic.SetSniff(false),
	)
	if err !=nil{
		return  nil
	}
	return  client
}
