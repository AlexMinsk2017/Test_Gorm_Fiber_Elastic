package orchestrator

type IUserOrchestrator interface {
	//GetByID(ctx context.Context, id uint) (*models.User, error)
	//Create(ctx context.Context, body *web.User) (*models.User, error)
	//DeleteMark(ctx context.Context, id uint) error
}
type UserOrchestrator struct {
	Engine *Engine
}

func NewUserOrchestrator(engine *Engine) IUserOrchestrator {
	return &UserOrchestrator{engine}
}
