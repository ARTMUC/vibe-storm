package dto

// GetUsersQuery represents query parameters for getting users
type GetUsersQuery struct {
	Page      int    `query:"page" validate:"min=1" example:"1" default:"1"`
	PerPage   int    `query:"per_page" validate:"min=1,max=100" example:"10" default:"10"`
	Search    string `query:"search" validate:"max=100" example:"john"`
	SortBy    string `query:"sort_by" validate:"omitempty,oneof=id email username created_at" example:"created_at" default:"created_at"`
	SortOrder string `query:"sort_order" validate:"omitempty,oneof=asc desc" example:"desc" default:"desc"`
	IsActive  *bool  `query:"is_active" example:"true"`
}

// GetUsersResponse represents the response for getting users with pagination
type GetUsersResponse struct {
	Users      []UserResponse `json:"users"`
	Pagination PaginationMeta `json:"pagination"`
}

// CreateUserResponse represents the response for creating a user
type CreateUserResponse struct {
	User    UserResponse `json:"user"`
	Message string       `json:"message" example:"User created successfully"`
}

// UpdateUserResponse represents the response for updating a user
type UpdateUserResponse struct {
	User    UserResponse `json:"user"`
	Message string       `json:"message" example:"User updated successfully"`
}

// DeleteUserResponse represents the response for deleting a user
type DeleteUserResponse struct {
	Message string `json:"message" example:"User deleted successfully"`
}

// GetUserByIDResponse represents the response for getting a user by ID
type GetUserByIDResponse struct {
	User UserResponse `json:"user"`
}

// PaginationParams represents common pagination parameters
type PaginationParams struct {
	Page    int `json:"page" query:"page" validate:"min=1" example:"1"`
	PerPage int `json:"per_page" query:"per_page" validate:"min=1,max=100" example:"10"`
}

// SetDefaults sets default values for pagination parameters
func (p *PaginationParams) SetDefaults() {
	if p.Page <= 0 {
		p.Page = 1
	}
	if p.PerPage <= 0 {
		p.PerPage = 10
	}
	if p.PerPage > 100 {
		p.PerPage = 100
	}
}

// GetOffset calculates the offset for database queries
func (p *PaginationParams) GetOffset() int {
	return (p.Page - 1) * p.PerPage
}

// GetLimit returns the limit for database queries
func (p *PaginationParams) GetLimit() int {
	return p.PerPage
}

// ToPaginationMeta converts pagination params to metadata
func (p *PaginationParams) ToPaginationMeta(totalCount int64) PaginationMeta {
	totalPages := int((totalCount + int64(p.PerPage) - 1) / int64(p.PerPage))
	if totalPages == 0 {
		totalPages = 1
	}

	return PaginationMeta{
		Page:       p.Page,
		PerPage:    p.PerPage,
		TotalCount: totalCount,
		TotalPages: totalPages,
	}
}
