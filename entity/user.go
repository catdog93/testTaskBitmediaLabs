package entity

type User struct {
	ID       uint64 `json:"id" binding:"required" bson:"_id"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserBody struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

//type User struct {
//	ID        primitive.ObjectID `json:"id" binding:"required" bson:"_id"`
//	*Author   `json:"author" binding:"required" bson:"author"`
//	*User `json:"user_auth" binding:"required" bson:"user_auth"`
//	Gender    Gender `json:"gender,omitempty" bson:"gender"`
//	Country   string `json:"country,omitempty" bson:"country"`
//	BirthDate string `json:"birth_date,omitempty" bson:"birthDate"`
//}
//
//type UserBody struct {
//	*AuthorBody `json:"author" binding:"required" bson:"author"`
//	Gender      Gender `json:"gender,omitempty" bson:"gender"`
//	Country     string `json:"country,omitempty" bson:"country"`
//	BirthDate   string `json:"birth_date,omitempty" bson:"birthDate"`
//}
//
//type Gender string
//
//const (
//	// Preferable values for field Gender
//	Male   Gender = "Male"
//	Female Gender = "Female"
//
//	MaleLower   Gender = "male"
//	FemaleLower Gender = "female"
//)
//
//func (userBody UserBody) ConvertUserBodyToUser() User {
//	user := User{
//		Gender:    userBody.Gender,
//		Country:   userBody.Country,
//		BirthDate: userBody.BirthDate,
//	}
//	user.Author.Nickname = userBody.AuthorBody.Nickname
//	return user
//}
