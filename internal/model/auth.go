package model

type ChangePasswordDto struct {
	Password        string `json:"password"`
	NewPassword     string `json:"newPassword"`
	ConfirmPassword string `json:"confirmPassword"`
}

type ResetPasswordDto struct {
	RefCode         string `json:"refCode"`
	Otp             string `json:"otp"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirmPassword"`
}

type GetProfileDto struct {
	User    string `json:"user"`
	Company string `json:"company"`
}
