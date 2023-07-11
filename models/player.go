package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Player struct {
	Id		 primitive.ObjectID  `json:"id"`
	Name     string 			 `json:"name"`
	Gols     string 			 `json:"gols"`
	Curso    string 			 `json:"curso"`
	Linkedin string 			 `json:"linkedin"`
}
