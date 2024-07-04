package neuralservice

type Service struct{
	Id int `json:"id"`
	Name string `json:"name"`
	Cost string `json:"call_cost"`
}