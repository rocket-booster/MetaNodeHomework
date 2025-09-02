package main

func removeDuplicates(nums []int) int {
    n := len(nums)
    if n <= 1{
        return n
    }
        
            
    slow := 0    
    for j := slow+1; j < n; j++{
        if nums[slow] != nums[j]{
            slow++;
            nums[slow] = nums[j]
        }        
    }
    return slow+1;
}