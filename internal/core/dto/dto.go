package dto

type UserAccountDTO struct {
	Username    string
	Password    string
	CreatedDate string
	UpdatedDate string
	Type        uint
}

type StudentDTO struct {
	Mahv         string
	Surname      string
	Name         string
	Gender       string
	DayOfBirth   string
	PlaceOfBirth string
}
