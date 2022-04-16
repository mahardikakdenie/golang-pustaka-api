package controller

import (
	"net/http"
	loan "pustaka-api/loan"
	"pustaka-api/middleware"
	"strconv"

	"github.com/gin-gonic/gin"
)

type loanController struct {
	loanService loan.Service
}

func NewLoanController(loanService loan.Service) *loanController {
	return &loanController{loanService}
}

func (controller *loanController) Index(ctx *gin.Context) {
	data, err := controller.loanService.FindAll()
	if err != nil {
		var meta = gin.H{
			"status":  false,
			"message": err.Error(),
		}

		var data = gin.H{}
		ctx.JSON(http.StatusBadRequest, gin.H{
			"meta": meta,
			"data": data,
		})
		return
	}

	var bookResponse []loan.LoanResponse
	for _, v := range data {
		bookResponse = append(bookResponse, loan.LoanResponse{
			ID:         v.ID,
			UserId:     int(v.UserId),
			BookId:     int(v.BookId),
			DateReturn: v.DateReturn,
			User:       v.User,
			Book:       v.Book,
			Status:     v.Status,
		})
	}

	var meta = gin.H{
		"status":  true,
		"message": "Success",
	}

	ctx.JSON(http.StatusOK, gin.H{
		"meta": meta,
		"data": bookResponse,
	})
}

func (controller *loanController) Store(ctx *gin.Context) {
	var loanInput loan.LoanRequest
	if err := ctx.ShouldBindJSON(&loanInput); err != nil {
		var meta = gin.H{
			"status":  false,
			"message": err.Error(),
		}

		var data = gin.H{}
		ctx.JSON(http.StatusBadRequest, gin.H{
			"meta": meta,
			"data": data,
		})
		return
	}

	loanInput.UserId = middleware.UserId
	loan, err := controller.loanService.Create(loanInput)

	if err != nil {
		var meta = gin.H{
			"status":  false,
			"message": err.Error(),
		}

		var data = gin.H{}
		ctx.JSON(http.StatusBadRequest, gin.H{
			"meta": meta,
			"data": data,
		})
		return
	}

	var meta = gin.H{
		"status":  true,
		"message": "Success",
	}

	ctx.JSON(http.StatusOK, gin.H{
		"meta": meta,
		"data": loan,
	})
}

func (controller *loanController) Show(ctx *gin.Context) {
	idString := ctx.Param("id")
	id, _ := strconv.Atoi(idString)

	data, err := controller.loanService.FindById(id)
	if err != nil {
		var meta = gin.H{
			"status":  false,
			"message": err.Error(),
		}

		var data = gin.H{}
		ctx.JSON(http.StatusBadRequest, gin.H{
			"meta": meta,
			"data": data,
		})
		return
	}

	if data.ID == 0 {
		var meta = gin.H{
			"status":  false,
			"message": "Data not found",
		}

		var data = gin.H{}
		ctx.JSON(http.StatusBadRequest, gin.H{
			"meta": meta,
			"data": data,
		})
		return
	}

	var bookResponse loan.LoanResponse
	bookResponse = loan.LoanResponse{
		ID:         data.ID,
		UserId:     int(data.UserId),
		BookId:     int(data.BookId),
		DateReturn: data.DateReturn,
		User:       data.User,
		Book:       data.Book,
		Status:     data.Status,
	}

	var meta = gin.H{
		"status":  true,
		"message": "Success",
	}

	ctx.JSON(http.StatusOK, gin.H{
		"meta": meta,
		"data": bookResponse,
	})
}

func (controller *loanController) Update(ctx *gin.Context) {
	idString := ctx.Param("id")
	id, _ := strconv.Atoi(idString)

	var loanInput loan.LoanRequest
	if err := ctx.ShouldBindJSON(&loanInput); err != nil {
		var meta = gin.H{
			"status":  false,
			"message": err.Error(),
		}

		var data = gin.H{}
		ctx.JSON(http.StatusBadRequest, gin.H{
			"meta": meta,
			"data": data,
		})
		return
	}

	loanInput.UserId = middleware.UserId

	var LoanRequest loan.LoanRequest
	LoanRequest = loan.LoanRequest{
		UserId:     loanInput.UserId,
		BookId:     loanInput.BookId,
		DateReturn: loanInput.DateReturn,
		Status:     loanInput.Status,
	}

	loan, err := controller.loanService.Update(LoanRequest, id)

	if err != nil {
		var meta = gin.H{
			"status":  false,
			"message": err.Error(),
		}

		var data = gin.H{}
		ctx.JSON(http.StatusBadRequest, gin.H{
			"meta": meta,
			"data": data,
		})
		return
	}

	var meta = gin.H{
		"status":  true,
		"message": "Success",
	}

	ctx.JSON(http.StatusOK, gin.H{
		"meta": meta,
		"data": loan,
	})
}

func (controller *loanController) DeleteData(ctx *gin.Context) {
	idString := ctx.Param("id")
	id, _ := strconv.Atoi(idString)

	data, err := controller.loanService.DestroyLoan(id)
	if err != nil {
		var meta = gin.H{
			"status":  false,
			"message": err.Error(),
		}

		var data = gin.H{}
		ctx.JSON(http.StatusBadRequest, gin.H{
			"meta": meta,
			"data": data,
		})
		return
	}

	var meta = gin.H{
		"status":  true,
		"message": "Success",
	}

	ctx.JSON(http.StatusOK, gin.H{
		"meta": meta,
		"data": data,
	})
}
