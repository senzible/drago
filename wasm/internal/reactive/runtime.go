package reactive

type Runtime struct {
	//slice of signal values
	signalValues      []any
	runningEffect     *EffectId
	signalSubscribers map[SignalId][]EffectId
	effects           []func()
}

func NewRuntime() *Runtime {
	return &Runtime{
		signalValues:      []any{},
		signalSubscribers: map[SignalId][]EffectId{},
		effects:           []func(){},
	}
}

func CreateSignal[T any](rt *Runtime, initial T) TypedSignal[T] {
	id := SignalId(len(rt.signalValues))
	rt.signalValues = append(rt.signalValues, initial)
	return TypedSignal[T]{rt, id}
}

func (rt *Runtime) CreateEffect(f func()) {
	effectId := EffectId(len(rt.effects))
	rt.effects = append(rt.effects, f)

	rt.runEffect(effectId)
}

func (rt *Runtime) runEffect(id EffectId) {
	//keep track of previous running effect
	previous := rt.runningEffect

	//set the current running effect
	rt.runningEffect = &id

	//run the effect
	rt.effects[id]()

	//reset the running effect
	rt.runningEffect = previous
}

type SignalId int
type EffectId int

type TypedSignal[T any] struct {
	*Runtime
	SignalId
}

func (s *TypedSignal[T]) Get() T {
	//get the value
	val := s.Runtime.signalValues[s.SignalId]

	//if there is a running effect add it to subscribers, only if it's not already subscribed
	if rt := s.Runtime; rt.runningEffect != nil {
		if _, ok := rt.signalSubscribers[s.SignalId]; !ok {
			rt.signalSubscribers[s.SignalId] = append(rt.signalSubscribers[s.SignalId], *rt.runningEffect)
		}
	}

	//return value
	return val.(T)
}

func (s *TypedSignal[T]) Set(v T) {
	//set the value
	s.Runtime.signalValues[s.SignalId] = v

	//notify subscribers
	for _, effectId := range s.Runtime.signalSubscribers[s.SignalId] {
		s.Runtime.runEffect(effectId)
	}
}
