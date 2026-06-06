func twoSum(nums []int, target int) []int {
    sum := make(map[int]int)
    for i,num := range nums {
        value := target - num
        if index, exists := sum[value]; exists {
            return []int{index, i}
        }
        sum[num] = i
    }
    return nil;
}
