package money

type Dollar struct {
	amount int
}

type Franc struct {
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

func (d *Dollar) Equal(other *Dollar) bool {
	return d.amount == other.amount
}

func NewFranc(amount int) *Franc {
	return &Franc{amount: amount}
}

func (f *Franc) Amount() int {
	return f.amount
}

func (f *Franc) Times(multiplier int) *Franc {
	return &Franc{amount: f.amount * multiplier}
}
