# 0026/Remove Duplicates from Sorted Array

<https://leetcode.com/problems/remove-duplicates-from-sorted-array/>
This problem asks to remove the duplicates from an integer array `nums`.
The removal should be done in-place. `nums` is already sorted in nondecreasing order.

I've solved this puzzle with two pointers[^1] approach.

[^1]: actually just slice indices, but these could be pointers too in e.g. `C`.
  I've always seen both pointers and indices called _pointers_ approach anyway.

### Algorithm analysis
First, `left` index is initialized to `0`.

Then `right` index is initialized to `1` and used to traverse the entire array (from second to last element) in a loop. On each iteration, the values in the array pointed to by `left` and `right` index are compared.
If the values are the same, nothing happens - `right` index proceeds to the next value.
If the values are different, then `left` index is increased by one, and the value from `right` index is written to the address pointed to by `left`.
As minor optimization, the overwriting `right` -> `left` happens only when `left` is less than `right`. This makes the overwriting not happen when `left` and `right` point to the same address. It would still result in valid answer.

At the end of such operation `left` points to the last unique value in the array. Since the puzzle specifies that `removeDuplicates` should return the length of uniques, `left + 1` is returned - Golang uses zero-indexing.

### Detailed run: Example 02
The input array is `[0, 0, 1, 1, 1, 2, 2, 3, 3, 4]`.
The algorithm performs as follows:
```
0, 0, 1, 1, 1, 2, 2, 3, 3, 4
L  R

0, 0, 1, 1, 1, 2, 2, 3, 3, 4
L     R

0, 1, 1, 1, 1, 2, 2, 3, 3, 4
   L     R

0, 1, 1, 1, 1, 2, 2, 3, 3, 4
   L        R

0, 1, 1, 1, 1, 2, 2, 3, 3, 4
   L           R

0, 1, 2, 1, 1, 2, 2, 3, 3, 4
      L           R

0, 1, 2, 1, 1, 2, 2, 3, 3, 4
      L              R

0, 1, 2, 3, 1, 2, 2, 3, 3, 4
         L              R

0, 1, 2, 3, 1, 2, 2, 3, 3, 4
         L                 R

0, 1, 2, 3, 4, 2, 2, 3, 3, 4
            L                 R

at the loop's end, L = 4; 4+1 = 5 is returned.
```

### Corner cases
For the degenerated case of an array of length 1 the algorithm still works correctly.
No modifications are made to the array in this case and the correct value of 1 is returned.

For the case of an array containing same integer repeated n times, `left` index never increases.
At the end of the loop, no modifications to the input array are made and the correct value of 1 is returned.

### Complexity analysis
Runtime: O(n), as `right` pointer has to traverse the entire `nums` array.
Memory: O(1), the algorithm requires space to hold two `int` indices, `left` and `right`.

