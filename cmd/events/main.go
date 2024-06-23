package main

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	httpHandler "github.com/Gustavohsdp/sales-go/internal/events/infra/http"
	"github.com/Gustavohsdp/sales-go/internal/events/infra/repository"
	"github.com/Gustavohsdp/sales-go/internal/events/infra/service"
	"github.com/Gustavohsdp/sales-go/internal/events/usecase"
)

func main() {
	db, err := sql.Open("mysql", "test_user:test_password@tcp(golang-mysql:3306)/test_db")
	if err != nil {
		panic(err)
	}

	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	eventRepo, err := repository.NewMysqlEventRepository(db)
	if err != nil {
		panic(err)
	}

	partnerBaseURLs := map[int]string{
		1: "http://localhost:3333/api/1",
		2: "http://localhost:3333/api/2",
	}

	partnerFactory := service.NewPartnerFactory(partnerBaseURLs)

	listEventsUseCase := usecase.NewListEventsUseCase(eventRepo)
	getEventUseCase := usecase.NewGetEventUseCase(eventRepo)
	listSpotsUseCase := usecase.NewListSpotsUseCase(eventRepo)
	buyTicketUseCase := usecase.NewBuyTicketsUseCase(eventRepo, partnerFactory)

	eventsHandler := httpHandler.NewEventsHandler(
		listEventsUseCase,
		listSpotsUseCase,
		getEventUseCase,
		buyTicketUseCase,
	)

	r := http.NewServeMux()
	r.HandleFunc("/events", eventsHandler.ListEvents)
	r.HandleFunc("/events/{eventID}", eventsHandler.GetEvent)
	r.HandleFunc("/events/{eventID}/spots", eventsHandler.ListSpots)
	r.HandleFunc("POST /checkout", eventsHandler.BuyTickets)

	log.Println("Server is running...")
	http.ListenAndServe(":3333", r)
}
