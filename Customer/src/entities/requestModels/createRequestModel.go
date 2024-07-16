package upRequestModel

import sharedentities "CustomerApi/shared/entities"

type RequestModel struct {
	FirstName string                  `json:"first_name"bson:"first_name"`
	LastName  string                  `json:"last_name"bson:"last_name"`
	Email     string                  `json:"email"bson:"email"`
	Phone     string                  `json:"phone"bson:"phone"`
	Address   *sharedentities.Address `json:"address"bson:"address"`
}
