package main

import (
	"context"
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/google/uuid"
	"github.com/sachatarba/course-db/internal/entity"
	"github.com/sachatarba/course-db/internal/orm"
	"github.com/sachatarba/course-db/internal/repository"
)

func main() {
	host := os.Getenv("HOST")
	pswd := os.Getenv("POSTGRES_PASSWORD")
	usr := os.Getenv("POSTGRES_USER")
	dbName := os.Getenv("POSTGRES_DB")

	dsn := fmt.Sprintf("host=%s user=%s dbname=%s password=%s sslmode=disable", host, usr, dbName, pswd)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Err while openning db: ", err)
	}

	err = db.AutoMigrate(
		&orm.Client{},
		&orm.Gym{},
		&orm.Equipment{},
		&orm.MembershipType{},
		&orm.ClientMembership{},
		&orm.Schedule{},
		&orm.Trainer{},
		&orm.Training{},
	)

	if err != nil {
		fmt.Println("Err while migrating db: ", err)
	}

	// clientID := uuid.New()
	// db.Create(&orm.Client{
	// 	ID:        clientID,
	// 	Fullname:  "Тарба Александр Вячеславович",
	// 	Email:     "sachatarba@rambler.ru",
	// 	Phone:     "+7-985-958-95-88",
	// 	Birthdate: "02-23-2003",
	// })

	gymID := uuid.New()
	db.Create(&orm.Gym{
		ID:      gymID,
		Name:    "Качалка №1",
		Phone:   "+7-999-999-99-99",
		City:    "Москва",
		Addres:  "Ангелов переулок, д. 10",
		IsChain: "false",
	})

	rep := repository.NewEquipmentRepo(db)
	err = rep.CreateNewEquipment(
		context.Background(),
		entity.Equipment{
			ID:          uuid.New(),
			Name:        "Планка для жима лежа",
			Description: "Для крепких качков",
			GymID:       gymID,
		},
	)

	eq, err := rep.ListEquipmentsByGymID(context.Background(), gymID)
	fmt.Println(eq)
	rep.ChangeEquipment(context.Background(),
		entity.Equipment{
			ID:          eq[0].ID,
			Name:        "Планка для жима стоя",
			Description: "Для очень крепких качков",
			GymID:       gymID,
		},
	)

	eq, err = rep.ListEquipmentsByGymID(context.Background(), gymID)
	fmt.Println(eq)

	e, err := rep.GetEquipmentByID(context.Background(), eq[0].ID)
	fmt.Println(e)
	err = rep.DeleteEquipment(context.Background(), e.ID)


	e, err = rep.GetEquipmentByID(context.Background(), eq[0].ID)
	fmt.Println(e, err)
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
