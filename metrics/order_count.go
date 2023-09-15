package metrics

// OrderCount is an anonymous metric that represents that an new order has been recieved and collects many products were in the order.
type OrderCount struct {
	Count int
	Tags  []Tag
}

// NewOrderCount creates a new OrderCount metric
func NewOrderCount(tags []Tag) OrderCount {
	return OrderCount{
		Count: 1,
		Tags:  tags,
	}
}
