package linkedlists

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAnimalShelterEnqueueCats(t *testing.T) {
	s := AnimalShelter{}
	s.enqueue(Cat{})
	s.enqueue(Cat{})
}
func TestAnimalShelterDequeueCats(t *testing.T) {
	s := AnimalShelter{}
	s.enqueue(Cat{})
	s.enqueue(Cat{})
	s.enqueue(Cat{})

	for i := 0; i < 3; i++ {
		_, ok := s.dequeueCat()
		assert.True(t, ok)
	}
	_, ok := s.dequeueCat()
	assert.False(t, ok)
}

// I won't bother with a test dequeue dogs test as it
// is the exact same logic as dequeue cats.

func TestAnimalShelterEnqueueAnimals(t *testing.T) {
	s := AnimalShelter{}
	s.enqueue(Cat{})
	s.enqueue(Dog{})
	s.enqueue(Dog{})
	s.enqueue(Dog{})
	s.enqueue(Cat{})
	s.enqueue(Cat{})
	s.enqueue(Dog{})
	s.enqueue(Cat{})
	s.enqueue(Cat{})
}
func TestAnimalShelterDequeueAnimals(t *testing.T) {
	s := AnimalShelter{}
	s.enqueue(Cat{})
	s.enqueue(Dog{})
	s.enqueue(Dog{})
	s.enqueue(Dog{})
	s.enqueue(Cat{})
	s.enqueue(Cat{})
	s.enqueue(Dog{})
	s.enqueue(Cat{})
	s.enqueue(Cat{})

	expectedAnimals := []Animal{
		Cat{},
		Dog{},
		Dog{},
		Dog{},
		Cat{},
		Cat{},
		Dog{},
		Cat{},
		Cat{},
	}

	for _, expectedAnimal := range expectedAnimals {
		actualAnimal, _ := s.dequeueAny()
		assert.IsType(t, expectedAnimal, actualAnimal)
	}
}
