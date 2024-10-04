package trees

import (
	"math"
)

func findClosestValue(root *Node[int], target float64) int {
	node := root
	closest := int(node.data)

	for node != nil {
		value := node.data
		if math.Abs(float64(value)-target) < math.Abs(float64(closest)-target) {
			closest = int(node.data)
		}

		if target > float64(node.data) {
			node = node.right
		} else {
			node = node.left
		}
	}
	return closest
}
