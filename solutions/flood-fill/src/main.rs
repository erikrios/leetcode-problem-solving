fn main() {
    println!(
        "{:?}",
        Solution::flood_fill(vec![vec![1, 1, 1], vec![1, 1, 0], vec![1, 0, 1]], 1, 1, 2)
    );
    println!(
        "{:?}",
        Solution::flood_fill(vec![vec![0, 0, 0], vec![0, 0, 0]], 0, 0, 0)
    );
}

struct Solution;

impl Solution {
    pub fn flood_fill(image: Vec<Vec<i32>>, sr: i32, sc: i32, color: i32) -> Vec<Vec<i32>> {
        let mut image = image;
        let target = image[sr as usize][sc as usize];

        Solution::fill(&mut image, sr, sc, color, target);

        image.to_vec()
    }

    fn fill(image: &mut Vec<Vec<i32>>, mut sr: i32, mut sc: i32, color: i32, target: i32) {
        if Solution::is_out_of_bounds(image, sr, sc) {
            return;
        }

        let current = image[sr as usize][sc as usize];
        if current != target {
            return;
        }

        if current == color {
            return;
        }

        image[sr as usize][sc as usize] = color;

        sr -= 1;
        Solution::fill(image, sr, sc, color, target);
        sr += 1;

        sc += 1;
        Solution::fill(image, sr, sc, color, target);
        sc -= 1;

        sr += 1;
        Solution::fill(image, sr, sc, color, target);
        sr -= 1;

        sc -= 1;
        Solution::fill(image, sr, sc, color, target);
    }

    fn is_out_of_bounds(image: &Vec<Vec<i32>>, sr: i32, sc: i32) -> bool {
        let m = image.len() as i32;
        let n = image[0].len() as i32;

        if sr < 0 || sr >= m {
            return true;
        }

        if sc < 0 || sc >= n {
            return true;
        }

        return false;
    }
}
