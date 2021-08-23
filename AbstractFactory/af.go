package af

// In this pattern, Add Duck type need, you need
// 1. Add lines code in FactoryProduceDuck
// 2. Add new type Duck in this file

type Cook interface {
	CookWithToast() string
	CookWithBeer() string
}

type BeijingDuckFactory struct{}

func (b BeijingDuckFactory) CookWithToast() string {
	return "Cook with toast"
}

func (b BeijingDuckFactory) CookWithBeer() string {
	return "Cook with beer"
}

type ShaoxingDuckFactory struct{}

func (s ShaoxingDuckFactory) CookWithBeer() string {
	return "Cook with beer"
}

func (s ShaoxingDuckFactory) CookWithToastr() string {
	return "Cook method invalid"
}
