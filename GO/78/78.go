package main
import "fmt"

func subsets(nums []int) [][]int {
    var ans [][]int
    for i:= 0;i<(1<<len(nums));i++{
        out := []int{}
        for key,val :=range nums{
            if i>>key & 1 == 1{
                out = append(out,val)
            }
        }
        ans = append(ans,append([]int(nil),out...))
    }
    return ans
}

func main()  {
    input :=[] int {1,2,3 } 
    output := subsets(input)
    fmt.Println(output)
}