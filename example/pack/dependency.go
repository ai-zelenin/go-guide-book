package factory

type Dependency struct {
}

func (d *Dependency) GetData() (string, error) {
	return "some data from Dependency", nil
}
