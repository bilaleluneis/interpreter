package common

type VType[T any] interface {
	GetCopy() T
}

func CopyOf[V VType[V]](v V) V {
	return v.GetCopy()
}
