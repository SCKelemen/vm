package obj

func NewScope(parent *Scope) *Scope {
	return Scope{store: make(map[string]Object), parent: parent}
}

type Scope struct {
	store  map[string]Object
	parent *Scope
}

func (s *Scope) Get(name string) (object Object, ok bool) {
	object, ok = s.store[name]
	if !ok && s.parent != nil {
		object, ok = s.parent.Get(name)
	}
	return
}

func (s *Scope) Set(name string, val Object) Object {
	s.store[name] = val
	return val
}
