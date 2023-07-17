package model

type RoleDto struct {
	Id string `json:"id"`

	Name string `json:"name"`
}

type ImportUserDto struct {
	File string `json:"file"`

	RoleName string `json:"roleName"`

	CompanyId int `json:"companyId"`
}

type UserCreateDto struct {
	Username string `json:"username"`

	Email string `json:"email"`

	RoleName string `json:"roleName"`

	CompanyId int `json:"companyId"`
}

type UserUpdateDto struct {
	Email string `json:"email"`

	RoleName string `json:"roleName"`

	CompanyId int `json:"companyId"`
}

type UserDto struct {
	Id string `json:"id"`

	Username string `json:"username"`

	Enabled bool `json:"enabled"`

	FirstName string `json:"firstName"`

	LastName string `json:"lastName"`

	Email string `json:"email"`

	CreatedTimestamp string `json:"createdTimestamp"`
}
