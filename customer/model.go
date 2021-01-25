package customer

import "go.mongodb.org/mongo-driver/bson/primitive"

//Customer model
type Customer struct {
	Id        primitive.ObjectID `bson:"_id" json:"id"`
	Email     string             `bson:"email" json:"email"`
	Password  string             `bson:"password" json:"password"`
	FirstName string             `bson:"firstname" json:"firstname"`
	Lastname  string             `bson:"lastname" json:"lastname"`
	Age       int                `bson:"age" json:"age"`
}

//CustomerRegistorBody model
type CustomerRegistorBody struct {
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required"`
	FirstName string `json:"firstname" validate:"required"`
	Lastname  string `json:"lastname" validate:"required"`
	Age       int    `json:"age" validate:"gte=1,lte=80"`
}

type CustomerResponseBody struct {
	Email     string `json:"email"`
	FirstName string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Age       int    `json:"age"`
}

//CustomerLoginBody model
type CustomerLoginBody struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

//CustomerChangePassword model
type CustomerChangePassword struct {
	OldPassword string `json:"oldPassword" validate:"required"`
	NewPassword string `json:"NewPassword" validate:"required"`
}
