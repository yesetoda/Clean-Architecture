package domain

type Task struct {
	Id          int `json:"id" bson:"id"`
	Title       string `json:"title" bson:"title"`
	Description string `json:"description" bson:"description"`
	Duedate     string `json:"duedate" bson:"duedate"`
	Status      string `json:"status" bson:"status"`
}
