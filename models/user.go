package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID           primitive.ObjectID `bson:"_id"`
	Guid         string             `json:"guid"`
	RefreshToken string             `json:"refresh_token"`
}

//type Tokens struct {
//	Token        jwt.StandardClaims
//	RefreshToken []byte
//}
