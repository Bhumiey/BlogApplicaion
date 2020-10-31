package Code

type PostgresRepository interface {
	AddNewBlog(b Blog) error
	GetAllBlog() (BlogList, error)
	DeleteBlog(id int) error
}

type PostgressService interface {
	AddNewBlog(blog Blog) error
	GetAllBlog() (BlogList, error)
	DeleteBlog(id int) error
}

type postgressService struct {
	pr PostgresRepository
}

func PostGressNewService(pr PostgresRepository) PostgressService {
	return &postgressService{pr}
}

func (s *postgressService) AddNewBlog(blog Blog) error {
	if err := s.pr.AddNewBlog(blog); err != nil {
		return err
	}
	return nil
}

func (s *postgressService) GetAllBlog() (BlogList, error) {
	blogList, err := s.pr.GetAllBlog()
	if err != nil {
		return blogList, err
	}
	return blogList, nil
}
func (s *postgressService) DeleteBlog(id int) error {
	if err := s.pr.DeleteBlog(id); err != nil {
		return err
	}
	return nil
}
