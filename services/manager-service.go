package services

//go:generate counterfeiter -o fakes/fake_manager_service.go ManagerService
type ManagerService interface {
	CreateManager(opts *ManagerOptions) (Manager, error)
}

func NewManagerService() ManagerService {
	return &managerService{}
}

type managerService struct {}

func (s *managerService) CreateManager(opts *ManagerOptions) (Manager, error) {
	return NewManager(opts)
}
