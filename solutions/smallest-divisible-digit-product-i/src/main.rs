fn main() {
    println!("{}", Solution::smallest_number(10, 2));
    println!("{}", Solution::smallest_number(15, 3));
}

struct Solution;

impl Solution {
    pub fn smallest_number(n: i32, t: i32) -> i32 {
        let mut smallest = n;

        loop {
            if Self::digit_product(smallest) % t == 0 {
                break;
            }
            smallest += 1;
        }

        smallest
    }

    fn digit_product(n: i32) -> i32 {
        let mut n = n;
        let mut product = 1;

        while n > 0 {
            let digit = n % 10;
            product *= digit;
            n /= 10;
        }

        product
    }
}
