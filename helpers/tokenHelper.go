package helpers

import (
	"TaskJWT/database"
	"context"
	"encoding/base64"
	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
	"log"
	"time"
)

var jwtKey = []byte("my_secret_key")

type SignedDetails struct {
	Guid string
	jwt.StandardClaims
}

var UserCollection *mongo.Collection = database.OpenCollection(database.Client, "user")

func GenerateAllToken(guid string) (token, refreshTokenToString string, err error) {
	claims := &SignedDetails{
		Guid: guid,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 60).Unix(),
		},
	}

	token, err = jwt.NewWithClaims(jwt.SigningMethodHS512, claims).SignedString(jwtKey)
	refreshToken, err := bcrypt.GenerateFromPassword(jwtKey, 14)
	if err != nil {
		log.Panic(err)
	}

	refreshTokenToString = base64.StdEncoding.EncodeToString(refreshToken)

	return token, refreshTokenToString, err
}

func CheckRefreshToken(refreshToken, jwtKey string) bool {
	decodeString, _ := base64.StdEncoding.DecodeString(refreshToken)
	err := bcrypt.CompareHashAndPassword(decodeString, []byte(jwtKey))
	return err == nil
}

// обновляет рефрешь токен в бд
func Update(guid string) {
	_, refreshToken, err := GenerateAllToken(guid)
	if err != nil {
		log.Fatal(err)
	}

	opts := options.Update().SetUpsert(true)
	filter := bson.M{"guid": guid}
	update := bson.D{{"$set", bson.D{{"refreshToken", refreshToken}}}}

	_, err = UserCollection.UpdateOne(context.TODO(), filter, update, opts)
	if err != nil {
		log.Fatal(err)
	}

	return
}
