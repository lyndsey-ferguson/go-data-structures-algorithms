package linkedlists

func findCircularListHeadWithHash[T comparable](list *Node[T]) *Node[T] {
	visitedNodes := make(map[*Node[T]]bool)

	for cursor := list; cursor != nil; cursor = cursor.next {
		_, ok := visitedNodes[cursor]
		if ok {
			return cursor
		}
		visitedNodes[cursor] = true
	}
	return nil
}

func findCircularListHeadWithRunners[T comparable](list *Node[T]) *Node[T] {

	if list == nil || list.next == nil {
		return nil
	}

	fastRunner := list
	slowRunner := list

	// find node where the runners collide
	for fastRunner != nil && fastRunner.next != nil {
		fastRunner = fastRunner.next.next
		slowRunner = slowRunner.next

		if slowRunner == fastRunner {
			break
		}
	}

	// exit early if the fastRunner made it to the end of the list
	if fastRunner == nil || fastRunner.next == nil {
		return nil
	}

	// now, the fastRunner and slowRunner have collided, and
	// according to their natures where fastRunner moves at
	// 2 nodes per iteration while slowRunner moves at 1 node
	// per iteration, AND fastRunner had a head start of "k"
	// nodes into the loop, where k is the number of nodes
	// for slowRunner to reach the loop, fastRunner will
	// have caught up to slowRunner in loopsize - k steps
	// leaving both k nodes before the beginning of the loop
	// (because fastRunner started k nodes into the loop, and
	// there will be loopsize - k nodes left). This allows
	// us to reset slowRunner to the head of the list
	// and then step each runner at the speed of 1 node per
	// iteration until they collide again
	slowRunner = list
	for ; slowRunner != fastRunner; slowRunner, fastRunner = slowRunner.next, fastRunner.next {
	}

	return slowRunner
}
