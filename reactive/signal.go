package reactive

type SignalId int

type Signal[T any] struct {
	*Runtime
	SignalId
}

func NewSignal[T any](initial T) Signal[T] {
	rt := GetRuntime()
	id := SignalId(len(rt.signalValues))
	rt.signalValues = append(rt.signalValues, initial)
	return Signal[T]{rt, id}
}

func NewSignalWithRuntime[T any](rt *Runtime, initial T) Signal[T] {
	id := SignalId(len(rt.signalValues))
	rt.signalValues = append(rt.signalValues, initial)
	return Signal[T]{rt, id}
}

func (s *Signal[T]) Get() T {
	val := s.Runtime.signalValues[s.SignalId]

	if rt := s.Runtime; rt.runningEffect != nil {
		if _, ok := rt.signalSubscribers[s.SignalId]; !ok {
			rt.signalSubscribers[s.SignalId] = append(rt.signalSubscribers[s.SignalId], *rt.runningEffect)
		}
	}

	return val.(T)
}

func (s *Signal[T]) Set(v T) {
	s.Runtime.signalValues[s.SignalId] = v

	if subscribers, ok := s.Runtime.signalSubscribers[s.SignalId]; ok {
		for _, effectId := range subscribers {
			runEffect(s.Runtime, effectId)
		}
	}
}
