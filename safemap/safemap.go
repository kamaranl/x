package safemap

type SafeMap struct{ *typeSafeMap[any] }

func NewSafeMap() *SafeMap { return &SafeMap{typeSafeMap: newtypeSafeMap[any]()} }

func GetTypedFromSafeMap[T any](s *SafeMap, key string) (T, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	v, ok := s.Data[key]
	if !ok {
		var z T
		return z, ok
	}

	value, ok := v.(T)

	return value, ok
}
