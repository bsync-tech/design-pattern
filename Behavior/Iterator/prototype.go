package prototype

import "sync"

type SingletonDuck struct {
	name string
}

func (s *SingletonDuck) GetName() string {
	return s.name
}

var instance *SingletonDuck

var once sync.Once

func GetInstance() *SingletonDuck {
	once.Do(func() {
		instance = &SingletonDuck{name: "small yellow duck"}
	})
	return instance
}
