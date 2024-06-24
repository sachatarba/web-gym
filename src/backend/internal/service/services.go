package service

type ApiServices struct {
	AuthorizationService    IAuthorizationService
	ClientMembershipService IClientMembershipsService
	ClientService           IClientService
	EquipmentService        IEquipmentService
	GymService              IGymService
	MembershipTypeService   IMembershipTypeService
	ScheduleService         IScheduleService
	TrainerService          ITrainerService
	TrainingService         ITrainingService
}

type ApiRepos struct {
	SessionRepo          ISessionRepository
	ClientMembershipRepo IClientMembershipsRepository
	ClientRepo           IClientRepository
	EquipmentRepo        IEquipmentRepository
	GymRepo              IGymRepository
	MembershipTypeRepo   IMembershipTypeRepository
	ScheduleRepo         IScheduleRepository
	TrainerRepo          ITrainerRepository
	TrainingRepo         ITrainingRepository
}
