package model

type CreateCardDto struct {
	EmployeeId     string `json:"employeeId"`
	NameTh         string `json:"nameTh"`
	NameEn         string `json:"nameEn"`
	SurnameTh      string `json:"surnameTh"`
	SurnameEn      string `json:"surnameEn"`
	NicknameTh     string `json:"nicknameTh"`
	NicknameEn     string `json:"nicknameEn"`
	Email          string `json:"email"`
	ContactNumber1 string `json:"contactNumber1"`
	ContactNumber2 string `json:"contactNumber2"`
	LineId         string `json:"lineId"`
	PositionTh     string `json:"positionTh"`
	PositionEn     string `json:"positionEn"`
	DepartmentTh   string `json:"departmentTh"`
	DepartmentEn   string `json:"departmentEn"`
}

type ImportCardDto struct {
	File string `json:"file"`
}

type QueryCard struct {
	Id         string `json:"id"`
	Page       string `json:"page"`
	PageSize   string `json:"pageSize"`
	EmployeeId string `json:"employeeId"`
	Search     string `json:"search"`
	SortBy     string `json:"sortBy"`
	OrderBy    string `json:"orderBy"`
}

type QueryImportCard struct {
	Token string `json:"token"`
}

type UpdateCardDto struct {
	NameTh         string `json:"nameTh"`
	NameEn         string `json:"nameEn"`
	SurnameTh      string `json:"surnameTh"`
	SurnameEn      string `json:"surnameEn"`
	NicknameTh     string `json:"nicknameTh"`
	NicknameEn     string `json:"nicknameEn"`
	Email          string `json:"email"`
	ContactNumber1 string `json:"contactNumber1"`
	ContactNumber2 string `json:"contactNumber2"`
	LineId         string `json:"lineId"`
	PositionTh     string `json:"positionTh"`
	PositionEn     string `json:"positionEn"`
	DepartmentTh   string `json:"departmentTh"`
	DepartmentEn   string `json:"departmentEn"`
}
