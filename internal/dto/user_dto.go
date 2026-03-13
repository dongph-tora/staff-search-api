package dto

type UpdateUserRequest struct {
	Name  *string `json:"name"`
	Bio   *string `json:"bio"`
	Phone *string `json:"phone"`
}
