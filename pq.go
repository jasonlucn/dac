package dac

//优先级队列
type PQI interface {
	GetMax() interface{}
	DelMax()
	Insert(interface{})
}
