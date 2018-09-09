package ev

import "reflect"

type Args struct {
	s []interface{}
}

func (a *Args) Len() int         { return len(a.s) }
func (a *Args) S() []interface{} { return a.s }

func (a *Args) String(i uint) string { return a.s[i].(string) }
func (a *Args) Int(i uint) int       { return a.s[i].(int) }
func (a *Args) Bool(i uint) bool     { return a.s[i].(bool) }

type handlerInfo struct {
	ptr  uintptr
	f    func(Args)
	once bool
}

type Ent struct {
	handlers []handlerInfo
}

func (e *Ent) On(h func(Args)) {
	e.handlers = append(e.handlers, handlerInfo{
		ptr:  reflect.ValueOf(h).Pointer(),
		f:    h,
		once: false,
	})
}

func (e *Ent) Off(h func(Args)) {
	hptr := reflect.ValueOf(h).Pointer()
	for i, hInfo := range e.handlers {
		if hInfo.ptr == hptr {
			e.handlers = append(e.handlers[:i], e.handlers[i+1:]...)
			return
		}
	}
	panic("handler not found")
}

func (e *Ent) LockOn(h func(Args)) {
	e.handlers = append(e.handlers, handlerInfo{
		f:    h,
		once: false,
	})
}

func (e *Ent) Once(h func(Args)) {
	e.handlers = append(e.handlers, handlerInfo{
		f:    h,
		once: true,
	})
}

func (e *Ent) Invoke(args ...interface{}) {
	a := Args{s: args}
	for i, hInfo := range e.handlers {
		hInfo.f(a)
		if hInfo.once {
			e.handlers = append(e.handlers[:i], e.handlers[i+1:]...)
		}
	}
}
