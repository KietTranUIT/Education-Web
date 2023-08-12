package request

type SignUpRequest struct {
	Username     string `Json:"username"`
	Password     string `Json:"password"`
	Surname      string `Json:"surname"`
	Name         string `Json:"name"`
	DateOfBirth  string `Json:"dateofbirth"`
	Gender       string `Json:"gender"`
	PlaceOfBirth string `Json:"placeofbirth"`
}
