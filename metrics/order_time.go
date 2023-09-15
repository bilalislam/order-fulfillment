package metrics

// OrderTime is an anonymous metric that represents that an new order has reached target steps in the prder process
type OrderTime struct {
	Count int
	Tags  []Tag
}

// NewOrderTime creates a new OrderTime metric
func NewOrderTime(tags []Tag) OrderTime {
	return OrderTime{
		Count: 1,
		Tags:  tags,
	}
}
