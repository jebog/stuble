package requests

type UserDetailsInput struct {
	UserId          uint   `binding:"required"`
	FirstName       string `binding:"required"`
	LastName        string `binding:"required"`
	PhoneNumber     string
	Description     string
	ProfileImage    string
	Email           string `binding:"required"`
	EmailVerifiedAt string
	RememberToken   string
}
