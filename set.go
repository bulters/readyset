package set

import (
  "strconv"
)

type Set map[interface{}]struct{}

func New(payload ...interface{}) Set {
  switch {
  case payload == nil:
    return NewEmptySet()
  default:
    return NewFilledSet(payload)
  }
}

func NewEmptySet() Set {
  return make(Set)
}

func NewFilledSet(payload []interface{}) Set {
  s := NewEmptySet()
  for _, i := range(payload) {
    s.Add(i)
  }
  return s
}

func (s *Set) Size() int {
  return len(*s)
}

func (s Set) Add(i interface{}) bool {
  _, exists := s[i]
  s[i] = struct{}{}
  return !exists
}

func (s Set) Remove(i interface{}) bool {
  _, exists := s[i]
  delete(s, i)
  return !exists
}

func (s Set) Includes(i interface{}) bool {
  _, exists := s[i]
  return exists
}

func (s Set) Union(o Set) Set {
  u := make(Set)
  for i, _ := range(s) {
    u.Add(i)
  }
  for i, _ := range(o) {
    _, e := s[i]
    if !e {
      u.Add(i)
    }
  }

  return u
}

func (s Set)Intersect(o Set) Set {
  is := make(Set)
  for i, _ := range(s) {
    _, e := is[i]
    if !e {
      _, es := s[i]
      _, eo := o[i]
      if es && eo {
        is.Add(i)
      }
    }
  }
  return is
}

// Formatting functions
func (s Set) String() string {
  sizeStr := strconv.Itoa(s.Size())
  return "Set(size: " + sizeStr + ")"
}
