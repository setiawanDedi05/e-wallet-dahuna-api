package controller

import (
	"mydiary_api/helper"
	"mydiary_api/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddTip(context *gin.Context) {
	var input model.Tip

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

	savedTip, err := input.Save()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	context.JSON(http.StatusCreated, gin.H{
		"data": savedTip,
	})
}

func GetAllTip(context *gin.Context) {
	_, err := helper.CurrentUser(context)

	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}

	tips, err := model.FindAllTip()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"data": tips,
	})
}
