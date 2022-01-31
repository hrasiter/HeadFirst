package factorymethod

type ChicagoCheesePizza struct {
	description string
}

func (p *ChicagoCheesePizza) GetDescription() string {
	return p.description
}

func (p *ChicagoCheesePizza) prepare() {

}

func (p *ChicagoCheesePizza) bake() {

}

func (p *ChicagoCheesePizza) cut() {

}

func (p *ChicagoCheesePizza) box() {

}
