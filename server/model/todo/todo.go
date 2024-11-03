package todo

import "go.mongodb.org/mongo-driver/bson/primitive"

type Model struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Task        string             `bson:"task" json:"task"`
	IsCompleted bool               `bson:"iscompleted" json:"iscompleted"`
}
