package types

type Context struct {
	Store map[string]ObjectJIPL
	Outer *Context // the outer scope
}

func NewContext() *Context {
	return &Context{
		Store: make(map[string]ObjectJIPL),
		Outer: nil,
	}
}

func NewContextWithOuter(outer *Context) *Context {
	ctx := NewContext()
	ctx.Outer = outer
	return ctx
}

func (ctx *Context) Get(key string) (ObjectJIPL, bool) {
	val, ok := ctx.Store[key]
	if !ok && ctx.Outer != nil {
		// recursively search for the key
		// in the outer context
		return ctx.Outer.Get(key)
	}
	return val, ok
}

func (ctx *Context) Set(key string, val ObjectJIPL) ObjectJIPL {
	ctx.Store[key] = val
	return val
}
