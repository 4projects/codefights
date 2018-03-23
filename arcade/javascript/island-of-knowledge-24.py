directions = [(i,j) for j in range(-1,2) for i in range(-1, 2)]

def minesweeper(matrix):
    arr = [[0 for _ in matrix[0] ] for _ in matrix]
    print(directions)
    r_bound = len(matrix) - 1 
    c_bound = len(matrix[0]) - 1 
    for k1, i in enumerate(matrix):
        for k2, j in enumerate(i):
            
            for t in directions:
                if t == (0, 0):
                    continue
                r = k1 + t[0]
                c = k2 + t[1]
                if r < 0 or r > r_bound:
                    continue
                if c < 0 or c > c_bound:
                    continue
                if matrix[r][c] == True:
                    arr [k1][k2] += 1
    return arr
                
            
