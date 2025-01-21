func robHeler(arr []int, n int, memo map[int]int)int{
    if data, ok:=memo[n]; ok{

        return data
    }
    if n< 0{

        return 0
    }

    rob:= arr[n]+robHeler(arr, n-2, memo)
    skip:= robHeler(arr, n-1, memo)

    max:=rob
    if skip>max{
        max = skip
    }
    memo[n] = max
    return memo[n]

}

func rob(nums []int) int {
    memo:= make(map[int]int)
    return robHeler(nums, len(nums)-1, memo)
}