package main

import "github.com/sachatarba/course-db/internal/api"

func main() {
	api := api.ApiServer{}
	api.Run()
	// host := os.Getenv("POSTGRES_HOST")
	// pswd := os.Getenv("POSTGRES_PASSWORD")
	// usr := os.Getenv("POSTGRES_USER")
	// dbName := os.Getenv("POSTGRES_DB")
	// sslMode := os.Getenv("POSTGRES_SSLMODE")

	// dsn := fmt.Sprintf("host=%s user=%s dbname=%s password=%s sslmode=%s", host, usr, dbName, pswd, sslMode)

	// db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	// if err != nil {
	// 	fmt.Println("Err while openning db: ", err)
	// }

	// err = db.AutoMigrate(
	// 	&orm.Client{},
	// 	&orm.Gym{},
	// 	&orm.Equipment{},
	// 	&orm.MembershipType{},
	// 	&orm.ClientMembership{},
	// 	&orm.Schedule{},
	// 	&orm.Trainer{},
	// 	&orm.Training{},
	// )

	// if err != nil {
	// 	fmt.Println("Err while migrating db: ", err)
	// }

	// redisHost := os.Getenv("REDIS_HOST")
	// redistPort := os.Getenv("REDIS_PORT")

	// rdb := redis.NewClient(&redis.Options{
	// 	Addr: fmt.Sprintf("%s:%s", redisHost, redistPort),
	// })

	// sessionRepository    := repository.NewSessionRepo(rdb)
	// clientMembershipRepository := repository.NewClientMembershipRepo(db)
	// clientRepository           := repository.NewClientRepo(db)
	// equipmentRepository        := repository.NewEquipmentRepo(db)
	// gymRepository              := repository.NewGymRepo(db)
	// membershipTypeRepository   := repository.NewMembershipTypeRepo(db)
	// scheduleRepository         := repository.NewScheduleRepo(db)
	// trainerRepository          := repository.NewTrainerRepo(db)
	// trainingRepository         := repository.NewTrainingRepo(db)

	// authorizationService := service.NewAuthorizationService(sessionRepository, clientRepository)
	// clientMembershipService := service.NewClientMembershipService(clientMembershipRepository)
	// clientService           := service.NewClientService(clientRepository)
	// equipmentService        := service.NewEquipmentService(equipmentRepository)
	// gymService              := service.NewGymService(gymRepository)
	// membershipTypeService   := service.NewMembershipTypeService(membershipTypeRepository)
	// scheduleService         := service.NewScheduleService(scheduleRepository)
	// trainerService          := service.NewTrainerService(trainerRepository)
	// trainingService         := service.NewTrainingService(trainingRepository)

	// handler := handler.NewHandler(clientService, authorizationService)

	// router := gin.Default()
	// router.GET("/ping", func(ctx *gin.Context) {
	// 	ctx.Status(http.StatusOK)
	// })
	// handler.Init(router.Group("/api"))
	// server := &http.Server{
	// 	Addr:    ":8080",
	// 	Handler: router,
	// }
	// server.ListenAndServe()

	// go func() {
	// 	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
	// 		log.Fatalf("listen: %s\n", err)
	// 	}
	// }()

	// quit := make(chan os.Signal, 1)
	// signal.Notify(quit, os.Interrupt)
	// <-quit
	// log.Println("Shutdown Server ...")

	// ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	// defer cancel()
	// if err := server.Shutdown(ctx); err != nil {
	// 	log.Fatal("Server Shutdown:", err)
	// }
	// log.Println("Server exiting")
	// clientID := uuid.New()
	// db.Create(&orm.Client{
	// 	ID:        clientID,
	// 	Fullname:  "Тарба Александр Вячеславович",
	// 	Email:     "sachatarba@rambler.ru",
	// 	Phone:     "+7-985-958-95-88",
	// 	Birthdate: "02-23-2003",
	// })

	// gymID := uuid.New()
	// db.Create(&orm.Gym{
	// 	ID:      gymID,
	// 	Name:    "Качалка №1",
	// 	Phone:   "+7-999-999-99-99",
	// 	City:    "Москва",
	// 	Addres:  "Ангелов переулок, д. 10",
	// 	IsChain: "false",
	// })

	// rep := repository.NewEquipmentRepo(db)
	// err = rep.CreateNewEquipment(
	// 	context.Background(),
	// 	entity.Equipment{
	// 		ID:          uuid.New(),
	// 		Name:        "Планка для жима лежа",
	// 		Description: "Для крепких качков",
	// 		GymID:       gymID,
	// 	},
	// )

	// eq, err := rep.ListEquipmentsByGymID(context.Background(), gymID)
	// fmt.Println(eq)
	// rep.ChangeEquipment(context.Background(),
	// 	entity.Equipment{
	// 		ID:          eq[0].ID,
	// 		Name:        "Планка для жима стоя",
	// 		Description: "Для очень крепких качков",
	// 		GymID:       gymID,
	// 	},
	// )

	// eq, err = rep.ListEquipmentsByGymID(context.Background(), gymID)
	// fmt.Println(eq)

	// e, err := rep.GetEquipmentByID(context.Background(), eq[0].ID)
	// fmt.Println(e)
	// err = rep.DeleteEquipment(context.Background(), e.ID)

	// e, err = rep.GetEquipmentByID(context.Background(), eq[0].ID)
	// fmt.Println(e, err)
	// db.Create(&orm.Equipment{
	// 	ID:          uuid.New(),
	// 	Name:        "Планка для жима лежа",
	// 	Description: "Для крепких качков",
	// 	GymID:       gymID,
	// })

	// membershipTypeID := uuid.New()
	// db.Create(&orm.MembershipType{
	// 	ID:           membershipTypeID,
	// 	Type:         "Начальный",
	// 	Description:  "Для новичков",
	// 	Price:        "300",
	// 	DaysDuration: 10,
	// 	GymID:        gymID,
	// })

	// clientMembershipID := uuid.New()
	// db.Create(&orm.ClientMembership{
	// 	ID:               clientMembershipID,
	// 	StartDate:        "04-09-2024",
	// 	EndDate:          "05-09-2024",
	// 	MembershipTypeID: membershipTypeID,
	// 	ClientID:         clientID,
	// })

	// trainerID := uuid.New()
	// db.Create(&orm.Trainer{
	// 	ID:            trainerID,
	// 	Fullname:      "Степаненко Андрей Сергеевич",
	// 	Email:         "sas@rambler.ru",
	// 	Phone:         "+7-985-958-95-88",
	// 	Qualification: "Мастер-тренер",
	// 	UnitPrice:     300,
	// })

	// trainingID := uuid.New()
	// db.Create(&orm.Training{
	// 	ID:           trainingID,
	// 	Title:        "Жетская разминка",
	// 	Description:  "Начинающим как раз",
	// 	TrainingType: "aerobic",
	// 	TrainerID:    trainerID,
	// })

	// sheduleID := uuid.New()
	// start := time.Now()
	// end := start.Add(2 * time.Hour)

	// db.Create(&orm.Schedule{
	// 	ID:            sheduleID,
	// 	DayOfTheeWeek: "04-09-2024",
	// 	StartTime:     start.Format(time.DateTime),
	// 	EndTime:       end.Format(time.DateTime),
	// 	ClientID:      clientID,
	// 	TrainingID:    trainingID,
	// })

	// users := []orm.Client{}

	// db.Preload("Schedules.Training").Preload("Schedules").Preload("ClientMemberships.MembershipType").Preload("ClientMemberships").Find(&users)
	// fmt.Println(users[0].Schedules[0].Training)
}
