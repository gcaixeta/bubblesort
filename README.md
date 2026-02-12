# Simple BubbleSort - Python - ED

A simple Python implementation of the Bubble Sort algorithm.
The program sorts an array of integers in ascending order and prints both the initial and final state of the array.

The implementation also includes a small optimization:
if during a full pass no swaps are performed, the algorithm stops early.

## How it works

Bubble Sort repeatedly traverses the list, compares adjacent elements, and swaps them if they are in the wrong order.
This process continues until the array is fully sorted.

Example input:

```
  [13, 2, 43, 25, 4, 4, 52, 63, 12]
```

Output:

```
  [2, 4, 4, 12, 13, 25, 43, 52, 63]
```

## Time Complexity

Best case: O(n)
Occurs when the array is already sorted (no swaps are needed).

Average case: O(n²)

Worst case: O(n²)
Occurs when the array is in reverse order.

## Space Complexity

O(1) (constant space)

The algorithm sorts the array in-place, using only a small number of extra variables.

## Technologies

Python 3

## Author

Gustavo Caixeta
