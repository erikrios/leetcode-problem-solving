fn main() {
    println!("{}", Solution::maximum_score(2, 4, 6));
    println!("{}", Solution::maximum_score(4, 4, 6));
    println!("{}", Solution::maximum_score(1, 8, 8));
}

struct Solution;

impl Solution {
    pub fn maximum_score(a: i32, b: i32, c: i32) -> i32 {
        let mut a = a;
        let mut b = b;
        let mut c = c;

        let mut zero_counter = 0;
        let mut score = 0;

        while zero_counter <= 1 {
            Solution::get_two_max_and_subtract(&mut a, &mut b, &mut c);

            zero_counter = 0;

            if a == 0 {
                zero_counter += 1;
            }
            if b == 0 {
                zero_counter += 1;
            }
            if c == 0 {
                zero_counter += 1;
            }

            score += 1;
        }

        score
    }

    fn get_two_max_and_subtract(a: &mut i32, b: &mut i32, c: &mut i32) {
        let mut refs = [&mut *a, &mut *b, &mut *c];
        refs.sort_unstable_by(|x, y| y.cmp(x));

        *refs[0] -= 1;
        *refs[1] -= 1;
    }
}
