package controller

import (
	"mydiary_api/helper"
	"mydiary_api/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddTransferHistory(context *gin.Context) {
	var input model.TransferHistory
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

	savedTranferHistory, err := input.Save()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	context.JSON(http.StatusCreated, gin.H{
		"data": savedTranferHistory,
	})
}

func GetAllTransferHistory(context *gin.Context) {
	_, err := helper.CurrentUser(context)

	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}

	transferHistories, err := model.FindAllTransferHistory()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"data": transferHistories,
	})
}
