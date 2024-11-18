func singleNumber(nums []int) int {
    
    hasMap := make(map[int]bool)

    for i:=0; i<len(nums); i++{
        _, ok:=hasMap[nums[i]]
        if ok{
            hasMap[nums[i]] = true
            continue
        }
        hasMap[nums[i]] = false
    }

    for k, _:= range hasMap{
        if !hasMap[k]{
            return k
        }
    }
    return 0
}