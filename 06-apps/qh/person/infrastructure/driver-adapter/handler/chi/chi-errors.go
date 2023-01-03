
import (
	"fmt"
)

type FuryHandlerError struct {
	StatusCode int
	Message    string
	Method     string
}

func (s FuryHandlerError) Error() string {
	return fmt.Sprintf("error at method %v - error message: %v", s.Method, s.Message)
}
