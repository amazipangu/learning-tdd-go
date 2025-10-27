package money

type Dollar struct {
	amount int
}

func NewDollar(amount int) *Dollar {
	return &Dollar{amount: amount}
}

func (d *Dollar) Amount() int {
	return d.amount
}

func (d *Dollar) Times(multiplier int) *Dollar {
	return &Dollar{amount: d.amount * multiplier}
}
