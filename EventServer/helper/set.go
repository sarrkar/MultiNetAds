package helper

type Set struct {
	map_ map[string]bool
}

func NewSet() Set {
	return Set{
		map_: map[string]bool{},
	}
}

func (s *Set) Add(item string) {
	s.map_[item] = true
}

func (s *Set) Check(item string) bool {
	_, ok := s.map_[item]
	return ok
}
