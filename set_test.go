package set

import (
  "testing"
)

func Test_New(t *testing.T) {
  s := New()
  if s.Size() != 0 {
    t.Error("New should return an empty set")
  }
}

func Test_New_Prefill(t *testing.T) {
  s := New(1, "123", 5i+1)
  if s.Size() != 3 {
    t.Error("Created Set should contain 3 items")
  }
}

func Test_Size(t *testing.T) {
  s := New(1,2,3)
  if s.Size() != 3 {
    t.Error("Size should report 3 items")
  }
}

func Test_Add(t *testing.T) {
  s := New()
  result := s.Add(1)
  if !result {
    t.Error("No addition reported while expected")
  }

  if s.Size() != 1 {
    t.Error("No item was added")
  }
}

func Test_Add_Twice(t *testing.T) {
  s := New()
  s.Add(1)
  result := s.Add(1)
  if result {
    t.Error("Adding same element twice should return false since already exists")
  }
}

func Test_Remove(t *testing.T) {
  s := New(1,2,3)
  s.Remove(3)
  if s.Size() != 2 {
    t.Error("Removing did not change size")
  }
}

func Test_Includes(t *testing.T) {
  s := New(1,2,3)
  if !s.Includes(1) {
    t.Error("Includes reported missing element while it is there")
  }

  if s.Includes(4) {
    t.Error("Includes reported included element while it is not there")
  }
}

/* Set operations */
func Test_Union(t *testing.T) {
  s1 := New(1,2,3)
  s2 := New(3,4,5)
  u := s1.Union(s2)
  if u.Size() != 5 {
    t.Error("Union resulted in wrong size")
  }
}

func Test_Intersect(t *testing.T) {
  s1 := New(1,2,3)
  s2 := New(3,4,5)
  i := s1.Intersect(s2)
  if i.Size() != 1 {
    t.Error("Intersect resulted in wrong size")
  }

  if !i.Includes(3) {
    t.Error("Intersect does not include only intended element")
  }
}

// General - higher level - tests
// for i.e. multiple types of items.
func Test_Multiple_Types_Size(t *testing.T) {
  s := New()
  s.Add(1)
  s.Add("123")
  if s.Size() != 2 {
    t.Error("Size reported incorrect size")
  }
}

func Test_Multiple_Types_Includes(t *testing.T) {
  s := New()
  s.Add(1)
  s.Add("123")
  if !s.Includes(1) || !s.Includes("123") {
    t.Error("One of the included items is reported not to be included")
  }
}

func Test_String(t *testing.T) {
  s := New(1,2,3)
  str := s.String()
  if str != "Set(size: 3)" {
    t.Error("String representation is incorrect\n\tIs: " + str + "\n\tExpected: Set(size: 3)")
  }
}
