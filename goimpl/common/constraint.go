package common

// VType is an interface that is used to define the GetCopy method
// if T passed by value then GetCopy should return T by value
// if T passed by pointer then GetCopy should return T by pointer i.e. *T
type VType[T any] interface {
	GetCopy() T
}

func CopyOf[V VType[V]](v V) V {
	return v.GetCopy()
}

type PtrType[T any] interface {
	*T
}

func PtrOf[T any, PT PtrType[T]](t T) PT {
	return PT(&t)
}

func PtrInstanceOf[T any, PT PtrType[T]]() PT {
	return PT(new(T))
}
