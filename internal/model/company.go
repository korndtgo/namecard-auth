package model

type CompanyCreateDto struct {
	DefaultLanguage string `json:"defaultLanguage"`

	Logo string `json:"logo"`

	NameTh string `json:"nameTh"`

	NameEn string `json:"nameEn"`

	AddressTh string `json:"addressTh"`

	AddressEn string `json:"addressEn"`

	ContactNumber1 string `json:"contactNumber1"`

	ContactNumber2 string `json:"contactNumber2"`

	BackgroundColor string `json:"backgroundColor"`

	FontColor string `json:"fontColor"`

	ButtonColor string `json:"buttonColor"`

	FontOnButtonColor string `json:"fontOnButtonColor"`
}

type CompanyUpdateDto struct {
	DefaultLanguage string `json:"defaultLanguage"`

	Logo string `json:"logo"`

	NameTh string `json:"nameTh"`

	NameEn string `json:"nameEn"`

	AddressTh string `json:"addressTh"`

	AddressEn string `json:"addressEn"`

	ContactNumber1 string `json:"contactNumber1"`

	ContactNumber2 string `json:"contactNumber2"`

	BackgroundColor string `json:"backgroundColor"`

	FontColor string `json:"fontColor"`

	ButtonColor string `json:"buttonColor"`

	FontOnButtonColor string `json:"fontOnButtonColor"`
}

type QueryCompany struct {
	Id       string `json:"id"`
	Page     string `json:"page"`
	PageSize string `json:"pageSize"`
	Search   string `json:"search"`
	SortBy   string `json:"sortBy"`
	OrderBy  string `json:"orderBy"`
}
