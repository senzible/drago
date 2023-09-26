package reactive

import "sync"

type Runtime struct {
	signalValues      []any
	runningEffect     *EffectId
	signalSubscribers map[SignalId][]EffectId
	effects           []func()
}

var lock = &sync.Mutex{}
var runtimeInstance *Runtime

func GetRuntime() *Runtime {
	if runtimeInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if runtimeInstance == nil {
			runtimeInstance = newRuntime()
		}
	}
	return runtimeInstance
}

func newRuntime() *Runtime {
	return &Runtime{
		signalValues:      []any{},
		signalSubscribers: map[SignalId][]EffectId{},
		effects:           []func(){},
	}
}

func runEffect(rt *Runtime, id EffectId) {
	previous := rt.runningEffect
	rt.runningEffect = &id
	rt.effects[id]()
	rt.runningEffect = previous
}
