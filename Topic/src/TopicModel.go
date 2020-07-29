package src

type Topic struct {
	TopicId int `json:"id"`
	TopicTitle string `json:"title"  form:"title" binding:"required"`
}

type Topics struct {
	TopicList []Topic `json:"Topics" binding:"gt=1,lt=3,dive,required"`
	TopicListSize int `json:"size"`
}


func CreateTopic(id int, title string) Topic {
	return Topic{id, title}
}

type TopicQuery struct {
	UserName string `json:"username" form:"username" binding:"required"`
	Page int `json:"page"`
	PageSize int `json:"page_size"`
}