package fbinterview

//https://www.careercup.com/question?id=5706117911937024

//Given a matrix of N * M, given the coordinates (x, y) present in a matrix,
//Find the number of paths that can reach (0, 0) from the (x, y) points with k steps (requires exactly k, k> = 0)
//From each point you can go up, down, left and right in four directions.
//
//For example, the following matrix:
//----------
//0 0 0 0 0
//0 0 0 0 0
//0 0 0 0 0
//0 0 0 0 0
//(x, y) = (1, 1), k = 2, the output is 2

type Task struct {
	M, N, K          int
	OpsBrut, OpsPoly int
}

type DP [][][]int

var moves = []struct{ dx, dy int }{
	{dx: 1, dy: 0},
	{dx: 0, dy: 1},
	{dx: -1, dy: 0},
	{dx: 0, dy: -1},
}

func (t *Task) FindPathBruteForce(x, y, path int) int {

	if !(0 <= x && x <= t.M && 0 <= y && y <= t.N) || path > t.K {
		return 0
	}

	if x == 0 && y == 0 && t.K == path {
		return 1
	}

	numPath := 0
	for _, mv := range moves {
		numPath += t.FindPathBruteForce(x+mv.dx, y+mv.dy, path+1)
		t.OpsBrut++
	}

	return numPath
}

func (t *Task) FindPathPoly(sx, sy, path int) int {

	dp := make(DP, t.M)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([][]int, t.N)

		for j := 0; j < len(dp[i]); j++ {
			dp[i][j] = make([]int, t.K+1)
		}
	}

	dp[sx][sy][0] = 1

	for k := 1; k <= t.K; k++ {
		for x := 0; x < t.M; x++ {
			for y := 0; y < t.N; y++ {

				for _, mv := range moves {
					nx := x + mv.dx
					ny := y + mv.dy

					if !(0 <= nx && nx < t.M) {
						continue
					}
					if !(0 <= ny && ny < t.N) {
						continue
					}

					dp[nx][ny][k] += dp[x][y][k-1]
					t.OpsPoly++
				}
			}
		}
	}

	return dp[0][0][t.K]
}
