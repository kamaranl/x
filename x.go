package x

import (
	"strings"
	"sync"
)

var state = NewLockedMap()

type AlertLevel int

const (
	Informational AlertLevel = iota
	Warning
	Critical
)

type Alert struct {
	Title   string
	Message string
	Level   AlertLevel
	OK      bool

	label string
}

func NewAlert(title, message string, level AlertLevel) *Alert {
	label := strings.ToLower(strings.ReplaceAll(title, " ", ""))

	return &Alert{
		Title:   title,
		Message: message,
		Level:   level,
		label:   "alert&" + label,
	}
}

func (a *Alert) Show() {
	if active, ok := GetTypedFromLockedMap[bool](state, a.label); ok && active {
		return
	}
	state.Set(a.label, true)

	go func() {
		a.alert()
		state.Set(a.label, false)
	}()
}

type LockedMap struct {
	Data map[string]any
	m    sync.RWMutex
}

func NewLockedMap() *LockedMap {
	return &LockedMap{Data: make(map[string]any)}
}

func (s *LockedMap) Get(key string) (any, bool) {
	s.m.RLock()
	defer s.m.RUnlock()
	v, ok := s.Data[key]

	return v, ok
}

func (s *LockedMap) Set(key string, value any) {
	s.m.Lock()
	s.Data[key] = value
	s.m.Unlock()
}

func (s *LockedMap) Delete(key string) {
	s.m.Lock()
	delete(s.Data, key)
	s.m.Unlock()
}

func (s *LockedMap) Clear() {
	s.m.Lock()
	s.Data = make(map[string]any)
	s.m.Unlock()
}

func (s *LockedMap) Keys() []string {
	keys := make([]string, 0, len(s.Data))
	for k := range s.Data {
		keys = append(keys, k)
	}

	return keys
}

func GetTypedFromLockedMap[T any](s *LockedMap, key string) (T, bool) {
	s.m.RLock()
	defer s.m.RUnlock()

	v, ok := s.Data[key]
	if !ok {
		var z T
		return z, ok
	}

	value, ok := v.(T)

	return value, ok
}
