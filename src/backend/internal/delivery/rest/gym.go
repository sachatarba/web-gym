package rest

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sachatarba/course-db/internal/delivery/request"
	"github.com/sachatarba/course-db/internal/entity"
)

func (h *Handler) ListGyms(ctx *gin.Context) {
	log.Print("GetGyms request:", ctx.Request)

	gyms, err := h.gymService.ListGyms(ctx.Request.Context())
	if err != nil {
		log.Print(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{"gyms": gyms})
}

func (h *Handler) CreateNewGym(ctx *gin.Context) {
	log.Print("CreateNewGym request:", ctx.Request)

	var req request.GymReq

	err := ctx.BindJSON(&req)
	if err != nil {
		log.Print(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})

		return
	}

	err = h.gymService.RegisterNewGym(ctx.Request.Context(), entity.Gym{
		ID:      req.ID,
		Name:    req.Name,
		Phone:   req.Phone,
		City:    req.City,
		Addres:  req.Addres,
		IsChain: req.IsChain,
	})
	if err != nil {
		log.Print(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})

		return
	}

	ctx.Status(http.StatusOK)
}

func (h *Handler) ChangeGym(ctx *gin.Context) {
	log.Print("ChangeGym request: ", ctx.Request)

	var req request.GymReq
	err := ctx.BindJSON(&req)
	if err != nil {
		log.Print(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})

		return
	}

	err = h.gymService.ChangeGym(ctx.Request.Context(), entity.Gym{
		ID:      req.ID,
		Name:    req.Name,
		Phone:   req.Phone,
		City:    req.City,
		Addres:  req.Addres,
		IsChain: req.IsChain,
	})
	if err != nil {
		log.Print(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})

		return
	}

	ctx.Status(http.StatusOK)
}

func (h *Handler) DeleteGym(ctx *gin.Context) {
	log.Print("ChangeGym request: ", ctx.Request)

	id, ok := ctx.Keys["id"]
	if !ok {
		log.Print()
		ctx.JSON(http.StatusInternalServerError, gin.H{"err": ErrNoKeyInRequest.Error()})
	}

	uuID, err := uuid.Parse(id.(string))
	if err != nil {
		log.Print(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})

		return
	}

	err = h.gymService.DeleteGym(ctx.Request.Context(), uuID)
	if err != nil {
		log.Print(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})

		return
	}

	ctx.Status(http.StatusOK)
}
