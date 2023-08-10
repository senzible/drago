package main

import (
	"fmt"
	"syscall/js"
)

type Runtime struct {
	//slice of signal values
	signalValues      []any
	runningEffect     *EffectId
	signalSubscribers map[SignalId][]EffectId
	effects           []func()
}

func (rt *Runtime) CreateSignal(initial any) Signal {
	id := SignalId(len(rt.signalValues))
	rt.signalValues = append(rt.signalValues, initial)
	return Signal{rt, id}
}

func (rt *Runtime) CreateEffect(f func()) {
	effectId := EffectId(len(rt.effects))
	rt.effects = append(rt.effects, f)

	rt.RunEffect(effectId)
}

func (rt *Runtime) RunEffect(id EffectId) {
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

type Signal struct {
	*Runtime
	SignalId
}

func (s *Signal) Get() any {
	//get the value
	val := s.Runtime.signalValues[s.SignalId]

	//if there is a running effect, add it to the subscribers
	//only if it is not already subscribed
	if rt := s.Runtime; rt.runningEffect != nil {
		if _, ok := rt.signalSubscribers[s.SignalId]; !ok {
			//append the effect to the subscribers
			rt.signalSubscribers[s.SignalId] = append(rt.signalSubscribers[s.SignalId], *rt.runningEffect)
		}
	}

	//return value
	return val
}

func (s *Signal) Set(v any) {
	//set the value
	s.Runtime.signalValues[s.SignalId] = v

	//notify subscribers
	for _, effectId := range s.Runtime.signalSubscribers[s.SignalId] {
		s.Runtime.RunEffect(effectId)
	}
}

func main() {
	c := make(chan struct{})

	window := js.Global()
	doc := window.Get("document")
	body := doc.Get("body")

	p := doc.Call("createElement", "p")
	p.Set("textContent", "Hello, WebAssembly! Not reactive yet!")

	inc := doc.Call("createElement", "button")
	inc.Set("textContent", "+1")
	dec := doc.Call("createElement", "button")
	dec.Set("textContent", "-1")

	body.Call("appendChild", inc)
	body.Call("appendChild", p)
	body.Call("appendChild", dec)

	rt := Runtime{
		signalValues:      []any{},
		signalSubscribers: map[SignalId][]EffectId{},
		effects:           []func(){},
	}
	count := rt.CreateSignal(0)

	rt.CreateEffect(func() {
		p.Set("textContent", fmt.Sprintf("Count: %d", count.Get()))
	})

	//add event listeners
	inc.Call("addEventListener", "click", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		count.Set(count.Get().(int) + 1)
		return nil
	}))

	dec.Call("addEventListener", "click", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		count.Set(count.Get().(int) - 1)
		return nil
	}))

	<-c
}
