package eago

type Selector interface {
	Select(indi Individuals) Individuals
}

type Tournament struct {
	NSelection uint
}

func (t Tournament) Select(indi Individuals) Individuals {

}