package response

type LoginResponse struct {
	Token  string `json:"token,omitempty"`
	Status bool   `json:"bool,omitempty"`
}

type AccountResponse struct {
	Token  string `json:"token,omitempty"`
	Status bool   `json:"bool,omitempty"`
}
