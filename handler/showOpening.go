package handler

import (
	"errors"
	"net/http"

	"github.com/ThuanyMendonca/gooportunities/schemas"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ShowOpeningHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}

	opening := schemas.Opening{}
	if err := db.First(&opening, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			sendError(ctx, http.StatusNotFound, "opening not found")
			return
		}
		sendError(ctx, http.StatusInternalServerError, "internal server error")
		return
	}

	sendSuccess(ctx, "show-opening", opening)
}
