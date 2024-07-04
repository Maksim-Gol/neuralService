package neuralservice
import "time"
type ModelCall struct {
	UserID int `json:"user_id"`
	ModelID int `json:"model_id"`
	CallTime time.Time `json:"call_time"`
}