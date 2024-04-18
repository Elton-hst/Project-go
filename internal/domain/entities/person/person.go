package person

import "errors"

type Person struct {
	Name     string
	Password string
	Email    string
	Phone    string
}

func NewPerson(name string, password string, email string, phone string) (*Person, error) {
	person := &Person{
		Name:     name,
		Password: password,
		Email:    email,
		Phone:    phone,
	}

	err := person.Validate()
	if err != nil {
		return nil, err
	}

	return person, nil
}

func (p *Person) Validate() error {

	if p.Name == "" || p.Email == "" || p.Password == "" {
		return errors.New("Person cannot be empty")
	}

	return nil
}
