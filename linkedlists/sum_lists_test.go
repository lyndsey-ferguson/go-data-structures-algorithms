package linkedlists

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
A test that helps me ensure that I am thinking correctly
about how to construct a reverse linked list with the
1's digit printed first -- or conceptually similar to
a linked list with the first element being the 1's digit.
*/
func TestPrintNumberByDigitsReversed(t *testing.T) {
	x := 912
	var b bytes.Buffer

	printNumberByDigitsReversed(int32(x), &b)

	assert.Equal(t, "219", b.String())
}

/*
This test makes sure that I can create a list from a number
where the 1's digits are at the head of the list.
*/
func TestNumberToReverseList(t *testing.T) {
	expectedList := createList([]int32{2, 1, 9})

	assert.Equal(t, expectedList.ToString(), numberToReverseList(912).ToString())
}

/*
This test makes sure that I can take a reversed list of numbers
where the 1's digits are at the head, and convert it to a number.
*/
func TestReverseListToNumber(t *testing.T) {
	list := createList([]int32{2, 1, 9})

	actualNumber, _ := reverseListToNumber(list)
	assert.Equal(t, int32(912), actualNumber)
}

/*
Let's make sure we get an error if we dared to send a node with multiple digits
*/
func TestReverseListToNumberWithMultipleDigitsReturnsError(t *testing.T) {
	list := createList([]int32{2, 11, 9})

	_, err := reverseListToNumber(list)
	assert.Error(t, err)
}

/*
Let's now test our first pass of summing lists of single digit numbers where
the head of the list represents the 1's digits.
*/
func TestSumReverseLists(t *testing.T) {
	list1 := createList([]int32{7, 1, 6})
	list2 := createList([]int32{5, 9, 2})

	expectedList := createList([]int32{2, 1, 9})

	actualList, _ := sumReverseLists(list1, list2)
	assert.Equal(t, expectedList.ToString(), actualList.ToString())
}

/*
This should run without a problem, but just to prove to ourselves, let's test
to make sure that adding two reversed lists that should create one more
digit, does.
*/

func TestSumReverseListsCreateLargerList(t *testing.T) {
	list1 := createList([]int32{7, 2, 9}) // 927
	list2 := createList([]int32{4, 9, 9}) // 994

	expectedList := createList([]int32{1, 2, 9, 1})

	actualList, _ := sumReverseLists(list1, list2)
	assert.Equal(t, expectedList.ToString(), actualList.ToString())
}

/*
Let's test our practice method for creating a list of nodes, printing
in order, digit by digit.
*/
func TestPrintNumberByDigits(t *testing.T) {
	x := 912
	var b bytes.Buffer

	printNumberByDigits(int32(x), &b)

	assert.Equal(t, "912", b.String())
}

/*
Let's test the method for creating a list, where the 1's digits are at the tail,
using our method that is similar to our printing of numbers by digits
*/
func TestNumberToList(t *testing.T) {
	expectedList := createList([]int32{9, 1, 2})

	assert.Equal(t, expectedList.ToString(), numberToList(912).ToString())
}

/*
Let's test the method for converting a list of digits, where the 1's digits
are at the end, into a number.
*/
func TestListToNumber(t *testing.T) {
	list := createList([]int32{9, 1, 2})
	actualNumber, _ := listToNumber(list)
	assert.Equal(t, int32(912), actualNumber)
}

/*
Let's make sure we get an error if we dared to send a node with multiple digits
*/
func TestListToNumberWithMultipleDigitsReturnsError(t *testing.T) {
	list := createList([]int32{2, 11, 9})

	_, err := listToNumber(list)
	assert.Error(t, err)
}

/*
Finally, let's test the sumLists method that uses the convenience
methods to create numbers from forward lists, and forward lists from
numbers.
*/
func TestSumLists(t *testing.T) {
	list1 := createList([]int32{6, 1, 7})
	list2 := createList([]int32{2, 9, 5})

	expectedList := createList([]int32{9, 1, 2})
	actualList, _ := sumLists(list1, list2)
	assert.Equal(t, expectedList.ToString(), actualList.ToString())
}

/*
Now, let's test the convenience methods we made for summing a forward
list in one pass.

First, summedNodeWithCarry
*/

func TestSummedNodeWithZeroCarryReturnsZeroCarry(t *testing.T) {
	carry, node, _ := summedNodeWithCarry(1, 2, 0)
	assert.Equal(t, carry, int32(0))
	assert.Equal(t, int32(3), node.data)
}

func TestSummedNodeWithZeroCarryReturnsNonZeroCarry(t *testing.T) {
	carry, node, _ := summedNodeWithCarry(9, 2, 0)
	assert.Equal(t, carry, int32(1))
	assert.Equal(t, int32(1), node.data)
}

