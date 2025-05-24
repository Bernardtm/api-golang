package menus

type MenusService interface {
	GetAll() ([]MenusResponse, error)
	GetByID(id string) (MenusResponse, error)
	Create(entity MenusRequest) (string, error)
	Update(id string, entity MenusRequest) error
	Delete(id string) error
	Paginate(page int, size int) ([]MenusResponse, error)
	GetMenusByUserID(userID string) ([]MenusResponse, error)
}

type menusService struct {
	repo                   MenusRepository
}

func NewMenusService(
	repo MenusRepository,
) *menusService {
	return &menusService{
		repo:                   repo,
	}
}

func (s *menusService) GetAll() ([]MenusResponse, error) {
	return s.repo.GetAll()
}

func (s *menusService) GetByID(id string) (MenusResponse, error) {
	return s.repo.GetByID(id)
}

func (s *menusService) Create(entity MenusRequest) (string, error) {
	return s.repo.Create(entity)
}

func (s *menusService) Update(id string, entity MenusRequest) error {
	return s.repo.Update(id, entity)
}

func (s *menusService) Delete(id string) error {
	return s.repo.Delete(id)
}

func (s *menusService) Paginate(page int, size int) ([]MenusResponse, error) {
	return s.repo.Paginate(page, size)
}

func (s *menusService) GetMenusByUserID(userID string) ([]MenusResponse, error) {
	menus, err := s.repo.GetAll()
	if err != nil {
		return []MenusResponse{}, err
	}
	return menus, nil
}
