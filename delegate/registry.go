package delegate

import "io"

type Module map[string]Delegator

type Registry struct {
	module  Module
	tracing io.Writer
}

type RegistryOpt func(*Registry)

func RegistryWithTracing(w io.Writer) RegistryOpt {
	return func(r *Registry) {
		r.tracing = w
	}
}

func NewRegistry(opts ...RegistryOpt) Registry {
	var r Registry

	for _, opt := range opts {
		opt(&r)
	}

	return r
}

func (r *Registry) Import(mods ...Module) {
	for _, mod := range mods {
		for k, d := range mod {
			r.Register(k, d)
		}
	}
}

func (r *Registry) Register(k string, f Delegator) {
	if r.module == nil {
		r.module = make(Module)
	}
	if r.tracing != nil {
		f = WithTracing(k, f, r.tracing)
	}
	r.module[k] = f
}

func (r *Registry) Get(k string) (Delegator, bool) {
	d, ok := r.module[k]
	return d, ok
}

func Import(mods ...Module) Registry {
	var r Registry
	r.Import(mods...)
	return r
}
