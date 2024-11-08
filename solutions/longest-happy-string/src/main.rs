fn main() {
    println!("{}", Solution::longest_diverse_string(1, 1, 7));
    println!("{}", Solution::longest_diverse_string(7, 1, 0));
    println!("{}", Solution::longest_diverse_string(2, 4, 1));
    println!("{}", Solution::longest_diverse_string(1, 0, 3));
    println!("{}", Solution::longest_diverse_string(5, 4, 3));
}

struct Solution;

impl Solution {
    pub fn longest_diverse_string(a: i32, b: i32, c: i32) -> String {
        let mut res = String::new();

        let total = a + b + c;

        let mut continuous_a = 0;
        let mut continuous_b = 0;
        let mut continuous_c = 0;

        let mut a = a;
        let mut b = b;
        let mut c = c;

        for _ in 0..total {
            if (a >= b && a >= c && continuous_a != 2)
                || (continuous_b == 2 && a > 0)
                || continuous_c == 2 && a > 0
            {
                res.push('a');
                a -= 1;
                continuous_a += 1;
                continuous_b = 0;
                continuous_c = 0;
            } else if (b >= a && b >= c && continuous_b != 2)
                || (continuous_a == 2 && b > 0)
                || continuous_c == 2 && b > 0
            {
                res.push('b');
                b -= 1;
                continuous_a = 0;
                continuous_b += 1;
                continuous_c = 0;
            } else if (c >= a && c >= b && continuous_c != 2)
                || (continuous_a == 2 && c > 0)
                || continuous_b == 2 && c > 0
            {
                res.push('c');
                c -= 1;
                continuous_a = 0;
                continuous_b = 0;
                continuous_c += 1;
            }
        }

        res
    }
}
