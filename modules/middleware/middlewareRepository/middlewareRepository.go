package middlewareRepository

// Create interface and struct
type (
	MiddlewareRepositoryService interface{}

	middlewareRepository struct{}
)

// Create constructor
func NewMiddlewareRepository() MiddlewareRepositoryService {
	return &middlewareRepository{}
}
