package seeders

import (
	"learn-goravel/app/facades"
	"learn-goravel/app/models"
)

type UserSeeder struct {
}

func (s *UserSeeder) Signature() string {
	return "UserSeeder"
}

func (s *UserSeeder) Run() error {
	age1 := 25
	age2 := 30
	age3 := 28

	city1 := "New York"
	city2 := "Los Angeles"
	city3 := "Chicago"

	users := []models.Users{
		{
			Name:     "John Doe",
			Email:    "john@example.com",
			Password: "password123",
			Age:      age1,
			City:     city1,
			IsActive: true,
		},
		{
			Name:     "Jane Smith",
			Email:    "jane@example.com",
			Password: "password123",
			Age:      age2,
			City:     city2,
			IsActive: true,
		},
		{
			Name:     "Bob Wilson",
			Email:    "bob@example.com",
			Password: "password123",
			Age:      age3,
			City:     city3,
			IsActive: false,
		},
	}

	for _, user := range users {
		err := facades.Orm().Query().Create(&user)
		if err != nil {
			return err
		}
	}

	return nil
}
