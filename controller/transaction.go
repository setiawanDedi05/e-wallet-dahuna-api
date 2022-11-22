package controller

import (
	"mydiary_api/helper"
	"mydiary_api/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddTransaction(context *gin.Context) {
	var input model.Transaction
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	_, err := helper.CurrentUser(context)

	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}

	savedTransaction, err := input.Save()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	context.JSON(http.StatusCreated, gin.H{
		"data": savedTransaction,
	})
}

func GetAllTransacion(context *gin.Context) {
	_, err := helper.CurrentUser(context)

	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}

	transactions, err := model.FindAllTransaction()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"data": transactions,
	})
}
