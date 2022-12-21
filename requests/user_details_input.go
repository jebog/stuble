package requests

type UserDetailsInput struct {
	UserId          string `json:"user_id" binding:"required"`
	FirstName       string `json:"first_name" binding:"required"`
	LastName        string `json:"last_name" binding:"required"`
	PhoneNumber     string `json:"phone_number" binding:"required"`
	Description     string `json:"description" binding:"required"`
	ProfileImage    string `json:"profile_image" binding:"required"`
	Email           string `json:"email" binding:"required"`
	EmailVerifiedAt string `json:"email_verified_at" binding:"required"`
	RememberToken   string `json:"remember_token" binding:"required"`
}
