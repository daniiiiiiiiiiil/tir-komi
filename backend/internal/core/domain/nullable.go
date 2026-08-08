package domain

type Nullable[T any] struct {
	Value *T
	Set   bool
}

func (n Nullable[T]) IsSet() bool {
	return n.Set
}

func (n Nullable[T]) IsNull() bool {
	return n.Set && n.Value == nil
}

func (n Nullable[T]) Get() (T, bool) {
	if n.Set && n.Value != nil {
		return *n.Value, true
	}
	var zero T
	return zero, false
}
