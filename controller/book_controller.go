package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"pustaka-api/book"
	"pustaka-api/middleware"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type bookController struct {
	bookService book.Service
}

func NewBookController(bookService book.Service) *bookController {
	return &bookController{bookService}
}

func (controller *bookController) Index(c *gin.Context) {
	entities := c.Query("entities")
	books, err := controller.bookService.FindAll(entities)

	if err != nil {
		var meta = gin.H{
			"status":  false,
			"message": err.Error(),
		}

		var data = gin.H{}
		c.JSON(http.StatusBadRequest, gin.H{
			"meta": meta,
			"data": data,
		})
		return
	}
	var meta = gin.H{
		"status":  true,
		"message": "Success",
	}
	var bookResponse []book.BookResponse
	for _, v := range books {
		bookResponses := responses(v)

		bookResponse = append(bookResponse, bookResponses)
	}
	c.JSON(http.StatusOK, gin.H{
		"meta": meta,
		"data": bookResponse,
	})
}

func (controller *bookController) Show(c *gin.Context) {
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	books, err := controller.bookService.FindById(id)
	fmt.Println("name =>", books.Title)
	if err != nil {
		var meta = gin.H{
			"status":  false,
			"message": err.Error(),
		}

		var data = gin.H{}
		c.JSON(http.StatusBadRequest, gin.H{
			"meta": meta,
			"data": data,
		})
		return
	}

	if books.ID == 0 {
		var meta = gin.H{
			"status":  false,
			"message": "Book not found",
		}

		c.JSON(http.StatusNotFound, gin.H{
			"meta": meta,
			"data": gin.H{},
		})

		return
	}
	var meta = gin.H{
		"status":  true,
		"message": "Success",
	}

	// var data entityResponses.Book

	// data = responses(books)

	c.JSON(http.StatusOK, gin.H{
		"meta": meta,
		"data": books,
	})
}

func (controller *bookController) PostBookHandler(c *gin.Context) {
	// title, price
	var BookInput book.BookRequest
	err := c.ShouldBindJSON(&BookInput)
	BookInput.UserId = middleware.UserId
	if err != nil {
		var errorMassages []string
		for _, e := range err.(validator.ValidationErrors) {
			errorMassage := e.Field() + " " + e.Tag() + " " + e.ActualTag()
			errorMassages = append(errorMassages, errorMassage)
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
			"e":     errorMassages,
		})
		return
	}

	book, err := controller.bookService.Create(BookInput)

	if err != nil {
		var errorMassages []string
		for _, e := range err.(validator.ValidationErrors) {
			errorMassage := e.Field() + " " + e.Tag() + " " + e.ActualTag()
			errorMassages = append(errorMassages, errorMassage)
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"status": false,
			"error":  err.Error(),
			"e":      errorMassages,
		})
		return
	}

	var meta = gin.H{
		"status":  true,
		"message": "Book created successfully",
	}

	c.JSON(http.StatusOK, gin.H{
		"meta": meta,
		"data": book,
	})
}

func (s *bookController) Update(c *gin.Context) {
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	bookId, err := s.bookService.FindById(id)

	if bookId.ID == 0 {
		var meta = gin.H{
			"status":  false,
			"message": "Book not found",
		}

		c.JSON(http.StatusNotFound, gin.H{
			"meta": meta,
			"data": gin.H{},
		})
		return
	}

	var BookInput book.BookRequest
	title := c.Request.FormValue("title")
	price := c.Request.FormValue("price")
	substitle := c.Request.FormValue("substitle")
	description := c.Request.FormValue("description")
	rating := c.Request.FormValue("rating")
	userId := c.Request.FormValue("user_id")
	user_id, _ := strconv.Atoi(userId)

	BookInput = book.BookRequest{
		Title:       ternaryTitle(title, bookId),
		Price:       ternaryPrice(json.Number(price), bookId),
		Substitle:   ternarySubstitle(substitle, bookId),
		Description: ternaryDescription(description, bookId),
		Rating:      ternaryRating(json.Number(rating), bookId),
		UserId:      ternaryUserId(user_id, bookId),
	}

	book, err := s.bookService.Update(BookInput, id)

	if err != nil {
		var errorMassages []string
		for _, e := range err.(validator.ValidationErrors) {
			errorMassage := e.Field() + " " + e.Tag() + " " + e.ActualTag()
			errorMassages = append(errorMassages, errorMassage)
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"status": false,
			"error":  err.Error(),
			"e":      errorMassages,
		})
		return
	}

	var meta = gin.H{
		"status":  true,
		"message": "Book updated successfully",
	}

	dataBook := responses(book)

	c.JSON(http.StatusOK, gin.H{
		"meta": meta,
		"data": dataBook,
	})
}

func (controller *bookController) Destroy(c *gin.Context) {
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	// indexBook, _ := controller.bookService.FindById(id)
	data, err := controller.bookService.Destroy(id)

	if data.ID == 0 {
		var meta = gin.H{
			"status":  false,
			"message": "Book not found",
		}

		c.JSON(http.StatusNotFound, gin.H{
			"meta": meta,
			"data": gin.H{},
		})

		return
	}

	if err != nil {
		var errorMassages []string
		for _, e := range err.(validator.ValidationErrors) {
			errorMassage := e.Field() + " " + e.Tag() + " " + e.ActualTag()
			errorMassages = append(errorMassages, errorMassage)
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"meta": gin.H{
				"status": false,
				"error":  err.Error(),
				"e":      errorMassages,
			},
			"data": gin.H{},
		})
		return
	}

	var meta = gin.H{
		"status":  true,
		"message": "Book deleted successfully",
		"code":    200,
	}

	c.JSON(http.StatusOK, gin.H{
		"meta": meta,
		"data": data,
	})
}
