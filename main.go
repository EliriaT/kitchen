package main

import (
	"github.com/EliriaT/kitchen/handlers"
	"github.com/EliriaT/kitchen/kitchen-elem"
	"github.com/gorilla/mux"
	"runtime"

	//_ "go.uber.org/automaxprocs"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	//fmt.Println(runtime.NumCPU())
	runtime.GOMAXPROCS(6)

	rand.Seed(time.Now().UnixNano())
	kitchen_elem.StartWorkDay()

	for i := 0; i < kitchen_elem.Stoves.Quantity; i++ {

		go kitchen_elem.Stoves.CookFood()
	}

	for i := 0; i < kitchen_elem.Ovens.Quantity; i++ {

		go kitchen_elem.Ovens.CookFood()
	}
	go kitchen_elem.ListenForOrders()

	for i, _ := range kitchen_elem.Cooks {
		go kitchen_elem.Cooks[i].ListenForFood()
	}

	//newroute for online orders
	r := mux.NewRouter()
	r.HandleFunc("/", handlers.GetCooks).Methods("GET")
	r.HandleFunc("/order", handlers.ReceiveOrder).Methods("POST")
	r.HandleFunc("/onlineOrder", handlers.ReceiveOnlineOrder).Methods("POST")

	log.Println("Kitchen server started..")
	log.Println("Quantum Apparatus: ", kitchen_elem.ApparatusQuantum)
	log.Fatal(http.ListenAndServe(kitchen_elem.Port, r))

}
