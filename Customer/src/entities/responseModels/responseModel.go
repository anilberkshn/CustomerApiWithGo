package responseModels

type CustomerResponseModel struct {
	Id        string `json:"id" bson:"_id"`
	FirstName string `json:"first_name,omitempty"bson:"first_name"`
	LastName  string `json:"last_name,omitempty"bson:"last_name"`
	Email     string `json:"email,omitempty"bson:"email"`
	CreatedAt string `json:"created_at,omitempty"bson:"created_at"`
	UpdatedAt string `json:"updated_at,omitempty"bson:"updated_at"`
}
