package puck

type Service struct {
	name string
}

func (s *Service) Name() string {
	return s.name
}
