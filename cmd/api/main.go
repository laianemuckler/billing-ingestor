package main

import (
	"billing-ingestor/internal/pulses/http/handler"
	"billing-ingestor/internal/pulses/http/routes"
	"billing-ingestor/internal/pulses/repository/memory"
	"billing-ingestor/internal/pulses/service"
	"billing-ingestor/internal/pulses/usecase"
	"billing-ingestor/pkg/fakequeue"
	"log"

	"github.com/labstack/echo/v4"
)

func main() {
	pulseRepository := memory.NewPulseRepository()

	pulseRepository.SeedAggregatedData()

	fakeQueue := fakequeue.NewFakeQueue(service.NewPulseService(pulseRepository))

	var pulseUsecase usecase.PulseUsecase = service.NewPulseService(pulseRepository)

	pulseHandler := handler.NewPulseHandler(pulseUsecase)

	e := echo.New()

	routes.RegisterRoutes(e, pulseHandler)

	go func() {
		fakeQueue.Start()
	}()

	log.Fatal(e.Start(":8081"))
}