func TestSummedNodeWithNonZeroCarryReturnsZeroCarry(t *testing.T) {
	carry, node, _ := summedNodeWithCarry(1, 2, 1)
	assert.Equal(t, carry, int32(0))
	assert.Equal(t, int32(4), node.data)
}

func TestSummedNodeWithNonZeroCarryReturnsNonZeroCarry(t *testing.T) {
	carry, node, _ := summedNodeWithCarry(9, 5, 1)
	assert.Equal(t, carry, int32(1))
	assert.Equal(t, int32(5), node.data)
}

// We should also test that if we dare to send a multidigit number in
// for the first or send value, we return an error

func TestSummedNodesWithMultiDigitForFirstParamReturnsError(t *testing.T) {
	_, _, err := summedNodeWithCarry(11, 9, 0)
	assert.Error(t, err)
}

func TestSummedNodesWithMultiDigitForSecondParamReturnsError(t *testing.T) {
	_, _, err := summedNodeWithCarry(1, 19, 0)
	assert.Error(t, err)
}

func TestSummedNodesWithMultiDigitForCarrayParamReturnsError(t *testing.T) {
	_, _, err := summedNodeWithCarry(1, 9, 10)
	assert.Error(t, err)
}

/*
Continuing the testing of the helper methods for summing a forward list
in one pass, we test the sumForwardListsWithCarry method.
*/

func TestSumForwardListsWithCarry(t *testing.T) {
	list1 := createList([]int32{6, 1, 7}) // 716
	list2 := createList([]int32{2, 9, 5}) // 592

	expectedList := createList([]int32{9, 1, 2}) // 912
	carry, actualList, _ := sumForwardListsWithCarry(list1, list2)
	assert.Equal(t, int32(0), carry)
	assert.Equal(t, expectedList.ToString(), actualList.ToString())
}

func TestSumForwardListsWithCarryReturnsNonZeroCarry(t *testing.T) {
	list1 := createList([]int32{9, 2, 7}) // 927
	list2 := createList([]int32{9, 9, 4}) // 994

	expectedList := createList([]int32{9, 2, 1}) // 921
	carry, actualList, _ := sumForwardListsWithCarry(list1, list2)
	assert.Equal(t, int32(1), carry)
	assert.Equal(t, expectedList.ToString(), actualList.ToString())
}

// We should also test that the summing of forward lists where one
// list contains multiple digits returns an error
func TestSumForwardListsWithCarryWithMultiDigitsReturnsError(t *testing.T) {
	list1 := createList([]int32{9, 2, 7})
	list2 := createList([]int32{9, 11, 4})

	_, _, err := sumForwardListsWithCarry(list1, list2)
	assert.Error(t, err)
}

/*
We have now tested the helper methods for summing a forward list in
one pass, testing the actual method should be straight forward (haha forward!)
*/
func TestSumForwardListsInOnePass(t *testing.T) {
	list1 := createList([]int32{6, 1, 7}) // 716
	list2 := createList([]int32{2, 9, 5}) // 592

	expectedList := createList([]int32{9, 1, 2}) // 912
	actualList, _ := sumForwardListsInOnePass(list1, list2)
	assert.Equal(t, expectedList.ToString(), actualList.ToString())
}

func TestSumForwardListsInOnePassWithCarry(t *testing.T) {
	list1 := createList([]int32{9, 2, 7}) // 927
	list2 := createList([]int32{9, 9, 4}) // 994

	expectedList := createList([]int32{1, 9, 2, 1})

	actualList, _ := sumForwardListsInOnePass(list1, list2)
	assert.Equal(t, expectedList.ToString(), actualList.ToString())
}

// we should test that calling the actual method with a list with
// mulitiple digits returns an error
func TestSumForwardListsInOnePassWithMultiDigitsReturnsError(t *testing.T) {
	list1 := createList([]int32{9, 12, 7})
	list2 := createList([]int32{9, 9, 4})

	_, err := sumForwardListsInOnePass(list1, list2)
	assert.Error(t, err)
}

// as we noted in the actual implementation file, we have to make sure
// that two lists with different node counts are summed up to the
// correct number
func TestSumForwardListsOfDifferentLengths(t *testing.T) {
	list1 := createList([]int32{9, 7})    // 97
	list2 := createList([]int32{9, 9, 4}) // 994

	expectedList := createList([]int32{1, 0, 9, 1}) // 1091
	actualList, _ := sumForwardListsInOnePass(list1, list2)
	assert.Equal(t, expectedList.ToString(), actualList.ToString())
}
