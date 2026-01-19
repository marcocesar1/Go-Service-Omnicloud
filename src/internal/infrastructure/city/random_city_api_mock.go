package city

type RandomCityApiMock struct {
	City  string
	Error error
}

func NewRandomCityApiMock() *RandomCityApiMock {
	return &RandomCityApiMock{
		City: "New York",
	}
}

func (r *RandomCityApiMock) GetCityName() (string, error) {
	if r.Error != nil {
		return "", r.Error
	}

	return r.City, nil
}
