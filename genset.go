package genset

type Set[T comparable] struct {
	m map[T]struct{}
}

func New[T comparable](s ...T) *Set[T] {
	set := &Set[T]{}
	set.Add(s...)
	return set
}

func (set *Set[T]) Add(s ...T) {
	if set.m == nil {
		set.m = make(map[T]struct{})
	}
	for _, s := range s {
		set.m[s] = struct{}{}
	}
}

func S(x ...string) *Set[string] {
	return New(x...)
}
