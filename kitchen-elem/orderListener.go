package kitchen_elem

import (
	"github.com/EliriaT/kitchen/dataStructures"
	"sync"
)

func ListenForOrders() {

	autoLockMutex := false
	var lowestPriority uint8 = 5 //highest is 1
	queue := dataStructures.NewHierarchicalQueue(lowestPriority, autoLockMutex)

	//constantly listening for this channel
	for order := range OrdersChannel {

		//generating a kitchen order
		kitchenOrder := OrderInKitchen{
			Id:            order.OrderId,
			Foods:         make([]FoodToCook, 0, len(order.Items)),
			ReceivedOrder: order,
			// new creates a pointer to WaitGroup
			Wg:       new(sync.WaitGroup),
			Priority: uint8(order.Priority),
		}

		//to wait for food prep, we add number of foods to wait for
		kitchenOrder.Wg.Add(len(order.Items))

		//generating the food to cook by cooks, this will be sent to cooks
		for _, foodId := range order.Items {
			newFood := FoodToCook{
				OrderId:          order.OrderId,
				FoodId:           foodId,
				CookingApparatus: Foods[foodId-1].CookingApparatus,
				PrepTime:         Foods[foodId-1].PreparationTime,
				//kitchenOrder.Wg is already a pointer
				Wg: kitchenOrder.Wg,
			}
			kitchenOrder.Foods = append(kitchenOrder.Foods, newFood)
		}

		//push to the priority queue
		queue.Enqueue(kitchenOrder, kitchenOrder.Priority)

		//If no cook is free, then take another order
		// TODO HERE BETTER USE PROFFICIENCY CHANNEL ?
		// TODO THREAD FOR EACH FOOD COMPLEXITY
		select {
		case CookFree <- 1:
			<-CookFree

		default:
			continue
		}

		//if queue.Len() < 3 {
		//	continue
		//}

		//TODO have here a load balancing

		//take the order with best priority (1 the best))
		orderInterface, _ := queue.Dequeue()
		kitchenOrder, _ = orderInterface.(OrderInKitchen)

		SendFoodsToCooks(kitchenOrder)

		// waiting for the foods to be prepared
		go kitchenOrder.WaitForOrder(kitchenOrder.Foods)

	}
}
