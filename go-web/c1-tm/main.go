package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

type transaction struct {
	Id       int     `json:"id"`
	Code     string  `json:"code" binding:"required" validate:"required"`
	Currency string  `json:"currency" binding:"required" validate:"required"`
	Amount   float64 `json:"amount" binding:"required" validate:"required"`
	Emitter  string  `json:"emitter" binding:"required" validate:"required"`
	Receiver string  `json:"receiver" binding:"required" validate:"required"`
	Date     string  `json:"date" binding:"required" validate:"required"`
}

func getAll(c *gin.Context) {
	jsonTransactions, err := os.ReadFile("./exercise-1/transactions.json")
	if err != nil {
		c.JSON(500, err)
	}
	var transactions []transaction
	if err := json.Unmarshal(jsonTransactions, &transactions); err != nil {
		c.JSON(500, err)
	}
	var filteredTransactions []transaction
	if c.Query("code") != "" {
		for _, transaction := range transactions {
			if c.Query("code") == transaction.Code {
				filteredTransactions = append(filteredTransactions, transaction)
			}
		}
		c.JSON(200, filteredTransactions)
		return
	}
	if c.Query("currency") != "" {
		for _, transaction := range transactions {
			if c.Query("currency") == transaction.Currency {
				filteredTransactions = append(filteredTransactions, transaction)
			}
		}
		c.JSON(200, filteredTransactions)
		return
	}
	if c.Query("amount") != "" {
		for _, transaction := range transactions {
			if amount, _ := strconv.ParseFloat(c.Query("amount"), 64); amount == transaction.Amount {
				filteredTransactions = append(filteredTransactions, transaction)
			}
		}
		c.JSON(200, filteredTransactions)
		return
	}
	if c.Query("emitter") != "" {
		for _, transaction := range transactions {
			if c.Query("emitter") == transaction.Emitter {
				filteredTransactions = append(filteredTransactions, transaction)
			}
		}
		c.JSON(200, filteredTransactions)
		return
	}
	if c.Query("receiver") != "" {
		for _, transaction := range transactions {
			if c.Query("receiver") == transaction.Receiver {
				filteredTransactions = append(filteredTransactions, transaction)
			}
		}
		c.JSON(200, filteredTransactions)
		return
	}
	if c.Query("date") != "" {
		for _, transaction := range transactions {
			if c.Query("date") == transaction.Date {
				filteredTransactions = append(filteredTransactions, transaction)
			}
		}
		c.JSON(200, filteredTransactions)
		return
	}
	c.JSON(200, transactions)
}

var transactions []transaction

func getByID(c *gin.Context) {
	jsonTransactions, err := os.ReadFile("./exercise-1/transactions.json")
	if err != nil {
		c.JSON(500, err)
	}
	var transactions []transaction
	if err := json.Unmarshal(jsonTransactions, &transactions); err != nil {
		c.JSON(500, err)
	}
	for _, transaction := range transactions {
		if id, _ := strconv.Atoi(c.Param("id")); id == transaction.Id {
			c.JSON(200, transaction)
			return
		}
		c.JSON(404, nil)
	}

	return

}

func store(c *gin.Context) {
	if err := validateToken(c); err != nil {
		c.JSON(400, gin.H{"Error": err.Error()})
		return
	}
	var req transaction

	var validate *validator.Validate
	validate = validator.New()

	if err := c.ShouldBindJSON(&req); err != nil {
		fmt.Print(req)
		if err := validate.Struct(&req); err != nil {
			fmt.Print("err")
			for _, err := range err.(validator.ValidationErrors) {
				customError := fmt.Errorf("el campo %s es %s", err.Field(), err.Tag())
				c.JSON(400, gin.H{"Error": customError.Error()})
			}
			return
		}
	}

	req.Id = len(transactions) + 1
	transactions = append(transactions, req)
	return
}

func validateToken(c *gin.Context) (err error) {
	token := c.GetHeader("token")
	if token != "1234" {
		err = fmt.Errorf("token inválido")
		return
	}
	return
}

func main() {
	router := gin.Default()

	router.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hola Luri",
		})
	})

	router.GET("transactions/:id", getByID)
	router.GET("/transactions", getAll)
	router.POST("/transactions", store)

	router.Run()
}
