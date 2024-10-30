fn main() {
    println!(
        "{}",
        Solution::count_battleships(vec![
            vec!['X', '.', '.', 'X'],
            vec!['.', '.', '.', 'X'],
            vec!['.', '.', '.', 'X']
        ])
    );
    println!("{}", Solution::count_battleships(vec![vec!['.']]));
}

struct Solution;

impl Solution {
    pub fn count_battleships(board: Vec<Vec<char>>) -> i32 {
        let m = board.len();
        let n = board[0].len();

        let mut board = board;

        let mut counter = 0;
        for i in 0..m {
            for j in 0..n {
                if board[i][j] != 'X' {
                    continue;
                }

                Self::dfs(&mut board, i as i32, j as i32);

                counter += 1;
            }
        }

        counter
    }

    fn dfs(board: &mut [Vec<char>], x: i32, y: i32) {
        let m = board.len();
        let n = board[0].len();

        if !(x >= 0 && x < m as i32 && y >= 0 && y < n as i32) {
            return;
        }

        let x = x as usize;
        let y = y as usize;
        let cell = board[x][y];
        if cell == '.' {
            return;
        }

        board[x][y] = '.';

        let directions: [(i32, i32); 4] = [(-1, 0), (1, 0), (0, -1), (0, 1)];

        for dir in directions {
            Self::dfs(board, x as i32 + dir.0, y as i32 + dir.1);
        }
    }
}
