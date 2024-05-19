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

// type ServiceDirector struct {
// 	Ctx            context.Context
// 	Config         config.Config
// 	ReposBuilder   IReposBuilder
// 	ServiceBuilder IServicesBuilder
// }

// func (d *ServiceDirector) NewServices() (*Services, error) {
// 	repos, err := d.ReposBuilder.NewRepos(d.Ctx, d.Config)
// 	if err != nil {
// 		return nil, err
// 	}

// 	services, err := d.ServiceBuilder.NewServices(d.Ctx, *repos)

// 	return services, err
// }

// type ApiServicesBuilder struct {
// 	Repos Repos
// }

// type ServiceBuilderCreator struct {
// 	ctx context.Context
// }

// func (creator *ServiceBuilderCreator) Create(repos Repos) IServicesBuilder {
// 	return &ServicesBuilder{
// 		Repos: repos,
// 	}
// }

// type ReposBuilderCreator struct {
// 	ctx context.Context
// }

// func (creator *ReposBuilderCreator) Create(conf config.Config) service.IReposBuilder{
// 	return &ReposBuilder{
// 		Conf: conf,
// 	}
// }

// func (br *ApiServicesBuilder) NewServices() (*ApiServices, error) {
// 	return &ApiServices{
// 		authorizationService:    NewAuthorizationService(br.Repos.SessionRepo, br.Repos.ClientRepo),
// 		clientMembershipService: NewClientMembershipService(br.Repos.ClientMembershipRepo),
// 		clientService:           NewClientService(br.Repos.ClientRepo),
// 		equipmentService:        NewEquipmentService(br.Repos.EquipmentRepo),
// 		gymService:              NewGymService(br.Repos.GymRepo),
// 		membershipTypeService:   NewMembershipTypeService(br.Repos.MembershipTypeRepo),
// 		scheduleService:         NewScheduleService(br.Repos.ScheduleRepo),
// 		trainerService:          NewTrainerService(br.Repos.TrainerRepo),
// 		trainingService:         NewTrainingService(br.Repos.TrainingRepo),
// 	}, nil
// }
