package password

type RecommendMinimumPasswordActionRequest struct {
	InitPassword string `json:"init_password" validate:"required,max=40"`
}

type RecommendMinimumPasswordActionResponse struct {
	NumOfSteps uint `json:"num_of_steps"`
}
