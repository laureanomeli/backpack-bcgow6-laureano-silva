package handler

import (
	"fmt"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/laureanomeli/backpack-bcgow6-laureano-silva/testing/c2-tt/internal/transactions"
	"github.com/laureanomeli/backpack-bcgow6-laureano-silva/testing/c2-tt/pkg/web"
)

type request struct {
	Code     string  `json:"code" binding:"required" validate:"required"`
	Currency string  `json:"currency" binding:"required" validate:"required"`
	Amount   float64 `json:"amount" binding:"required" validate:"required"`
	Emitter  string  `json:"emitter" binding:"required" validate:"required"`
	Receiver string  `json:"receiver" binding:"required" validate:"required"`
	Date     string  `json:"date" binding:"required" validate:"required"`
}

type requestModify struct {
	Code   string  `json:"code" binding:"required" validate:"required"`
	Amount float64 `json:"amount" binding:"required" validate:"required"`
}

type Transaction struct {
	service transactions.Service
}

func NewTransaction(t transactions.Service) *Transaction {
	return &Transaction{
		service: t,
	}
}

// @Summary List transactions
// @Tags Transactions
// @Description get transactions
// @Produce json
// @Param token header string true "token"
// @Success 200 {Object} web.Response
// @Failure 401 {Object} web.Response
// @Router /transactions [get]
func (c *Transaction) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if err := validateToken(ctx); err != nil {
			ctx.JSON(401, web.NewResponse(401, nil, err.Error()))
			return
		}

		p, err := c.service.GetAll()
		if err != nil {
			ctx.JSON(404, web.NewResponse(401, nil, err.Error()))
			return
		}
		ctx.JSON(200, web.NewResponse(200, p, ""))
	}
}

// @Summary Store a transaction
// @Tags Transactions
// @Description store a transaction
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Param transaction body request true "Transaction to store"
// @Success 200 {Object} web.Response
// @Failure 400 {Object} web.Response
// @Failure 401 {Object} web.Response
// @Router /transactions [post]
func (c *Transaction) Store() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if err := validateToken(ctx); err != nil {
			ctx.JSON(401, web.NewResponse(401, nil, err.Error()))
		}
		var req request

		var validate = validator.New()

		if err := ctx.ShouldBindJSON(&req); err != nil {
			fmt.Print(req)
			if err := validate.Struct(&req); err != nil {
				fmt.Print("err")
				for _, err := range err.(validator.ValidationErrors) {
					customError := fmt.Errorf("el campo %s es %s", err.Field(), err.Tag())
					ctx.JSON(400, web.NewResponse(400, nil, customError.Error()))
				}
				return
			}
		}
		transaction, err := c.service.Store(req.Code, req.Currency, req.Amount, req.Emitter, req.Receiver, req.Date)
		if err != nil {
			ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}
		ctx.JSON(200, web.NewResponse(200, transaction, ""))
	}
}

// @Summary Update a transaction
// @Tags Transactions
// @Description update a transaction fields
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Param id path int true "id"
// @Param transaction body request true "Transaction to update"
// @Success 200 {Object} web.Response
// @Failure 400 {Object} web.Response
// @Failure 401 {Object} web.Response
// @Failure 404 {Object} web.Response
// @Router /transactions/{id} [put]
func (c *Transaction) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if err := validateToken(ctx); err != nil {
			ctx.JSON(401, web.NewResponse(401, nil, err.Error()))
		}
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, err.Error()))
			return
		}
		var req request

		var validate = validator.New()

		if err := ctx.ShouldBindJSON(&req); err != nil {
			fmt.Print(req)
			if err := validate.Struct(&req); err != nil {
				fmt.Print("err")
				for _, err := range err.(validator.ValidationErrors) {
					customError := fmt.Errorf("el campo %s es %s", err.Field(), err.Tag())
					ctx.JSON(400, web.NewResponse(400, nil, customError.Error()))
				}
				return
			}
		}
		transaction, err := c.service.Update(int(id), req.Code, req.Currency, req.Amount, req.Emitter, req.Receiver, req.Date)
		if err != nil {
			ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}
		ctx.JSON(200, web.NewResponse(200, transaction, ""))
	}
}

// @Summary Delete a transaction
// @Tags Transactions
// @Description Delete a transaction
// @Param token header string true "token"
// @Param id path int true "id"
// @Success 200
// @Failure 400 {Object} web.Response
// @Failure 401 {Object} web.Response
// @Failure 404 {Object} web.Response
// @Router /transactions/{id} [delete]
func (t *Transaction) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if err := validateToken(ctx); err != nil {
			ctx.JSON(401, web.NewResponse(401, nil, ""))
			return
		}
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, "invalid id"))
			return
		}
		if err := t.service.Delete(int(id)); err != nil {
			ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}
		fmt.Print("entra")
		ctx.JSON(204, web.NewResponse(204, "ok", ""))
	}
}

// @Summary Update a transaction
// @Tags Transactions
// @Description Update code and amount of a transaction
// @Param token header string true "token"
// @Param id path int true "id"
// @Param transaction body requestModify true "Transaction to update caode and amount"
// @Success 200 {Object} web.Response
// @Failure 400 {Object} web.Response
// @Failure 401 {Object} web.Response
// @Failure 404 {Object} web.Response
// @Router /transactions/{id} [patch]
func (t *Transaction) UpdateCodeAndAmount() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if err := validateToken(ctx); err != nil {
			ctx.JSON(401, web.NewResponse(401, nil, err.Error()))
		}
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, err.Error()))
			return
		}

		var reqModify requestModify

		var validate = validator.New()

		if err := ctx.ShouldBindJSON(&reqModify); err != nil {
			fmt.Print(reqModify)
			if err := validate.Struct(&reqModify); err != nil {
				fmt.Print("err")
				for _, err := range err.(validator.ValidationErrors) {
					customError := fmt.Errorf("el campo %s es %s", err.Field(), err.Tag())
					ctx.JSON(400, web.NewResponse(400, nil, customError.Error()))
				}
				return
			}
		}
		transaction, err := t.service.UpdateCodeAndAmount(int(id), reqModify.Code, reqModify.Amount)
		if err != nil {
			ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}
		ctx.JSON(200, web.NewResponse(200, transaction, ""))
	}
}

func validateToken(c *gin.Context) (err error) {
	token := c.GetHeader("token")
	if token != os.Getenv("TOKEN") {
		err = fmt.Errorf("token inválido")
		return
	}
	return
}
