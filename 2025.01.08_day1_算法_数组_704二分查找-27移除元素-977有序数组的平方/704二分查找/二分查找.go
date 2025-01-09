package main

import "fmt"

func main(){
    nums := [...]int{-1,0,3,5,9,12}
    target := 9

    // 初始化左右边界
    left := 0
    right := len(nums)-1

    // 定义查找区间为左闭右闭[left,right]
    for left <= right{
        // 求中间位置的下标
        mid := left + (right-left)>>1
        // 根据nums[mid]和target大小确定查找区间的范围
        if nums[mid] == target{
            fmt.Println("target下标为:",mid)
            return 
        }else if nums[mid] < target{
            left = mid + 1
        }else{
            right = mid -1
        }
    }
    fmt.Println("未找到target")
    return
}

