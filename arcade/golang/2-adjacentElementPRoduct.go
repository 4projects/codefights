
func adjacentElementsProduct(inputArray []int) int { 
    last := inputArray[0]
    mProduct := last * inputArray[1]
    for _, i := range inputArray[1:] {
        prod := i * last
        if prod > mProduct {
            mProduct = prod
        }
        last = i
    }
    return mProduct
}
