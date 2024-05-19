package rest

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sachatarba/course-db/internal/delivery/request"
	"github.com/sachatarba/course-db/internal/entity"
)

func (h *Handler) IsAuthorize(ctx *gin.Context) {
	log.Print("IsAuthorize request:", ctx.Request)

	session, err := ctx.Cookie("session")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	id, err := uuid.Parse(session)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	ok, err := h.authorizationService.IsAuthorize(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	if !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "authorized"})
}

func (h *Handler) Logout(ctx *gin.Context) {
	log.Print("Logout request:", ctx.Request)

	session, err := ctx.Cookie("session")
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"err": err.Error()})
	}

	ctx.SetCookie(
		"session",
		session,
		0,
		"/",
		"",
		false,
		true,
	)
	ctx.Status(http.StatusOK)
}

func (h *Handler) Login(ctx *gin.Context) {
	log.Print("Login request:", ctx.Request)

	var req request.LoginReq

	err := ctx.BindJSON(&req)
	if err != nil {
		log.Print(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})

		return
	}

	session, err := h.authorizationService.Authorize(ctx.Request.Context(), req.Login, req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})

		return
	}

	ctx.SetCookie(
		"session",
		session.SessionID.String(),
		int(session.TTL.UnixMilli()-time.Now().UnixMilli()),
		"/",
		"",
		false,
		true,
	)
	ctx.Status(http.StatusOK)
}

func (h *Handler) RegisterNewUser(ctx *gin.Context) {
	log.Print("RegisterNewUser request:", ctx.Request)

	var req request.ClientReq

	err := ctx.BindJSON(&req)
	if err != nil {
		log.Print(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})

		return
	}

	session, err := h.authorizationService.Register(ctx.Request.Context(), entity.Client{
		ID:        req.ID,
		Login:     req.Login,
		Password:  req.Password,
		Fullname:  req.Fullname,
		Email:     req.Email,
		Phone:     req.Phone,
		Birthdate: req.Birthdate,
	})

	if err != nil {
		log.Print(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})

		return
	}

	ctx.SetCookie(
		"session",
		session.SessionID.String(),
		session.TTL.Second(),
		"/",
		"",
		false,
		true,
	)
	ctx.Status(http.StatusOK)
}
