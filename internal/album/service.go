package album

var _ ServiceI = (*Service)(nil)

type ServiceI interface {
	CreateAlbum(*Album) error
	GetAlbums() ([]Album, error)
	GetAlbumByID(id int64) (*Album, error)
}
type Service struct {
	repo RepositoryI
}

func NewService(repo RepositoryI) *Service {
	return &Service{repo: repo}
}

func (s *Service) CreateAlbum(album *Album) error {
	return s.repo.Create(album)
}

func (s *Service) GetAlbums() ([]Album, error) {
	return s.repo.GetAll()
}

func (s *Service) GetAlbumByID(id int64) (*Album, error) {
	return s.repo.GetByID(id)
}
