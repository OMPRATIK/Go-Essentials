package main

type Order struct {
	ID     int
	Status string
}

func GenerateOrders(count int) []Order {
	orders := make([]Order, count)

	for idx := range orders {
		orders[idx] = Order{
			ID:     idx + 1,
			Status: "Pending",
		}
	}

	return orders
}