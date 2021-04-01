package alacpretty

// Gauge encapsulates everything required to create an adjustable gauge.
type Gauge struct {
	// Current value of gauge.
	Value float64

	// Lowest value of gauge.
	MinVal float64

	// Highest value of gauge.
	MaxVal float64
}

func InitGauge(value int) *Gauge {

}

func (g *gauge) Render() {

}

func (g *gauge) Update(key string) {

}

func (g *gauge) PercentFilled() float64 {

}

func (g *gauge) PercentFilledPresenter() string {

}

func (g *gauge) ShiftRight() {

}

func (g *gauge) ShiftLeft() {

}
