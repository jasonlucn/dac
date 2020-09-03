package dac

type FSA struct {
	state      int
	stateTable map[int]int
}

func (f *FSA) Init() {

}
