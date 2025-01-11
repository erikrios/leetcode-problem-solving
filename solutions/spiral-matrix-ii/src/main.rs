fn main() {
    println!("{:?}", Solution::generate_matrix(3));
    println!("{:?}", Solution::generate_matrix(1));
}

struct Solution;

enum Direction {
    Right,
    Bottom,
    Left,
    Top,
}

impl Solution {
    pub fn generate_matrix(n: i32) -> Vec<Vec<i32>> {
        let n = n as usize;
        let mut matrix = Vec::with_capacity(n);

        for _ in 0..n {
            let nums = vec![0; n];
            matrix.push(nums);
        }

        Self::generate(&mut matrix, (0, 0), Direction::Right, 1);

        matrix
    }

    fn generate(matrix: &mut [Vec<i32>], coordinates: (usize, usize), dir: Direction, i: usize) {
        let n = matrix.len();

        matrix[coordinates.0][coordinates.1] = i as i32;

        match dir {
            Direction::Right => {
                let (possible_next_direction, new_coordinates) = if coordinates.1 + 1 < n
                    && matrix[coordinates.0][coordinates.1 + 1] == 0
                {
                    (Direction::Right, (coordinates.0, coordinates.1 + 1))
                } else if coordinates.0 + 1 < n && matrix[coordinates.0 + 1][coordinates.1] == 0 {
                    (Direction::Bottom, (coordinates.0 + 1, coordinates.1))
                } else {
                    return;
                };
                Self::generate(matrix, new_coordinates, possible_next_direction, i + 1)
            }
            Direction::Bottom => {
                let (possible_next_direction, new_coordinates) =
                    if coordinates.0 + 1 < n && matrix[coordinates.0 + 1][coordinates.1] == 0 {
                        (Direction::Bottom, (coordinates.0 + 1, coordinates.1))
                    } else if coordinates.1 > 0 && matrix[coordinates.0][coordinates.1 - 1] == 0 {
                        (Direction::Left, (coordinates.0, coordinates.1 - 1))
                    } else {
                        return;
                    };
                Self::generate(matrix, new_coordinates, possible_next_direction, i + 1)
            }
            Direction::Left => {
                let (possible_next_direction, new_coordinates) =
                    if coordinates.1 > 0 && matrix[coordinates.0][coordinates.1 - 1] == 0 {
                        (Direction::Left, (coordinates.0, coordinates.1 - 1))
                    } else if coordinates.0 > 0 && matrix[coordinates.0 - 1][coordinates.1] == 0 {
                        (Direction::Top, (coordinates.0 - 1, coordinates.1))
                    } else {
                        return;
                    };
                Self::generate(matrix, new_coordinates, possible_next_direction, i + 1)
            }
            Direction::Top => {
                let (possible_next_direction, new_coordinates) = if coordinates.0 > 0
                    && matrix[coordinates.0 - 1][coordinates.1] == 0
                {
                    (Direction::Top, (coordinates.0 - 1, coordinates.1))
                } else if coordinates.1 + 1 < n && matrix[coordinates.0][coordinates.1 + 1] == 0 {
                    (Direction::Right, (coordinates.0, coordinates.1 + 1))
                } else {
                    return;
                };
                Self::generate(matrix, new_coordinates, possible_next_direction, i + 1)
            }
        }
    }
}
