package linkedlists

import (
	"fmt"
	"io"
	"os"
)

/*
Problem: Sum Lists

	You have two numbers represented by a linked list, where each node contains a single digit.
	The digits are stored in reverse order, such that the 1's digit is at the head of the list.
	Write a function that adds the two numbers and returns the sum as a linked list.

	EXAMPLE:

	input: ([7] -> [1] -> [6]) + ([5] -> [9] -> [2]). That is, 617 + 295.
	output:  [2] -> [1] -> [9]. That is 912.

	FOLLOW UP:
	Suppose that the digits are stored in "forward" order, repeat the above problem:

	input: ([6] -> [1] -> [7]) + ([2] -> [9] -> [5]). That is, 617 + 295.
	output: [9] -> [1] -> [2]. That is 912.
*/

/*
My first inclination was to write a method that would take a list and convert it to a number.
I could then add those two numbers to get the value.
Then, I would take that value, and pass it to a method that would convert it to a list.

I could write it fairly straightforward for "reverse" order, where the list starts with the
1's digit.
*/

/*
First, I need to visualize out to use recursion to print out a number, digit by digit
using recursion

This method lets me use an io.Writer to either use stdout, or in a unit test, my
own io.Writer so that I can compare the values.
*/
func printNumberByDigitsReversed(x int32, out io.Writer) {
	if out == nil {
		out = os.Stdout
	}
	digits := []byte("0123456789")
	if x < 10 {
		fmt.Fprintf(out, "%c", digits[x])
		return
	}
	fmt.Fprintf(out, "%c", digits[x%10])
	printNumberByDigitsReversed(x/10, out)
}

/*
Now that I have a concept of how to construct a string with the 1's digits
at the beginning, I can take this understanding to construct a linked list
of a number with the 1's digits at the beginning.

This will be used after the number value of a reversed list is retrieved from
both lists and added in this method.
*/

func numberToReverseList[T int32](i T) *Node[T] {
	if i < 10 {
		// this is just like the condition x < 10 above where we "print" out the node as is.
		return CreateNode(i)
	}
	list := CreateNode(i % 10)              // this is like the condition where we print out the value
	list.next = numberToReverseList(i / 10) // and we continue printing the remainder of the list

	return list
}

/*
Now that I have a method to convert any number to a reversed list with the 1's digits
at the head, I create a method to take a reversed list, with 1's digits at the head, and
convert that to the corresponding number.

Follow up: never trust a programmer to do the right thing. We were told that we would get
a single digit for each node, but let's ensure that we do.
*/
func reverseListToNumber[T int32](list *Node[T]) (T, error) {
	var number T
	for i, cursor := T(1), list; cursor != nil; i, cursor = i*10, cursor.next {
		if cursor.data > 9 {
			return 0, fmt.Errorf("error: nodes must contain only single digits! [%d]", cursor.data)
		}
		number += cursor.data * i
	}
	return number, nil
}

/*
Now that I have a method to convert a reversed list of digits where the 1's digits are at the
head of the list, and I have a method to convert a number to a reversed list of digits in that
same order, I can get the value of each list and create a list from the sum of those values

Note: since there is a possibility that we'll get an invalid node of more than one digit,
we have to let the caller know that they have to pay attention for that in the return value
*/
func sumReverseLists[T int32](list1 *Node[T], list2 *Node[T]) (*Node[T], error) {
	number1, err1 := reverseListToNumber(list1)
	if err1 != nil {
		return nil, err1
	}
	number2, err2 := reverseListToNumber(list2)
	if err2 != nil {
		return nil, err2
	}
	return numberToReverseList(number1 + number2), nil
}

/*
Using the same methodology, let's create functions that take a number and convert it to a list
with the 1's digits as the last node and a function that takes such as a list and converts it to
a number. We'll use that to make a method that sums two lists in that order.

But first, let's use the same method to help us visualize that by printing to a string in the
correct order. That will help us determine when to print, and when to dig down into the list.
*/

func printNumberByDigits(x int32, out io.Writer) {
	if out == nil {
		out = os.Stdout
	}
	digits := []byte("0123456789")
	if x < 10 {
		fmt.Fprintf(out, "%c", []any{digits[x]}...)
		return
	}
	printNumberByDigits(x/10, out)
	fmt.Fprintf(out, "%c", []any{digits[x%10]}...)
}

