package linkedlists

func summedNodeWithCarry[T int32](v1 T, v2 T, carry T) (T, *Node[T]) {
	sum := v1 + v2 + carry
	if sum < T(10) {
		return T(0), CreateNode(sum)
	}
	return sum / T(10), CreateNode(sum % T(10))
}

func sumListsWithCarry[T int32](list1 *Node[T], list2 *Node[T]) (T, *Node[T]) {
	if list1 == nil || list2 == nil {
		return 0, nil
	}
	if list1.next == nil && list2.next == nil {
		return summedNodeWithCarry(list1.data, list2.data, 0)
	}
	carry, summedChildList := sumListsWithCarry(list1.next, list2.next)
	carry, summedList := summedNodeWithCarry(list1.data, list2.data, carry)
	AppendNode(summedList, summedChildList)
	return carry, summedList
}

func sumListsInOnePass[T int32](list1 *Node[T], list2 *Node[T]) *Node[T] {
	carry, summedList := sumListsWithCarry(list1, list2)
	if carry != 0 {
		head := CreateNode(carry)
		head.next = summedList
		return head
	}
	return summedList
}
