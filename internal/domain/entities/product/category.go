package product

import "errors"

type Category struct {
	Name string
}

func NewCategory(name string) (*Category, error) {
	category := &Category{Name: name}

	err := category.Validate()
	if err != nil {
		return nil, err
	}

	return category, nil
}

func (c *Category) Validate() error {
	if c.Name == "" {
		return errors.New("Category name cannot be empty")
	}
	return nil
}