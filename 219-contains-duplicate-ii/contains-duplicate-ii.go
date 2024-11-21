func containsNearbyDuplicate(nums []int, k int) bool {
    
    numsMap:=make(map[int]int)

    for index, _ := range nums{
        
        valPosition, ok:= numsMap[nums[index]]

        if ok && index != valPosition && (index- valPosition) <=k{
            return true
        }

        numsMap[nums[index]] = index    
    }

    return false

}