package controller

import (
	"TaskJWT/database"
	"TaskJWT/helpers"
	"TaskJWT/models"
	"context"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"net/http"
	"time"
)

var userCollection *mongo.Collection = database.OpenCollection(database.Client, "user")
var jwtKey = []byte("my_secret_key")

func Login(w http.ResponseWriter, r *http.Request) {
	guid := r.URL.Query().Get("guid")
	token, refreshToken, err := helpers.GenerateAllToken(guid)
	if err != nil {
		http.Error(w, "Server error", 500)
		return
	}

	users := bson.D{{"guid", guid}, {"refreshToken", refreshToken}}
	result, err := userCollection.InsertOne(context.TODO(), users)
	if err != nil {
		fmt.Errorf("USER %v", err)
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   token,
		Expires: time.Now().Add(60 * time.Minute),
	})

	fmt.Fprintf(w, "Token %s", token)
	fmt.Fprintf(w, "\nRefreshToken %s", refreshToken)
	fmt.Fprintf(w, "\nInsertID нового пользователя %s", result.InsertedID)

}


type SignedDetails struct {
	Guid string
	jwt.StandardClaims
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	guid := r.URL.Query().Get("guid")
	guid1 := r.URL.Query().Get("refreshToken")

	user, err := getUser(guid)
	if err != nil {
		log.Fatal(err)
	}

	token, _, _ := helpers.GenerateAllToken(guid)

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   token,
		Expires: time.Now().Add(60 * time.Minute),
	})

	checkValidRefreshToken := helpers.CheckRefreshToken(guid1, string(jwtKey))

	if checkValidRefreshToken == true {
		helpers.Update(guid)
	} else {
		http.Error(w, "API Error: No authorization token was found", 401)
	}

	fmt.Fprintf(w, "\nCheckValidToken %v", checkValidRefreshToken)
	fmt.Fprintf(w, "\nUserID %v", user.ID)

}

func getUser(guid string) (*models.User, error) {
	var user *models.User

	err := userCollection.FindOne(context.TODO(), bson.M{"guid": guid}).Decode(&user)
	if err != nil {
		log.Fatal(err)
	}

	return user, err
}
