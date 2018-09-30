package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsInitialAdjMatrixCorrect(t *testing.T) {
	graph := NewGraph(6)
	actualNodeNumber := graph.GetNodeNumber()
	assert.Equal(t, 6, actualNodeNumber)
	for i := 0; i < 6; i++ {
		for j := 0; j < 6; j++ {
			actualEdgeWeight, _ := graph.GetEdgeWeight(i, j)
			assert.Equal(t, -1, actualEdgeWeight)
		}
	}
}

func TestSettingEdgeWeight(t *testing.T) {
	graph := NewGraph(6)
	graph.SetEdgeWeight(2, 3, 11)
	weight, _ := graph.GetEdgeWeight(2, 3)
	assert.Equal(t, 11, weight)
	weight, _ = graph.GetEdgeWeight(3, 2)
	assert.Equal(t, 11, weight)
}

func TestSettingEdgeWeightError(t *testing.T) {
	graph := NewGraph(6)
	actualError := graph.SetEdgeWeight(2, 6, 11)
	assert.Error(t, actualError)
	actualError = graph.SetEdgeWeight(5, 5, 11)
	assert.Error(t, actualError)
	actualNoError := graph.SetEdgeWeight(0, 5, 11)
	assert.NoError(t, actualNoError)
	largeArgError := graph.SetEdgeWeight(0, 6, 11)
	sameArgError := graph.SetEdgeWeight(5, 5, 11)
	assert.EqualError(t, largeArgError, "out of range error")
	assert.EqualError(t, sameArgError, "same argument error")
}

func TestGettingEdgeWeightError(t *testing.T) {
	graph := NewGraph(6)
	_, largeArgError := graph.GetEdgeWeight(0, 6)
	_, sameArgError := graph.GetEdgeWeight(5, 5)
	weight, noError := graph.GetEdgeWeight(0, 1)
	assert.EqualError(t, largeArgError, "out of range error")
	assert.EqualError(t, sameArgError, "same argument error")
	assert.Equal(t, -1, weight)
	assert.NoError(t, noError)
}
