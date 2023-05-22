package utils

import (
	"github.com/sixfwa/fiber-api/dto"
	"github.com/sixfwa/fiber-api/models"
)

func CreateResponseUser(user models.User) dto.User {
	return dto.User{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
	}
}

func CreateResponseOrder(order models.Order, user dto.User, product dto.Product) dto.Order {
	return dto.Order{
		ID:      order.ID,
		User:    user,
		Product: product,
	}
}

func CreateResponseProduct(product models.Product) dto.Product {
	return dto.Product{
		ID:   product.ID,
		Name: product.Name,
		SKU:  product.SKU,
	}
}
