Two Pointers Approach

    1. Sorted or Monotonic Data
Many two-pointer problems either require the data to be sorted beforehand or naturally involve sorted or monotonic data (increasing or decreasing). Sorting helps in making efficient decisions about pointer movements.

    2. Pointer Initialization
Two pointers are typically initialized at specific positions:
    • Start and End: Common in problems involving searching for pairs or specific conditions (e.g., finding two elements that sum to a target in a sorted array).
    • Adjacent Start Positions: Useful in problems like sliding window techniques (e.g., finding subarrays with certain properties).
    
    3. Iterative Pointer Movement
The two pointers move towards each other or in a specific direction based on certain conditions. Movement is usually done in a loop until a termination condition is met (e.g., pointers cross each other).

    4. Conditional Logic for Pointer Movement
The movement of the pointers depends on the evaluation of certain conditions:
    • Comparison: Comparing the sum/product/difference of elements at the pointers with a target value.
    • Window Constraints: Checking if a certain window size or condition is met.
    
    5. Efficient Pair Counting or Window Management
The goal is often to count pairs or manage a window of elements efficiently:
    • Counting Valid Pairs: As seen in the two-pointer sum problem.
    • Maintaining a Window: Adjusting window size based on conditions to find the minimum/maximum length subarray meeting a criterion.
    
    6. Optimization Over Brute Force
Two-pointer approaches optimize the time complexity compared to naive methods

    7. Applications in Various Problems
Two-pointer techniques are used in a variety of problems, such as:
    • Pair Problems: Finding pairs with specific properties (e.g., sum, difference).
    • Subarray Problems: Finding subarrays that meet certain criteria (e.g., sum, product).
    • String Problems: Manipulating or analyzing substrings.



Let's explain both pointer initialization techniques with examples in Go.

    1. Start and End Pointers

Problem: Finding Two Elements that Sum to a Target in a Sorted Array**

Given a sorted array `arr` and a target sum `target`, find two elements in the array that add up to the target.

Algorithm:
1. Initialize two pointers: `left` at the start (0) and `right` at the end (len(arr)-1).
2. Calculate the sum of the elements at these pointers.
3. If the sum is equal to the target, return the indices of the elements.
4. If the sum is less than the target, increment the `left` pointer to increase the sum.
5. If the sum is greater than the target, decrement the `right` pointer to decrease the sum.
6. Repeat steps 2-5 until the pointers meet.

Code Example:

func twoSumSorted(arr []int, target int) (int, int) {
    left, right := 0, len(arr)-1

    for left < right {
        sum := arr[left] + arr[right]
        if sum == target {
            return left, right
        } else if sum < target {
            left++
        } else {
            right--
        }
    }
    return -1, -1
}

func main() {
    arr := []int{1, 2, 3, 4, 5, 6}
    target := 9
    left, right := twoSumSorted(arr, target)
    if left != -1 && right != -1 {
        fmt.Printf("Indices: %d, %d\n", left, right)
    } else {
        fmt.Println("No solution found.")
    }
}


     2. Adjacent Start Positions

Problem: Finding the Longest Subarray with Sum Less Than or Equal to a Target**

Given an array `arr` and a target sum `target`, find the longest subarray whose sum is less than or equal to the target.

Algorithm:
1. Initialize two pointers: `start` and `end` at the beginning of the array.
2. Initialize a variable to keep track of the current sum.
3. Expand the `end` pointer to include elements in the window and update the current sum.
4. If the current sum exceeds the target, move the `start` pointer to reduce the sum.
5. Keep track of the maximum length of subarrays found during the process.
6. Repeat steps 3-5 until the `end` pointer reaches the end of the array.

Code Example:

func longestSubarrayWithSum(arr []int, target int) int {
    i, j:= 0, 0
    currentSum := 0
    maxLength := 0

    for j< len(arr) {
        currentSum += arr[j]
        
        if j-i+1 > maxLength {
            maxLength = j- i+ 1
        }
        

        for currentSum > target && i<= j{
            currentSum -= arr[i]
            i++
        }

        j++
    }

    return maxLength
}

func main() {
    arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    target := 15
    maxLength := longestSubarrayWithSum(arr, target)
    fmt.Printf("The longest subarray length with sum <= %d is %d\n", target, maxLength)
}




Summary

- Start and End Pointers are useful for problems like finding pairs that meet a condition, often with sorted arrays.
- Adjacent Start Positions are used in sliding window techniques to find subarrays with specific properties.


SLIDING WINDOW

Two types:
    1. fixed size window           : -  K  will be given
    2. variable size window      : -  condition will be given (Max/min or largest/smallest)

Identification : 
    1.  array/string problem
    2.  K or Condition   given
    3.  subarray , substring

    
    Notes:
    Window will be represent with 2 pointers : i and j
    Always size of window  :   j- i + 1



Top 10 Sliding Window Interview Problems

    1. Best time to buy sell stocks
    
    2. Longest Substring Without Repeating Characters
Determine the length of the longest substring without duplicate characters.
    
    3. Longest Repeating Character Replacement
Find the length of the longest substring that can be obtained by replacing at most K characters to make all characters the same.
    
    4. Maximum Sum Subarray of Size K
Find the maximum sum of any contiguous subarray of size K.
    
    5. Longest Subarray with Ones after Replacement : https://leetcode.com/problems/max-consecutive-ones-iii/description/ 
    Given a binary array, find the length of the longest subarray containing only 1s after replacing at most K 0s.
    (Same as problem number 3)
    
    6. Smallest Subarray with Sum Greater than S/target : https://leetcode.com/problems/minimum-size-subarray-sum/description/ 
    Find the minimal length of a contiguous subarray of which the sum is greater than or equal to a given value S/target.
    
    7. Fruits into Baskets
    Given an array representing types of fruits, find the length of the longest subarray with at most two different types.
    
    8. Find All Anagrams in a String
    Find all start indices of p's anagrams in s.
    
    9. Minimum Window Substring
Identify the smallest substring in a given string that contains all characters of another string.
    
    10. Permutation in a String
Check if s2 contains a permutation of s1.
    
    11. Count Distinct Elements In Every Window of Size K: https://leetcode.com/problems/subarrays-with-k-different-integers/description/
Given an array and a number K, count distinct elements in all windows of size K.


    Heap:

Ways to identify heap problem:
    1. Kth
    2. Smallest/Largest

 K + smallest ---> max heap
K + largest.  -----> min heap

