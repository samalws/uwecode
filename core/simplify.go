package core

const maxDepth = uint(10000)

func min(a uint, b uint) uint {
	if a > b {
		return b
	} else {
		return a
	}
}
func SimplifyUntilNoPanic(f func(Obj) (bool, interface{}), obj Obj, state SimplifyState) (bool, interface{}) {
	newObj := obj.Simplify(state)
	isGood, val := f(newObj)
	if isGood {
		return true, val
	} else {
		return false, newObj
	}
}
func SimplifyUntil(f func(Obj) (bool, interface{}), obj Obj, state SimplifyState) interface{} {
	b, val := SimplifyUntilNoPanic(f, obj, state)
	if !b {
		panic("Failed to simplify to expected value")
	} else {
		return val
	}
}
