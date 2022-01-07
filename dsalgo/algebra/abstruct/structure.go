package abstract




type Monoid struct {
	Op func(a, b interface{}) interface{}
	E func() interface{}
	Commutative bool
}


type Group struct {
	Op func(a, b interface{}) interface{}
	E func() interface{}
	Commutative bool
	Inverse func(interface{}) interface{}
}