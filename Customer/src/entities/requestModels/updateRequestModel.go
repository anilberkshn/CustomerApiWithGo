package upRequestModel

import (
	sharedentities "CustomerApi/shared/entities"
)

type UpdateRequestModel struct {
	FirstName string                  `json:"first_name"bson:"first_name" validate:"required"`
	LastName  string                  `json:"last_name" bson:"last_name" validate:"required"`
	Email     string                  `json:"email" bson:"email" validate:"required"`
	Phone     string                  `json:"phone" bson:"phone" validate:"required"`
	Address   *sharedentities.Address `json:"address" bson:"address" validate:"required"`
}
