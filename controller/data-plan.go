package controller

import (
	"mydiary_api/helper"
	"mydiary_api/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddDataPlan(context *gin.Context) {
	var input model.DataPlan
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

	savedDataPlan, err := input.Save()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	context.JSON(http.StatusOK, gin.H{
		"data": savedDataPlan,
	})
}

func GetAllDataPlan(context *gin.Context) {
	_, err := helper.CurrentUser(context)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	dataPlans, err := model.FindAllDataPlan()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
	}

	context.JSON(http.StatusOK, gin.H{
		"data": dataPlans,
	})
}
