package demo

type Showcase interface {
	Show() (string, error)
}

type MyShowcase struct {
	cfg *Config
}

func NewMyShowcase(cfg *Config) *MyShowcase {
	return &MyShowcase{cfg: cfg}
}

func (o *MyShowcase) Show() (string, error) {
	return "You started your Go project from scratch 🌱. Well done! 🎉", nil
}
