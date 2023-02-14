package account

type Adress struct {
	country string `json:"country"`
}

type Account struct {
	ID          string `json:"id"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	AvatarUrl   string `json:"avatarUrl"`
	DOB         string `json:"dob"`
	Password    string `json:"password"`
	Language    string `json:"language"`
	Adress      Adress `json:"adress"`
	PhoneNumber string `json:"phoneNumber"`
}
