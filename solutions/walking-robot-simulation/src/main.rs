use std::collections::HashSet;

fn main() {
    println!("{}", Solution::robot_sim(vec![4, -1, 3], vec![]));
    println!(
        "{}",
        Solution::robot_sim(vec![4, -1, 4, -2, 4], vec![vec![2, 4]])
    );
    println!("{}", Solution::robot_sim(vec![6, -1, -1, 6], vec![]));
}

struct Solution;

enum Direction {
    North,
    East,
    South,
    West,
}

impl Solution {
    pub fn robot_sim(commands: Vec<i32>, obstacles: Vec<Vec<i32>>) -> i32 {
        let mut x = 0;
        let mut y = 0;
        let mut current_direction = Direction::North;
        let mut obstacles_set = HashSet::new();

        for obstacle in obstacles {
            obstacles_set.insert((obstacle[0], obstacle[1]));
        }

        let mut max_euclidean_distance = 0;

        'outer: for command in commands {
            match command {
                -2 => match current_direction {
                    Direction::North => current_direction = Direction::West,
                    Direction::East => current_direction = Direction::North,
                    Direction::South => current_direction = Direction::East,
                    Direction::West => current_direction = Direction::South,
                },
                -1 => match current_direction {
                    Direction::North => current_direction = Direction::East,
                    Direction::East => current_direction = Direction::South,
                    Direction::South => current_direction = Direction::West,
                    Direction::West => current_direction = Direction::North,
                },
                units => {
                    let mut dir = (0, 0);

                    match current_direction {
                        Direction::North => {
                            dir.1 += 1;
                        }
                        Direction::East => {
                            dir.0 += 1;
                        }
                        Direction::South => {
                            dir.1 -= 1;
                        }
                        Direction::West => {
                            dir.0 -= 1;
                        }
                    }

                    for _ in 0..units {
                        x += dir.0;
                        y += dir.1;

                        if obstacles_set.contains(&(x, y)) {
                            x -= dir.0;
                            y -= dir.1;
                            continue 'outer;
                        }

                        max_euclidean_distance = max_euclidean_distance.max(x * x + y * y)
                    }
                }
            }
        }

        max_euclidean_distance
    }
}
