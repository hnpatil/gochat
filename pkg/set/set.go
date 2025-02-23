package set

type Set[T comparable] struct {
	values map[T]struct{}
}

func New[T comparable](values ...T) *Set[T] {
	valueMap := make(map[T]struct{}, len(values))
	for _, value := range values {
		valueMap[value] = struct{}{}
	}

	return &Set[T]{values: valueMap}
}

func (s *Set[T]) Values() []T {
	values := make([]T, 0, len(s.values))
	for v := range s.values {
		values = append(values, v)
	}

	return values
}

func (s *Set[T]) Size() int {
	return len(s.values)
}

func (s *Set[T]) Add(value T) {
	s.values[value] = struct{}{}
}
