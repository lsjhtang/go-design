package classes

import (
	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic/v7"
	"wserver/goft"
	"wserver/initalize"
)

type Book struct {
	*goft.GormAdapter
}

func NewBook() *Book {
	return &Book{}
}

func MapToSilce(re *elastic.SearchResult, key string) []interface{} {
	results := make([]interface{}, 0)
	for _, hit := range re.Hits.Hits {
		results = append(results, hit.Fields[key].([]interface{})[0])
	}
	return results
}

func (this *Book) PressList(context *gin.Context) goft.Models {
	bp := elastic.NewCollapseBuilder("BookPress")
	cx, err := initalize.GetEsClient().Search().Collapse(bp).FetchSource(false).Index("books").Size(50).Do(context)
	goft.Error(err)
	return goft.MakeModels(MapToSilce(cx, "BookPress"))
}

func (this *Book) Build(goft *goft.Goft) {
	goft.Handle("GET", "/book_helper", this.PressList)
	goft.Handle("GET", "/book_list", this.BookList)
	goft.Handle("GET", "/book_query", this.BookQuery)
	goft.Handle("GET", "/book_search", this.BookSearch)
}
