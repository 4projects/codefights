function makeArrayConsecutive2(statues) {
    arr = []
    
    statues.map(a=>arr[a]=a)
    firstNotEmpty = arr.find((i,a)=>i!==undefined)
    return arr.length - statues.length - firstNotEmpty
}
