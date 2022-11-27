package runtime

type Value interface {
	~Primitive | ~Object | ~Empty
}

type Empty struct {
}

type Primitive interface {
	~int32 | ~int64 | ~byte | ~uint16 | ~bool | ~float32 | ~float64
}

type Object struct {
}
