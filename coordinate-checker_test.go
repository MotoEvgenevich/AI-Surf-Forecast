package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCoordinateChecker(t *testing.T) {
	// Пример использования функции
	coords := Coordinate{
		Latitude: 42.13076, Longitude: 41.65917,
	}
	fmt.Println(CoordinateChecker(coords))
	{
		result, err := CoordinateChecker(coords)
		assert.NotNil(t, result)
		assert.NoError(t, err)

	}

}
