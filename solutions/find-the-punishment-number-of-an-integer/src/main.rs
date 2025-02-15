fn main() {
    println!("{}", Solution::punishment_number(10));
    println!("{}", Solution::punishment_number(37));
}

struct Solution;

impl Solution {
    pub fn punishment_number(n: i32) -> i32 {
        let mut sum = 0;

        for i in 1..=n as usize {
            let square = i * i;
            let digits = Self::extract_digits(square);
            if Self::is_punishment_number(i, &digits, 0, &mut vec![]) {
                sum += square;
            }
        }

        sum as i32
    }

    fn extract_digits(n: usize) -> Vec<usize> {
        let mut n = n;
        let mut digits = Vec::new();

        while n > 0 {
            let digit = n % 10;
            digits.insert(0, digit);
            n /= 10;
        }

        digits
    }

    fn is_punishment_number(
        n: usize,
        n_digits: &[usize],
        start: usize,
        partitions: &mut Vec<usize>,
    ) -> bool {
        let digit_len = n_digits.len();

        if start == digit_len {
            if partitions.iter().sum::<usize>() == n {
                return true;
            }

            return false;
        }

        let mut temp = 0;
        for i in start..digit_len {
            temp *= 10;
            let digit = n_digits[i];
            temp += digit;
            partitions.push(temp);

            if Self::is_punishment_number(n, n_digits, i + 1, partitions) {
                return true;
            }

            partitions.pop();
        }

        false
    }
}
