package api

type Hello struct {
	Language string `json:"language"`
	Message string	`json:"message"`
}

type HelloApi struct {
	HelloAll []Hello `json:"helloAll"`
}
