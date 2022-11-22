package controller

import (
	"mydiary_api/helper"
	"mydiary_api/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddPaymentMethod(context *gin.Context) {
	var input model.PaymentMethod
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

	savedPaymentMethod, err := input.Save()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"data": savedPaymentMethod,
	})

}

func GetAllPaymentMethods(context *gin.Context) {
	_, err := helper.CurrentUser(context)
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
		})
		return
	}

	paymentMethods, err := model.FindAllPaymentMethod()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"data": paymentMethods,
	})
}
