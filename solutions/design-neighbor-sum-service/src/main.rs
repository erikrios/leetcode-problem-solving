use std::collections::HashMap;

fn main() {
    let neigbor_sum = neighborSum::new(vec![vec![0, 1, 2], vec![3, 4, 5], vec![6, 7, 8]]);
    println!(
        "{} {} {} {}",
        neigbor_sum.adjacent_sum(1),
        neigbor_sum.adjacent_sum(4),
        neigbor_sum.diagonal_sum(4),
        neigbor_sum.diagonal_sum(8)
    );

    let neigbor_sum = neighborSum::new(vec![
        vec![1, 2, 0, 3],
        vec![4, 7, 15, 6],
        vec![8, 9, 10, 11],
        vec![12, 13, 14, 5],
    ]);
    println!(
        "{} {}",
        neigbor_sum.adjacent_sum(15),
        neigbor_sum.diagonal_sum(9),
    );
}

struct neighborSum {
    index_mapper: HashMap<usize, [usize; 2]>,
    n: usize,
    grid: Vec<Vec<i32>>,
}

impl neighborSum {
    fn new(grid: Vec<Vec<i32>>) -> Self {
        let mut index_mapper = HashMap::new();

        for (i, nums) in grid.iter().enumerate() {
            for (j, &num) in nums.iter().enumerate() {
                index_mapper.insert(num as usize, [i, j]);
            }
        }

        let n = grid.len();

        Self {
            index_mapper,
            n,
            grid,
        }
    }

    fn adjacent_sum(&self, value: i32) -> i32 {
        let indexes = self.index_mapper.get(&(value as usize)).unwrap();
        let i = indexes[0];
        let j = indexes[1];

        let mut sum = 0;

        if self.is_inside(self.n, i as i32 - 1, j as i32) {
            sum += self.grid[i - 1][j];
        }

        if self.is_inside(self.n, i as i32 + 1, j as i32) {
            sum += self.grid[i + 1][j];
        }

        if self.is_inside(self.n, i as i32, j as i32 - 1) {
            sum += self.grid[i][j - 1];
        }

        if self.is_inside(self.n, i as i32, j as i32 + 1) {
            sum += self.grid[i][j + 1];
        }

        sum
    }

    fn diagonal_sum(&self, value: i32) -> i32 {
        let indexes = self.index_mapper.get(&(value as usize)).unwrap();
        let i = indexes[0];
        let j = indexes[1];

        let mut sum = 0;

        if self.is_inside(self.n, i as i32 - 1, j as i32 - 1) {
            sum += self.grid[i - 1][j - 1];
        }

        if self.is_inside(self.n, i as i32 + 1, j as i32 + 1) {
            sum += self.grid[i + 1][j + 1];
        }

        if self.is_inside(self.n, i as i32 + 1, j as i32 - 1) {
            sum += self.grid[i + 1][j - 1];
        }

        if self.is_inside(self.n, i as i32 - 1, j as i32 + 1) {
            sum += self.grid[i - 1][j + 1];
        }

        sum
    }

    fn is_inside(&self, n: usize, i: i32, j: i32) -> bool {
        i >= 0 && i < n as i32 && j >= 0 && j < n as i32
    }
}
