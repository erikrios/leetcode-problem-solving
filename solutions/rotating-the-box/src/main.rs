fn main() {
    println!("{:?}", Solution::rotate_the_box(vec![vec!['#', '.', '#']]));
    println!(
        "{:?}",
        Solution::rotate_the_box(vec![vec!['#', '.', '*', '.'], vec!['#', '#', '*', '.']])
    );
    println!(
        "{:?}",
        Solution::rotate_the_box(vec![
            vec!['#', '#', '*', '.', '*', '.'],
            vec!['#', '#', '#', '*', '.', '.'],
            vec!['#', '#', '#', '.', '#', '.']
        ])
    );
}

struct Solution;

impl Solution {
    pub fn rotate_the_box(r#box: Vec<Vec<char>>) -> Vec<Vec<char>> {
        let mut r#box = r#box;
        Self::apply_gravity(&mut r#box);
        Self::rotate(r#box)
    }

    fn apply_gravity(r#box: &mut [Vec<char>]) {
        let m = r#box.len();
        let n = r#box[0].len();

        for chars in r#box.iter_mut().take(m) {
            for j in (0..n).rev() {
                let cur = chars[j];

                if cur == '.' {
                    let mut k = j;
                    while k > 0 {
                        k -= 1;
                        let prev = chars[k];
                        if prev == '*' {
                            break;
                        } else if prev == '#' {
                            chars[k] = '.';
                            chars[j] = '#';
                            break;
                        }
                    }
                }
            }
        }
    }

    fn rotate(r#box: Vec<Vec<char>>) -> Vec<Vec<char>> {
        let m = r#box.len();
        let n = r#box[0].len();

        let mut rotated_box = vec![vec![' '; m]; n];

        for (i, chars) in r#box.into_iter().enumerate() {
            for (j, ch) in chars.into_iter().enumerate() {
                rotated_box[j][m - 1 - i] = ch;
            }
        }

        rotated_box
    }
}
