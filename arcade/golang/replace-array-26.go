func arrayReplace(inputArray []int, elemToReplace int, substitutionElem int) []int {
    for k, i := range(inputArray) {
        if i == elemToReplace {
            inputArray[k] = substitutionElem
        }
    }
    return inputArray
}
