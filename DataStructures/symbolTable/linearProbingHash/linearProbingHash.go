package linearProbingHash

type Key *string
type Value *int

type st struct {
        len  int
        cap  int
        keys []Key
        vals []Value
}

func New(cap int) *st {
        s := new(st)
        s.len = 0
        s.cap = cap
        s.keys = make([]Key, s.cap)
        s.vals = make([]Value, s.cap)
        return s
}

func (s st) Size() int {
        return s.len
}

func (s st) IsEmpty() bool {
        return s.Size() == 0
}

func (s st) Contains(key Key) bool {
        return s.Get(key) != nil
}

func (s st) Get(key Key) Value {
        for i := s.hash(key); s.keys[i] != nil; i = (i+1)%s.cap {
                if key == s.keys[i] {
                        return s.vals[i]
                }
        }
        return nil
}

func (s st) hash(key Key) int {
        return (s.hashCode(key) & 0x7fffffff) %s.cap
}

func (s st) hashCode(key Key) int {
        ks := *key
        l := len(ks)
        hash := 0
        for i := 0; i < l; i++ {
                hash = 31 * hash + int(ks[i])
        }
        return hash
}

func (s *st) resize(cap int) {
        temp := New(cap)
        for i := 0; i < s.cap; i++ {
                if s.keys[i] != nil {
                        temp.Put(s.keys[i], s.vals[i])
                }
        }
        s.keys = temp.keys
        s.vals = temp.vals
        s.cap = temp.cap
}

func (s *st) Put(key Key, value Value)  {
        if value == nil {
                s.Delete(key)
                return
        }

        if s.len > s.cap/2 {
                s.resize(s.cap*2)
        }

        i := s.hash(key)
        for ; s.keys[i] != nil; i = (i+1)%s.cap {
                if key == s.keys[i] {
                        s.vals[i] = value
                        return
                }
        }
        s.keys[i] = key
        s.vals[i] = value
        s.len++
}

func (s *st) Delete(key Key)  {
        if !s.Contains(key) {
                return
        }

        i := s.hash(key)
        for key != s.keys[i] {
                i = (i+1)%s.cap
        }

        s.keys[i] = nil
        s.vals[i] = nil

        i = (i+1)%s.cap
        for s.keys[i] != nil {
                newKey := s.keys[i]
                newVal := s.vals[i]
                s.keys[i] = nil
                s.vals[i] = nil
                s.len--
                s.Put(newKey, newVal)
                i = (i+1)%s.cap
        }

        s.len--

        if s.len > 0 && s.len <= s.cap/8 {
                s.resize(s.cap/2)
        }
}