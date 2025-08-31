package main


func twoSum(nums []int, target int) []int {
    n := len(nums)
    partNumMap := make(map[int]int , n)

    for i := 0; i < n; i++{
        if p, ok := partNumMap[target - nums[i]]; ok{
            return []int{i, p}
        }
        partNumMap[nums[i]] = i 
    }

    return nil
}