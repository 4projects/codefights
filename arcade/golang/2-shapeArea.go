func shapeArea(n int) int {
    s := 0
    for i:=1; i <n; i++ {
        s += (i*2-1) * 2
    }
    
    return s + n * 2 - 1
}
