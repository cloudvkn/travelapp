package server

import (
    "github.com/go-micro/micro/v3"
    "github.com/go-micro/micro/v3/api"
    "github.com/cloudvkn/travelapp/handler"
)

func Run() {
    srv := micro.NewService(
        micro.Name("com.example.travelapp"),
        micro.Version("latest"),
    )

    srv.Init()

    apiHandler := api.NewHandler()
    apiHandler.HandleFunc("/booking", handler.Book)
    apiHandler.HandleFunc("/itinerary", handler.PlanItinerary)
    apiHandler.HandleFunc("/recommendations", handler.GetRecommendations)

    srv.Handle(apiHandler)

    if err := srv.Run(); err != nil {
        panic(err)
    }
}
