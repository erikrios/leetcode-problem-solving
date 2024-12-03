fn main() {
    println!(
        "{}",
        Solution::count_unguarded(
            4,
            6,
            vec![vec![0, 0], vec![1, 1], vec![2, 3]],
            vec![vec![0, 1], vec![2, 2], vec![1, 4]]
        )
    );
    println!(
        "{}",
        Solution::count_unguarded(
            3,
            3,
            vec![vec![1, 1]],
            vec![vec![0, 1], vec![1, 0], vec![2, 1], vec![1, 2]]
        )
    );
    println!(
        "{}",
        Solution::count_unguarded(1, 100000, vec![vec![0, 0]], vec![vec![0, 1]])
    );
}

struct Solution;

#[derive(Clone, Debug, PartialEq)]
enum CellType {
    Guard,
    Wall,
    Guarded(GuardedType),
    Unguarded,
}

#[derive(Clone, Debug, PartialEq)]
enum GuardedType {
    Vertical,
    Horizontal,
    VerticalHorizontal,
}

#[derive(Debug)]
enum Direction {
    Left((i32, i32)),
    Right((i32, i32)),
    Up((i32, i32)),
    Down((i32, i32)),
    Unspecified,
}

const DIRECTIONS: [Direction; 4] = [
    Direction::Left((0, -1)),
    Direction::Right((0, 1)),
    Direction::Up((-1, 0)),
    Direction::Down((1, 0)),
];

impl Solution {
    pub fn count_unguarded(m: i32, n: i32, guards: Vec<Vec<i32>>, walls: Vec<Vec<i32>>) -> i32 {
        let mut matrix = vec![vec![CellType::Unguarded; n as usize]; m as usize];

        for wall in walls {
            matrix[wall[0] as usize][wall[1] as usize] = CellType::Wall;
        }

        for guard in &guards {
            matrix[guard[0] as usize][guard[1] as usize] = CellType::Guard;
        }

        for guard in guards {
            Self::mark_guarded(
                guard[0] as usize,
                guard[1] as usize,
                &mut matrix,
                Direction::Unspecified,
            );
        }

        let mut counter = 0;

        for i in 0..m as usize {
            for j in 0..n as usize {
                if matrix[i][j] == CellType::Unguarded {
                    counter += 1;
                }
            }
        }

        counter
    }

    fn mark_guarded(i: usize, j: usize, matrix: &mut [Vec<CellType>], direction: Direction) {
        let cell = &matrix[i][j];
        let m = matrix.len();
        let n = matrix[0].len();

        match direction {
            Direction::Unspecified => {
                for dir in DIRECTIONS {
                    match dir {
                        Direction::Left((i_val, j_val)) => {
                            if j > 0 {
                                Self::mark_guarded(
                                    (i as i32 + i_val) as usize,
                                    (j as i32 + j_val) as usize,
                                    matrix,
                                    dir,
                                )
                            }
                        }
                        Direction::Right((i_val, j_val)) => {
                            if j < n - 1 {
                                Self::mark_guarded(
                                    (i as i32 + i_val) as usize,
                                    (j as i32 + j_val) as usize,
                                    matrix,
                                    dir,
                                )
                            }
                        }
                        Direction::Up((i_val, j_val)) => {
                            if i > 0 {
                                Self::mark_guarded(
                                    (i as i32 + i_val) as usize,
                                    (j as i32 + j_val) as usize,
                                    matrix,
                                    dir,
                                )
                            }
                        }
                        Direction::Down((i_val, j_val)) => {
                            if i < m - 1 {
                                Self::mark_guarded(
                                    (i as i32 + i_val) as usize,
                                    (j as i32 + j_val) as usize,
                                    matrix,
                                    dir,
                                )
                            }
                        }
                        _ => (),
                    }
                }
            }
            dir => match cell {
                CellType::Wall | CellType::Guard => (),
                CellType::Guarded(guarded_type) => match *guarded_type {
                    GuardedType::Vertical => match dir {
                        Direction::Left((i_val, j_val)) => {
                            if j > 0 {
                                matrix[i][j] = CellType::Guarded(GuardedType::VerticalHorizontal);
                                Self::mark_guarded(
                                    (i as i32 + i_val) as usize,
                                    (j as i32 + j_val) as usize,
                                    matrix,
                                    dir,
                                )
                            }
                        }
                        Direction::Right((i_val, j_val)) => {
                            if j < n - 1 {
                                matrix[i][j] = CellType::Guarded(GuardedType::VerticalHorizontal);
                                Self::mark_guarded(
                                    (i as i32 + i_val) as usize,
                                    (j as i32 + j_val) as usize,
                                    matrix,
                                    dir,
                                )
                            }
                        }
                        _ => (),
                    },
                    GuardedType::Horizontal => match dir {
                        Direction::Up((i_val, j_val)) => {
                            matrix[i][j] = CellType::Guarded(GuardedType::VerticalHorizontal);
                            if i > 0 {
                                Self::mark_guarded(
                                    (i as i32 + i_val) as usize,
                                    (j as i32 + j_val) as usize,
                                    matrix,
                                    dir,
                                )
                            }
                        }
                        Direction::Down((i_val, j_val)) => {
                            matrix[i][j] = CellType::Guarded(GuardedType::VerticalHorizontal);
                            if i < m - 1 {
                                Self::mark_guarded(
                                    (i as i32 + i_val) as usize,
                                    (j as i32 + j_val) as usize,
                                    matrix,
                                    dir,
                                )
                            }
                        }
                        _ => (),
                    },
                    GuardedType::VerticalHorizontal => (),
                },
                _ => match dir {
                    Direction::Left((i_val, j_val)) => {
                        matrix[i][j] = CellType::Guarded(GuardedType::Horizontal);
                        if j > 0 {
                            Self::mark_guarded(
                                (i as i32 + i_val) as usize,
                                (j as i32 + j_val) as usize,
                                matrix,
                                dir,
                            )
                        }
                    }
                    Direction::Right((i_val, j_val)) => {
                        matrix[i][j] = CellType::Guarded(GuardedType::Horizontal);
                        if j < n - 1 {
                            Self::mark_guarded(
                                (i as i32 + i_val) as usize,
                                (j as i32 + j_val) as usize,
                                matrix,
                                dir,
                            )
                        }
                    }
                    Direction::Up((i_val, j_val)) => {
                        matrix[i][j] = CellType::Guarded(GuardedType::Vertical);
                        if i > 0 {
                            Self::mark_guarded(
                                (i as i32 + i_val) as usize,
                                (j as i32 + j_val) as usize,
                                matrix,
                                dir,
                            )
                        }
                    }
                    Direction::Down((i_val, j_val)) => {
                        matrix[i][j] = CellType::Guarded(GuardedType::Vertical);
                        if i < m - 1 {
                            Self::mark_guarded(
                                (i as i32 + i_val) as usize,
                                (j as i32 + j_val) as usize,
                                matrix,
                                dir,
                            )
                        }
                    }
                    _ => (),
                },
            },
        }
    }
}
