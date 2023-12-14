package linkedlists

/*
Given a circular linked list, implement an algorithm that returns the node at the beginning of the loop.

Definition:
Circular linked list: A (corrupt) linked list in which a node’s next pointer points to an earlier node, so as to make a loop in the linked list.

Example:
Input: A -> B -> C -> D -> E -> C (the same C as earlier)
Output: C

# Linked List Loop Detection

![[Pasted image 20231128212148.png]]

With regards to the problem where one has to find a cycle in a linked list, the thought is that your fast runner is already "k" nodes inside the loop when the slow runner first enters the loop. The fast runner is also "k" nodes ahead of the slow runner because of its speed: it moves 2 nodes each step for the 1 node that the slow runner moves.

![[Pasted image 20231128212317.png]]

As the fast runner is moving twice as fast as the slow runner, it will eventually catch up to the slow runner, and that happens to be k nodes from the beginning of the loop.

So I it does end up that the slow and fast runners collide K nodes away from the entrance to the
Loop. The idea is that the fast runner will reach the slow runner in the number of steps equaling the  loop size minus the number of modes from the head to the loop entrance because the fast runner is already k nodes in the loop. Then, that means that the runners will meet k nodes away from the loop entrance within the loop.

![[Pasted image 20231128213155.png]]

Once we have the collision point, so we reset the slow pointer at the head and move both runners at the slow runner rate. When the runners collide again, we have the loop entrance.

![[Pasted image 20231128215200.png]]


![[Pasted image 20231128220215.png]]

#LinkedLists
#DataStructuresAndAlgorithms

*/

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
