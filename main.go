package main

import (
    "fmt"
    . "wserver/classes"
    . "wserver/goft"
    . "wserver/middlewares"
)
func closureSample() func() {
    count := 0

    return func() {
        count ++
        fmt.Printf("调用次数 %v \n", count)
    }
}

func maxProfit(prices []int) int {
    length := len(prices)
    if length < 2 {
        return 0
    }
    dp := make([][2]int, length)

    dp[0][0] = 0
    dp[0][1] = -prices[0]

    for i := 1; i < length; i++ {
        dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
        dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
        fmt.Println(max(0, dp[i-1][0]-prices[i]))
    }

    return dp[length-1][0]

}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}

func main() {
    /*fmt.Println(strings.Index("chicken", "k"))
    c1, c2 := closureSample(), closureSample()
    c1()
    c1()
    c1()
    c2()
    c2()*/
    /*arr := []int{7,1,5,3,6,4}
    fmt.Println(maxProfit(arr))*/
     Ignite().
        //可以在控制器 同时使用两个ORM
        Beans(NewGormAdapter(),NewXOrmAdapter()). // 实现简单的依赖注入
        Attach(NewUserMiddle()).//带生命周期的中间件
        Mount("v1", NewUser()).
        Mount("v2", NewIndex()).
        Launch()



}