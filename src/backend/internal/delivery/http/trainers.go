package http

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h *Handler) GetListTrainersByGymID(ctx *gin.Context) {
	log.Print("GetListTrainersByGymID:", ctx.Request)

	id, ok := ctx.Keys["id"]
	if !ok {
		log.Print()
		ctx.JSON(http.StatusInternalServerError, gin.H{"err": ErrNoKeyInRequest.Error()})

		return
	}

	uuID, err := uuid.Parse(id.(string))
	if err != nil {
		log.Print(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})

		return
	}

	trainers, err := h.trainerService.ListTrainersByGymID(ctx.Request.Context(), uuID)
	if err != nil {
		log.Print(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{"trainers": trainers})
}
