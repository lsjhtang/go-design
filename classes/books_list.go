package classes

import (
	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic/v7"
	"reflect"
	"wserver/goft"
	"wserver/initalize"
	"wserver/models"
)

func ToMap(re *elastic.SearchResult) []*models.Books {
	booklist := models.BookList{}
	var book *models.Books
	for _, result := range re.Each(reflect.TypeOf(book)) {
		booklist = append(booklist, result.(*models.Books))
	}
	return booklist
}

func (this *Book) BookList(ctx *gin.Context) goft.Models {
	cx, err := initalize.GetEsClient().Search().Index("books").Do(ctx)
	goft.Error(err)
	return goft.MakeModels(ToMap(cx))
}

func (this *Book) BookQuery(ctx *gin.Context) goft.Models {
	param, _ := ctx.Params.Get("press")
	query := elastic.NewTermQuery("BookPress", param)
	cx, err := initalize.GetEsClient().Search().Query(query).Index("books").Do(ctx)
	goft.Error(err)
	return goft.MakeModels(ToMap(cx))
}
