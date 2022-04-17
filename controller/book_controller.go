package controller

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"pustaka-api/book"
	"pustaka-api/dot_env"
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

func (controller *bookController) FileUpload(ctx *gin.Context) {
	idString := ctx.Param("id")
	id, _ := strconv.Atoi(idString)
	file, header, err := ctx.Request.FormFile("file")
	if err != nil {
		var meta = gin.H{
			"status":  false,
			"message": "File not found",
		}
		ctx.JSON(http.StatusNotFound, gin.H{
			"meta": meta,
			"data": gin.H{},
		})
		return
	}

	fileName := header.Filename
	// make directory
	out, err := os.Create("public/" + fileName)

	if err != nil {
		var meta = gin.H{
			"status":  false,
			"message": "File not found Public",
		}
		ctx.JSON(http.StatusNotFound, gin.H{
			"meta": meta,
			"data": gin.H{},
		})
		return
	}

	defer out.Close()

	_, err = io.Copy(out, file)

	APP_URL := dot_env.GoDotEnvVariable("APP_URL")

	filepath := APP_URL + "/file" + fileName
	data, err := controller.bookService.FileUpload(id, filepath)

	ctx.JSON(http.StatusOK, gin.H{
		"meta": gin.H{
			"status":  true,
			"message": "File uploaded successfully",
		},
		"data": data,
	})

	// delete Path => os.remove(path)
}

func (controller *bookController) ChangeImage(ctx *gin.Context) {
	id := ctx.Param("id")
	idInt, _ := strconv.Atoi(id)

	file, header, err := ctx.Request.FormFile("file")
	// if err != nil {
	// 	var meta = gin.H{
	// 		"status":  false,
	// 		"message": "File not found",
	// 	}
	// 	ctx.JSON(http.StatusNotFound, gin.H{
	// 		"meta": meta,
	// 		"data": gin.H{},
	// 	})
	// 	return
	// }
	filename := header.Filename
	out, err := os.Create("public/" + filename)
	// if err != nil {
	// 	var meta = gin.H{
	// 		"status":  false,
	// 		"message": "File not found Public",
	// 	}
	// 	ctx.JSON(http.StatusNotFound, gin.H{
	// 		"meta": meta,
	// 		"data": gin.H{},
	// 	})
	// 	return
	// }
	defer out.Close()
	_, err = io.Copy(out, file)
	APP_URL := dot_env.GoDotEnvVariable("APP_URL")
	filepath := APP_URL + "/file" + filename
	data, err := controller.bookService.ChangeImage(idInt, filepath)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"meta": gin.H{
				"status":  false,
				"message": "File not found Public",
			},
			"data": gin.H{},
		})
	}
	ctx.JSON(http.StatusOK, gin.H{
		"meta": gin.H{
			"status":  true,
			"message": "File uploaded successfully",
		},
		"data": data,
	})
}
