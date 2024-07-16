package sharedentities

type Address struct {
	AddressLine string `json:"address_line"bson:"address_line"	validate: "required"`
	City        string `json:"city"bson:"city"	validate: "required"`
	Country     string `json:"country"bson:"country"	validate: "required"`
	CityCode    int    `json:"city_code"bson:"city_code"	validate: "required"`
}
