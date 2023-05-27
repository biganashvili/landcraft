// Package main runs the landcraft and performs an Order
package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/biganashvili/landcraft/cmd/api/config"
	"github.com/biganashvili/landcraft/cmd/api/handler"
	"github.com/biganashvili/landcraft/domain/land"
	"github.com/biganashvili/landcraft/services/app"
	"github.com/biganashvili/landcraft/services/order"
	"github.com/biganashvili/landcraft/services/user"
)

func main() {

	us, err := user.NewUserService(
		user.WithMemoryUserRepository(),
	)
	if err != nil {
		log.Fatal(err)
	}
	lands := landInventory()

	os, err := order.NewOrderService(
		order.WithMemoryOrderRepository(),
		order.WithMemoryLandRepository(lands),
		order.WithMemoryUserRepository(us.GetUserRepo()),
	)
	if err != nil {
		log.Fatal(err)
	}

	tt, err := app.NewLandcraft(
		app.WithUserService(us),
		app.WithOrderService(os),
	)
	if err != nil {
		log.Fatal(err)
	}
	config.App = tt

	http.HandleFunc("/user", handler.AddUser)
	http.HandleFunc("/user/list", handler.UsersList)

	http.HandleFunc("/order", handler.AddOrder)
	http.HandleFunc("/order/list", handler.OrdersList)

	fmt.Printf("Starting server at port 8080\n")
	for i, l := range lands {
		fmt.Printf("land %d: %s\n", i, l.GetID())
	}
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}

	// lands := landInventory()
	// // Create Order Service to use in landcraft
	// os, err := order.NewOrderService(
	// 	order.WithMemoryUserRepository(),
	// 	order.WithMemoryLandRepository(lands),
	// )
	// if err != nil {
	// 	panic(err)
	// }
	// // Create landcraft service
	// landcraft, err := app.NewLandcraft(
	// 	app.WithOrderService(os),
	// )
	// if err != nil {
	// 	panic(err)
	// }

	// uid, err := os.AddUser("Sergi")
	// if err != nil {
	// 	panic(err)
	// }
	// order := []uuid.UUID{
	// 	lands[0].GetID(),
	// }
	// // Execute Order
	// err = landcraft.Order(uid, order)
	// if err != nil {
	// 	panic(err)
	// }
}

func landInventory() []land.Land {
	landNearLake, err := land.NewLand("Land near lake", 1.99)
	if err != nil {
		panic(err)
	}
	landNearSea, err := land.NewLand("Land near sea", 0.99)
	if err != nil {
		panic(err)
	}
	landNearForest, err := land.NewLand("Land near forest", 0.99)
	if err != nil {
		panic(err)
	}
	lands := []land.Land{
		landNearLake, landNearSea, landNearForest,
	}
	return lands
}