/*
Now that we have a visual concept of how to create the list of characters in a string, we can
use the same model to create a linked list from a number, where the 1's digit is at the tail.

Imagine as if your "out" above, is the linked list you are creating here.
*/
func numberToList[T int32](i T) *Node[T] {
	if i < 10 {
		// this is like the printNumberByDigits step where we check if x < 10 and append the node to "out"
		// our, in our case, our list will append this node to it's "out"
		return CreateNode(i)
	}
	list := numberToList(i / 10)       // as above, we recurse into ourselves with x/10, or set "out" to that list
	AppendNode(list, CreateNode(i%10)) // and as above, we append a new node with the remainder to our list

	return list
}

/*
To wrap up, we need a method that converts a list, where the 1's digits are at the end of the list to
an actual number. We should also account for any craziness where the programmer gave us a multidigit
node.
*/
func listToNumber[T int32](list *Node[T]) (T, error) {
	var number T
	for cursor := list; cursor != nil; cursor = cursor.next {
		if cursor.data > 9 {
			return 0, fmt.Errorf("error: nodes must contain only single digits! [%d]", cursor.data)
		}
		number = number * T(10)
		number += cursor.data
	}

	return number, nil
}

/*
Now that we have our methods that convert a number into a list with the 1's digits at the end, and a
list with 1's digits at the end to a number, we can create the method that adds two lists following that
same digit order.
*/
func sumLists[T int32](list1 *Node[T], list2 *Node[T]) (*Node[T], error) {
	number1, err1 := listToNumber(list1)
	if err1 != nil {
		return nil, err1
	}
	number2, err2 := listToNumber(list2)
	if err2 != nil {
		return nil, err1
	}
	return numberToList(number1 + number2), nil
}

/*
Okay, so we have convenient methods to sum forward and reverse lists.
However, this has a runtime, for both the forward and reverse lists, of:

1. O(n) for list 1
2. O(n) for list 2
3. O(n) to create list 3

That's 3N

Can we do better? Yes, we can add each node, node by node. That would give us a runtime of O(n).
The only thing is that we have to concern ourselves with the carry.
That is to say, if we have two lists [9] and [7], that should create lists:

reverse order: (9) + (7) = [6] -> [1] // 16
forward order: (9) + (7) = [1] -> [6] // 16

For now, let's assume that each list is the same number of digits
*/

/*
I need to iterate further and further down the list until I reach the end. Once I reach the end
I need to sum the two values. If the value is greater than two digits, I need to provide the
"carry" value (to be used to the next two digits, or to create a new node at the front of the list)

We should also confirm that we get an error if we dare to send a multidigit number in as v1 or v2.

*/

/*
I think summedNodeWithCarry by itself is fine, I need to figure out how to
follow one list down when another list has ended to set the pointer in the
correct position
*/
func summedNodeWithCarry[T int32](v1 T, v2 T, carry T) (T, *Node[T], error) {
	if v1 > 9 || v2 > 9 || carry > 9 {
		biggestNumber := max(v1, v2, carry)
		return 0, nil, fmt.Errorf("error: nodes and the carry value must be only single digits! [%d]", biggestNumber)
	}
	sum := v1 + v2 + carry
	if sum < T(10) {
		return T(0), CreateNode(sum), nil
	}
	return sum / T(10), CreateNode(sum % T(10)), nil
}

func sumForwardListsWithCarry[T int32](list1 *Node[T], list2 *Node[T]) (T, *Node[T], error) {
	if list1 == nil || list2 == nil {
		return 0, nil, nil
	}
	if list1.next == nil && list2.next == nil {
		return summedNodeWithCarry(list1.data, list2.data, 0)
	}
	carry, summedChildForwardList, err := sumForwardListsWithCarry(list1.next, list2.next)
	if err != nil {
		return 0, nil, err
	}
	carry, summedForwardList, err := summedNodeWithCarry(list1.data, list2.data, carry)
	if err != nil {
		return 0, nil, err
	}
	AppendNode(summedForwardList, summedChildForwardList)
	return carry, summedForwardList, nil
}

func sumForwardListsInOnePass[T int32](list1 *Node[T], list2 *Node[T]) (*Node[T], error) {
	carry, summedForwardList, err := sumForwardListsWithCarry(list1, list2)
	if err != nil {
		return nil, err
	}
	if carry != 0 {
		head := CreateNode(carry)
		head.next = summedForwardList
		return head, nil
	}
	return summedForwardList, nil
}

