package main

// Set is a collection of unique elements
type Set struct {
	elements map[string]struct{}
}

// NewSet creates a new set
func MakeSet() *Set {
	return &Set{
		elements: make(map[string]struct{}),
	}
}

// NewSet creates a new set
func ToSet(elements []string) *Set {
	set := MakeSet()

	for _, element := range elements {
		set.Add(element)
	}

	return set
}

// Add inserts an element into the set
func (s *Set) Add(value string) {
	s.elements[value] = struct{}{}
}

func (s *Set) Append(values []string) {
	for _, value := range values {
		s.Add(value)
	}
}

// Remove deletes an element from the set
func (s *Set) Remove(value string) {
	delete(s.elements, value)
}

// Contains checks if an element is in the set
func (s *Set) Contains(value string) bool {
	_, found := s.elements[value]
	return found
}

// Size returns the number of elements in the set
func (s *Set) Size() int {
	return len(s.elements)
}

// List returns all elements in the set as a slice
func (s *Set) List() []string {
	keys := make([]string, 0, len(s.elements))
	for key := range s.elements {
		keys = append(keys, key)
	}
	return keys
}

func (s *Set) Union(other *Set) *Set {
	result := MakeSet()
	for key := range s.elements {
		result.Add(key)
	}
	for key := range other.elements {
		result.Add(key)
	}
	return result
}

func (s *Set) Intersection(other *Set) *Set {
	result := MakeSet()
	for key := range s.elements {
		if other.Contains(key) {
			result.Add(key)
		}
	}
	return result
}

func (s *Set) Difference(other *Set) *Set {
	result := MakeSet()
	for key := range s.elements {
		if !other.Contains(key) {
			result.Add(key)
		}
	}
	return result
}
