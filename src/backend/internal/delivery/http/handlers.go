package http

import (
	"github.com/gin-gonic/gin"
	"github.com/sachatarba/course-db/internal/service"
)

type Handler struct {
	clientService        service.IClientService
	authorizationService service.IAuthorizationService
	gymService           service.IGymService
	equipmentService     service.IEquipmentService
}

type Services struct {
	ClientService        service.IClientService
	AuthorizationService service.IAuthorizationService
	GymService           service.IGymService
	EquipmentService     service.IEquipmentService
}

func NewHandler(services Services) *Handler {
	return &Handler{
		clientService:        services.ClientService,
		authorizationService: services.AuthorizationService,
		gymService:           services.GymService,
		equipmentService:     services.EquipmentService,
	}
}

func (h *Handler) Init(api *gin.RouterGroup) {
	v1 := api.Group("/v1")
	{
		v1.POST("/register", h.RegisterNewUser)
		v1.POST("/login", h.Login)

		v1.GET("/isauthorize", h.IsAuthorize)
		v1.GET("/logout", h.Logout)
		v1.GET("/clients", h.GetClients)
	}
}
