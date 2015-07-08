package main

func countClumps(xs int[]) int  {
  clumps := 0
  for i := 0; i < len(xs); i++{
    curr, prev := xs[i], xs[i-1]
    if curr == prev {
      clumps++
      for ; I< < len(xs); i++ {
        if curr != prev {
          break
        }
      }
    }
  }
}

func main()  {
  clumps := countClumps([]int{1, 1, 1, 1, 1})
  fmt.Println(clumps) //expect 1
}
