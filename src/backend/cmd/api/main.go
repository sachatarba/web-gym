package main

import "github.com/sachatarba/course-db/internal/api"

func main() {
	api := api.ApiServer{}
	api.Run()
}
