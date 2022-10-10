package kitchen_elem

func SendFoodsToCooks(kitchenOrder OrderInKitchen) {
	for i, foodToCook := range kitchenOrder.Foods {

		NrFoodsQueue++
		foodID := foodToCook.FoodId

		switch complexity := Foods[foodID-1].Complexity; complexity {

		case 3:
			select {
			case Cooks[0].ProfficiencyChan <- 1:
				kitchenOrder.Foods[i].CookId = 1

				Cooks[0].FoodChan <- foodToCook

			default:
				Cooks[0].ProfficiencyChan <- 1
				kitchenOrder.Foods[i].CookId = 1
				Cooks[0].FoodChan <- foodToCook
			}

		case 2:
			select {
			case Cooks[0].ProfficiencyChan <- 1:

				kitchenOrder.Foods[i].CookId = 1
				Cooks[0].FoodChan <- foodToCook

			case Cooks[1].ProfficiencyChan <- 1:

				kitchenOrder.Foods[i].CookId = 2
				Cooks[1].FoodChan <- foodToCook

			case Cooks[2].ProfficiencyChan <- 1:

				kitchenOrder.Foods[i].CookId = 3
				Cooks[2].FoodChan <- foodToCook

			default:
				Cooks[1].ProfficiencyChan <- 1
				kitchenOrder.Foods[i].CookId = 2
				Cooks[1].FoodChan <- foodToCook
			}

		case 1:
			select {
			case Cooks[0].ProfficiencyChan <- 1:

				kitchenOrder.Foods[i].CookId = 1
				Cooks[0].FoodChan <- foodToCook

			case Cooks[1].ProfficiencyChan <- 1:

				kitchenOrder.Foods[i].CookId = 2
				Cooks[1].FoodChan <- foodToCook

			case Cooks[2].ProfficiencyChan <- 1:

				kitchenOrder.Foods[i].CookId = 3
				Cooks[2].FoodChan <- foodToCook

			case Cooks[3].ProfficiencyChan <- 1:

				kitchenOrder.Foods[i].CookId = 4
				Cooks[3].FoodChan <- foodToCook
			default:
				//foodToCook.CookId = 1
				//kitchen_elem.Cooks[0].FoodChan <- foodToCook
				Cooks[3].ProfficiencyChan <- 1
				kitchenOrder.Foods[i].CookId = 4
				Cooks[3].FoodChan <- foodToCook
			}

		}

	}
}
