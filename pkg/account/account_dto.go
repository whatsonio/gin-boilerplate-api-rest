package account

type AccountCreateInput struct {
	Name string `json:"name" binding:"required"`
}

type AccountUpdateInput struct {
	Name string `json:"name" binding:"required"`
}
