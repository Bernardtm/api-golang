package status

type StatusService interface {
	GetAll() ([]StatusResponse, error)
	GetByID(id string) (StatusResponse, error)
	Create(entity StatusRequest) (string, error)
	Update(id string, entity StatusRequest) error
	Delete(id string) error
	Paginate(page int, size int) ([]StatusResponse, error)
}

type statusService struct {
	repo StatusRepository
}

func NewStatusService(repo StatusRepository) *statusService {
	return &statusService{repo: repo}
}

func (s *statusService) GetAll() ([]StatusResponse, error) {
	return s.repo.GetAll()
}

func (s *statusService) GetByID(id string) (StatusResponse, error) {
	return s.repo.GetByID(id)
}

func (s *statusService) Create(entity StatusRequest) (string, error) {

	return s.repo.Create(entity)
}

func (s *statusService) Update(id string, entity StatusRequest) error {
	return s.repo.Update(id, entity)
}

func (s *statusService) Delete(id string) error {
	return s.repo.Delete(id)
}

func (s *statusService) Paginate(page int, size int) ([]StatusResponse, error) {
	return s.repo.Paginate(page, size)
}
