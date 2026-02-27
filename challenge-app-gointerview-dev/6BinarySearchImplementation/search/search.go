package search

/**
The arrays can be assumed to be sorted in ascending order.
All functions must implement the binary search algorithm, which has O(log n) time complexity.
BinarySearchRecursive must use recursion to solve the problem.
If multiple occurrences of the target exist, return the index of any occurrence.
**/

// BinarySearch performs a standard binary search to find the target in the sorted array.
// Returns the index of the target if found, or -1 if not found.
func BinarySearch(arr []int, target int) int {
	var foundIndex int
	length := len(arr)

	if length <= 0 {
		foundIndex = -1
	} else {
		foundIndex = -1
		leftIndex := 0
		rightIndex := length - 1
		var middleIndex int

		for leftIndex <= rightIndex {
			middleIndex = (leftIndex + rightIndex) / 2

			if arr[middleIndex] == target {
				foundIndex = middleIndex
				break
			} else if arr[middleIndex] > target {
				rightIndex = middleIndex - 1
			} else if arr[middleIndex] < target {
				leftIndex = middleIndex + 1
			}
		}
	}

	return foundIndex
}

// BinarySearchRecursive performs binary search using recursion.
// Returns the index of the target if found, or -1 if not found.
func BinarySearchRecursive(arr []int, target int, left int, right int) int {
	var foundIndex int

	if left < 0 || right < 0 {
		foundIndex = -1
	} else {
		length := right - left + 1

		if length <= 0 {
			foundIndex = -1
		} else {
			foundIndex = -1
			leftIndex := left
			rightIndex := right

			if leftIndex > rightIndex {
				foundIndex = -1
			} else {
				middleIndex := (leftIndex + rightIndex) / 2
				if arr[middleIndex] == target {
					foundIndex = middleIndex
				} else if arr[middleIndex] > target {
					foundIndex = BinarySearchRecursive(arr, target, leftIndex, middleIndex-1)
				} else if arr[middleIndex] < target {
					foundIndex = BinarySearchRecursive(arr, target, middleIndex+1, rightIndex)
				}
			}
		}
	}

	return foundIndex
}

// FindInsertPosition returns the index where the target should be inserted
// to maintain the sorted order of the array.
func FindInsertPosition(arr []int, target int) int {
	return 0
}
