package controller

import (
	"mydiary_api/helper"
	"mydiary_api/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddOperatorCard(context *gin.Context) {
	var input model.OperatorCard
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	_, err := helper.CurrentUser(context)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	savedOperatorCard, err := input.Save()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	context.JSON(http.StatusCreated, gin.H{
		"data": savedOperatorCard,
	})

}

func GetAllOperatorCards(context *gin.Context) {
	_, err := helper.CurrentUser(context)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	operatorCards, err := model.FindAllOperatorCard()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	context.JSON(http.StatusOK, gin.H{
		"data": operatorCards,
	})
}
