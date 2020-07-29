package main

import (
	"fmt"
	/*"github.com/gin-gonic/gin"
	. "wserver/Topic/src"*/
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)


func main() {
	db, err := gorm.Open("mysql", "user:password@/dbname?charset=utf8&parseTime=True&loc=Local")
	defer db.Close()
	fmt.Print(err)
	/*r := gin.Default()
	v1 := r.Group("/v1/topic")
	v1.GET(":topic_id", GetTopicDetail)
	v1.Use(MustLogin())
	{
		v1.POST("", NewTopic)
		v1.DELETE(":topic_id", DelTopic)
	}


	r.GET("/ping/:abc", func(c *gin.Context) {
		c.JSON(200, gin.H{"hello": len(c.Params)})
	})
	r.Run()*/ // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