/*
Now, we have it working for lists that have the same number of nodes. However, we are assuming that
we have the same number of digits in each list. What if we have a smaller
list for one of the parameters. It should fail right now, let's add a test to demonstrate that
failure

9 -> 7
9 -> 9 -> 4

This is more complicated, this should really be lists like so:

0 -> 9 -> 7
9 -> 9 -> 4

But, that is not what I am getting.

9 -> 7
^

9 -> 9 -> 4
^

Can we do some trick with the carry?

# Let's say we are at the last digit for one of the lists

9 -> 7

	^

9 -> 9 -> 4

	^

We see that one of the lists ends. Can we add the 7 and 4 and return a big
enough carry that it propogates up? It may not be that easy, because
we could have a list like this:

9 -> 7

	^

9 -> 9 -> 4 -> 5

	^

Perhaps instead, when we reach the end of one list, we need to continue the
list down, and do somethihg fancy when we come back up.

9 -> 7

	^

9 -> 9 -> 4 -> 5

	^

Then, we add 7 and 5. Then, we check the depth as we come back up whether or
not to come back up on both lists, or just one (adding the carry).

Okay, what if we had two stacks, we push all the elements for one list to one stack
and all the other elements for the other list to the other stack.

l1: 9 -> 7 (97)

l2: 9 -> 9 -> 4 -> 5 (9945)

summed list: 1 -> 0 -> 0 -> 4 -> 2 (10042)

Then, we pop each element from each stack and add them together, carrying the "carry". If one stack
is empty, we simply add the left over carry to the values of the non-empty popped element

s1:

	7
	9

s2:

	5
	4
	9
	9

iteration 1:

0 (carry) + 7 + 5 = 12 => 1, [2] // 1 carry, a node with value 2
1 + 4 + 9 = 14 => 1, [4] -> [2]; // 1 carry, add the 4 node in front of the 2 node from last iteration
1 + 9 + 0 (nil item) = 1, [0] -> [4] -> [2]; // 1 carry, and add the [0] node in front of the list from last iteration
1 + 9 + 0 (nil item) = 1, [0] -> [0] -> [4] -> [2]; // 1 carry, and add the [0] node in front of the list from last iteration
// at this point, no more elements in either stack, but there is a 1 carry, let's add that as a node in front of the list from the last iteration

[1] -> [0] -> [0] -> [4] -> [2]
*/
func _stackTwoForwardLists[T int32](list1 *Node[T], list2 *Node[T], stack1 *Stack[*Node[T]], stack2 *Stack[*Node[T]]) error {

	for c1, c2 := list1, list2; /* keep going while one of these are != nil */ c1 != nil || c2 != nil; {
		if c1 != nil {
			if c1.data > 9 {
				return fmt.Errorf("error: nodes must contain only single digits! [%d]", c1.data)
			}
			(*stack1).Push(c1)
			c1 = c1.next
		}
		if c2 != nil {
			if c2.data > 9 {
				return fmt.Errorf("error: nodes must contain only single digits! [%d]", c2.data)
			}
			(*stack2).Push(c2)
			c2 = c2.next
		}
	}
	return nil
}

func sumUnevenForwardLists[T int32](list1 *Node[T], list2 *Node[T]) (*Node[T], error) {
	if list2 == nil && list1 == nil {
		return nil, nil
	} else if list1 == nil {
		return list2, nil
	} else if list2 == nil {
		return list1, nil
	}

	var stack1 Stack[*Node[T]]
	var stack2 Stack[*Node[T]]
	err := _stackTwoForwardLists(list1, list2, &stack1, &stack2)
	if err != nil {
		return nil, err
	}
	// Add list elements to their respective stacks
	// Add the elements at the top of the stack, with a potential carry
	// If one stack no longer has elements, use the value of 0.
	carry := T(0)
	var head *Node[T]
	var node *Node[T]

	top1, ok1 := stack1.Pop()
	top2, ok2 := stack2.Pop()
	for top1 != nil || top2 != nil {
		v1, v2 := T(0), T(0)
		if ok1 {
			v1 = top1.data
			top1, ok1 = stack1.Pop()
		}
		if ok2 {
			v2 = top2.data
			top2, ok2 = stack2.Pop()
		}
		carry, node, _ = summedNodeWithCarry(v1, v2, carry)
		if head == nil {
			head = node
		} else {
			node.next = head
			head = node
		}
	}
	if carry != 0 {
		node := CreateNode(carry)
		node.next = head
		head = node
	}
	return head, nil
}
