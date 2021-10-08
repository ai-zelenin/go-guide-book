package factory

type Dependency interface {
	LoadData() (string, error)
}

type dependency struct {
}

func NewDependency() Dependency {
	return &dependency{}
}

func (d *dependency) LoadData() (string, error) {
	return "some data from Dependency", nil
}
