package factorymethod

type NYCheesePizza struct {
	description string
}

func (p *NYCheesePizza) GetDescription() string {
	return p.description
}

func (p *NYCheesePizza) prepare() {

}

func (p *NYCheesePizza) bake() {

}

func (p *NYCheesePizza) cut() {

}

func (p *NYCheesePizza) box() {

}
