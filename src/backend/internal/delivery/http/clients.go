package http

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetClients(ctx *gin.Context) {
	log.Print("GetClients request:", ctx.Request)

	clients, err := h.clientService.ListClients(ctx.Request.Context())
	if err != nil {
		log.Print(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
	}

	ctx.JSON(http.StatusOK, gin.H{"clients": clients})
}

func (h *Handler) GetClientByLogin(ctx *gin.Context) {
	log.Print("GetClientByLogin request:", ctx.Request)

	clients, err := h.clientService.ListClients(ctx.Request.Context())
	if err != nil {
		log.Print(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
	}

	ctx.JSON(http.StatusOK, gin.H{"clients": clients})
}
