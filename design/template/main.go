package template


type Cooker interface {
	open()
	fire()
	doCook(name string)
	outfire()
	close()
}


type xihongshi struct {

}

func (xihongshi) doCook(name string) {
	panic("implement me")
}

func (xihongshi) open() {
	panic("implement me")
}

func (xihongshi) fire() {
	panic("implement me")
}

func (xihongshi) outfire() {
	panic("implement me")
}

func (xihongshi) close() {
	panic("implement me")
}

func (xi xihongshi) do(name string)  {
	xi.open()
	xi.fire()
	xi.doCook(name)
	xi.outfire()
	xi.close()
}