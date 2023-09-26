package reactive

type EffectId int

func NewEffect(f func()) {
	rt := GetRuntime()
	effectId := EffectId(len(rt.effects))
	rt.effects = append(rt.effects, f)

	runEffect(rt, effectId)
}

func NewEffectWithRuntime(rt *Runtime, f func()) {
	effectId := EffectId(len(rt.effects))
	rt.effects = append(rt.effects, f)

	runEffect(rt, effectId)
}
