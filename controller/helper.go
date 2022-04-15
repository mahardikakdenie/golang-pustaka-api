package controller

import (
	"encoding/json"
	"pustaka-api/book"
	"pustaka-api/entity"

	"golang.org/x/crypto/bcrypt"
)

func responses(books entity.Book) book.BookResponse {
	var bookResponseWithUser book.BookResponse

	if books.UserId != 0 {
		bookResponseWithUser = book.BookResponse{
			Id:          int(books.ID),
			Title:       books.Title,
			Price:       books.Price,
			Rating:      books.Rating,
			Substitle:   books.Substitle,
			Description: books.Description,
			UserId:      int(books.UserId),
			User:        books.User,
			CreatedAt:   books.CreatedAt,
			UpdatedAt:   books.UpdatedAt,
		}
	} else {
		bookResponseWithUser = book.BookResponse{
			Id:          int(books.ID),
			Title:       books.Title,
			Price:       books.Price,
			Rating:      books.Rating,
			Substitle:   books.Substitle,
			Description: books.Description,
			UserId:      int(books.UserId),
			CreatedAt:   books.CreatedAt,
			UpdatedAt:   books.UpdatedAt,
		}
	}

	return bookResponseWithUser

}

func ternaryTitle(title string, indexBook entity.Book) string {
	if title != "" {
		indexBook.Title = title
	} else {
		title = indexBook.Title
		indexBook.Title = title
	}

	return indexBook.Title
}

func ternaryPrice(price json.Number, indexBook entity.Book) json.Number {
	if price != "" {
		indexBook.Price = price
	} else {
		price = indexBook.Price
		indexBook.Price = price
	}
	return indexBook.Price
}

func ternarySubstitle(substitle string, indexBook entity.Book) string {
	if substitle != "" {
		indexBook.Substitle = substitle
	} else {
		substitle = indexBook.Substitle
		indexBook.Substitle = substitle
	}

	return indexBook.Substitle
}

func ternaryRating(rating json.Number, indexBook entity.Book) json.Number {
	if rating != "" {
		indexBook.Rating = rating
	} else {
		rating = indexBook.Rating
		indexBook.Rating = rating
	}
	return indexBook.Rating
}

func ternaryDescription(description string, indexBook entity.Book) string {
	if description != "" {
		indexBook.Description = description
	} else {
		description = indexBook.Description
		indexBook.Description = description
	}
	return indexBook.Description
}

func ternaryUserId(userId int, indexBook entity.Book) int {

	if userId != 0 {
		indexBook.UserId = userId
	} else {
		userId = indexBook.UserId
		indexBook.UserId = userId
	}
	return indexBook.UserId
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)

	return string(bytes), err
}

func CheckPasswordHash(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func checkerName(form entity.User, name string) string {
	if name != "" {
		form.Name = name
	}

	if name == "" {
		name = form.Name
		form.Name = name
	}

	return form.Name
}

func checkerEmail(form entity.User, email string) string {
	if email != "" {
		form.Email = email
	}

	if email == "" {
		email = form.Email
		form.Email = email
	}

	return form.Email
}

func checkerPassword(form entity.User, password string) string {
	if password != "" {
		form.Password = password
	}

	if password == "" {
		password = form.Password
		form.Password = password
	}

	return form.Password
}
