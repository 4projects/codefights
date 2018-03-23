

func evenDigitsOnly(n int) bool {
    s := strconv.Itoa(n)
    
    for _, i := range(s) {
        v, _ := strconv.Atoi(string(i))
        if v % 2 == 1 {
            return false
        }
    }
    return true
}
