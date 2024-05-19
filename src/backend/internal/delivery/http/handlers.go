package http

import (
	"github.com/gin-gonic/gin"
	"github.com/sachatarba/course-db/internal/service"
)

type Handler struct {
	authorizationService    service.IAuthorizationService
	clientMembershipService service.IClientMembershipsService
	clientService           service.IClientService
	equipmentService        service.IEquipmentService
	gymService              service.IGymService
	membershipTypeService   service.IMembershipTypeService
	scheduleService         service.IScheduleService
	trainerService          service.ITrainerService
	trainingService         service.ITrainingService
}

func NewHandler(services service.ApiServices) *Handler {
	return &Handler{
		authorizationService:    services.AuthorizationService,
		clientMembershipService: services.ClientMembershipService,
		clientService:           services.ClientService,
		equipmentService:        services.EquipmentService,
		gymService:              services.GymService,
		membershipTypeService:   services.MembershipTypeService,
		scheduleService:         services.ScheduleService,
		trainerService:          services.TrainerService,
		trainingService:         services.TrainingService,
	}
}

func (h *Handler) InitApi(router gin.IRouter) {
	api := router.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			v1.POST("/register", h.RegisterNewUser)
			v1.POST("/login", h.Login)

			v1.GET("/isauthorize", h.IsAuthorize)
			v1.GET("/logout", h.Logout)
			v1.GET("/clients", h.GetClients)
		}
	}
}
