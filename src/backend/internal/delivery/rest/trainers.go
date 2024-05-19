package rest

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sachatarba/course-db/internal/delivery/request"
	"github.com/sachatarba/course-db/internal/entity"
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

func (h *Handler) CreateNewTrainer(ctx *gin.Context) {
	log.Print("CreateNewTrainer:", ctx.Request)

	var req request.TrainerReq

	err := ctx.BindJSON(&req)
	if err != nil {
		log.Print(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})

		return
	}

	err = h.trainerService.RegisterNewTrainer(ctx.Request.Context(), entity.Trainer{
		ID: req.ID,
		Fullname: req.Fullname,
		Email: req.Email,
		Phone: req.Phone,
		Qualification: req.Qualification,
		UnitPrice: req.UnitPrice,
		GymsID: req.GymsID,
	})
	if err != nil {
		log.Print(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})

		return
	}

	ctx.Status(http.StatusOK)
}

func (h *Handler) ChangeTrainer(ctx *gin.Context) {
	log.Print("ChangeTrainer: ", ctx.Request)

	var req request.TrainerReq
	err := ctx.BindJSON(&req)
	if err != nil {
		log.Print(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})

		return
	}

	err = h.trainerService.ChangeTrainer(ctx.Request.Context(), entity.Trainer{
		ID: req.ID,
		Fullname: req.Fullname,
		Email: req.Email,
		Phone: req.Phone,
		Qualification: req.Qualification,
		UnitPrice: req.UnitPrice,
		GymsID: req.GymsID,
	})
	if err != nil {
		log.Print(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})

		return
	}

	ctx.Status(http.StatusOK)
}

func (h *Handler) DeleteTrainer(ctx *gin.Context) {
	log.Print("DeleteTrainer request: ", ctx.Request)

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

	err = h.trainerService.DeleteTrainer(ctx.Request.Context(), uuID)
	if err != nil {
		log.Print(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})

		return
	}

	ctx.Status(http.StatusOK)
}
