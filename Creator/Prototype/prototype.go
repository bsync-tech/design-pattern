package prototype

type Duck struct {
	name string
}

func (s Duck) String() string {
	return s.name
}

func (s *Duck) Clone() Duck {
	return *s
}
