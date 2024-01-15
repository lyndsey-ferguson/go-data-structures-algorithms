package linkedlists

import (
	"fmt"
	"math"
	"time"
)

/*
An animal shelter, which holds only dogs and cats, operates on a structly
"first in, first out" basis. People must adopt either the "oldest" (based on
arrival time) of all animals at the shelter, or they can select whether they
would prefer a dog or a cat (and will receive the oldest animal of that type).

They cannot select which specific animal they would like.
Create the data structures to maintain this system, and implement operations
such as enqueue, dequeueAny, dequeueCat, and dequeueDog. You may use the
built-in LinkedList data structure.

At first, I thought about creating a normal queue, and then looking for the first
cat or the first dog for the dequeueCat and dequeueDog methods. But that would
be a lot of searching for the first dog or first cat, and would be wasteful
because of 1) duplicate work, and 2) wasted work if there are no animals of the
desired type.

It was better to create two queues: one for cats, and one for dogs. However, we
would need to know which animal from each queue was brought in first in order
to satisfy the method dequeueAny. Each animal is brought in at a separate time
so let's store a timestamp for each animal and when dequeueAny is called
get the smallest timestamp for each queue (setting a default to max in case there
are no animals of a given type), and then return the animal that has been in
the shelter the longest time.
*/

type Animal interface {
	Speak()
}

type Cat struct{}
type Dog struct{}

func (c Cat) Speak() {
	fmt.Println("Meow")
}
func (d Dog) Speak() {
	fmt.Println("Ruff")
}

type TimestampedNode[T comparable] struct {
	data               T
	next               *TimestampedNode[T]
	timestampSinceEpoc int64
}

func CreateTimestampedNode[T comparable](value T, timestamp int64) *TimestampedNode[T] {
	if timestamp == 0 {
		timestamp = time.Now().UnixNano()
	}
	n := TimestampedNode[T]{
		data:               value,
		timestampSinceEpoc: timestamp,
	}
	return &n
}

type AnimalShelter struct {
	creationTimestamp int64 // added to mock the time between animals sheltered
	firstCat          *TimestampedNode[Cat]
	lastCat           *TimestampedNode[Cat]
	firstDog          *TimestampedNode[Dog]
	lastDog           *TimestampedNode[Dog]
}

func (s *AnimalShelter) enqueue(a Animal) {
	s.creationTimestamp += 1
	if c, ok := a.(Cat); ok {
		s._enqueueCat(c)
	} else if d, ok := a.(Dog); ok {
		s._enqueueDog(d)
	}
}
func (s *AnimalShelter) _enqueueCat(c Cat) {
	n := CreateTimestampedNode(c, s.creationTimestamp)
	if s.lastCat != nil {
		s.lastCat.next = n
	}
	s.lastCat = n
	if s.firstCat == nil {
		s.firstCat = s.lastCat
	}
}
func (s *AnimalShelter) _enqueueDog(c Dog) {
	n := CreateTimestampedNode(c, s.creationTimestamp)
	if s.lastDog != nil {
		s.lastDog.next = n
	}
	s.lastDog = n
	if s.firstDog == nil {
		s.firstDog = s.lastDog
	}
}
func (s *AnimalShelter) dequeueCat() (Cat, bool) {
	if s.firstCat == nil {
		var nonExistentCat Cat
		return nonExistentCat, false
	}
	c := s.firstCat.data
	if s.lastCat == s.firstCat {
		s.lastCat = nil
		s.firstCat = nil
	} else {
		s.firstCat = s.firstCat.next
	}
	return c, true
}
func (s *AnimalShelter) dequeueDog() (Dog, bool) {
	if s.firstDog == nil {
		var nonExistentDog Dog
		return nonExistentDog, false
	}
	c := s.firstDog.data
	if s.lastDog == s.firstDog {
		s.lastDog = nil
		s.firstDog = nil
	} else {
		s.firstDog = s.firstDog.next
	}
	return c, true
}
func (s *AnimalShelter) dequeueAny() (Animal, bool) {
	if s.firstCat == nil && s.firstDog == nil {
		var nonExistentAnimal Animal
		return nonExistentAnimal, false
	}
	firstCatTimestamp := int64(math.MaxInt64)
	firstDogTimestamp := int64(math.MaxInt64)
	if s.firstCat != nil {
		firstCatTimestamp = s.firstCat.timestampSinceEpoc
	}
	if s.firstDog != nil {
		firstDogTimestamp = s.firstDog.timestampSinceEpoc
	}
	if firstCatTimestamp < firstDogTimestamp {
		return s.dequeueCat()
	}
	return s.dequeueDog()
}
