package lesson4

//Solve: 查找被‘X'围绕的以’O'为标记的区域
/*parameter
board: 以'X"、'O'为值的byte二维数组
*/
func Solve(board [][]byte) {

	//行数、列数
	rows, cols := len(board), len(board[0])
	if rows <= 0 || cols <= 0 {

		return
	}

	//以深度优先的方式查找二维数组board边界中的‘O'值并暂时标记为'#'
	//row: 行索引
	//col: 列索引
	var dfs func(row, col int)
	dfs = func(row, col int) {
		if row < 0 || row >= rows || col < 0 || col >= cols || board[row][col] != 'O' {

			return
		}

		board[row][col] = '#'
		//向下搜
		dfs(row+1, col)
		//向上搜
		dfs(row-1, col)
		//向右搜
		dfs(row, col+1)
		//向左搜
		dfs(row, col-1)
	}

	//搜索边界的’O'值并标记为'#'
	for row := 0; row < rows; row++ {

		dfs(row, 0)
		dfs(row, cols-1)
	}
	for col := 1; col < cols-1; col++ {

		dfs(0, col)
		dfs(rows-1, col)
	}

	//遍历行列，还原'#'为‘O',并将其余'O'值变更为'X’
	for row := 0; row < rows; row++ {

		for col := 0; col < cols; col++ {

			if board[row][col] == '#' {

				board[row][col] = 'O'
			} else {

				board[row][col] = 'X'
			}
		}
	}
}
