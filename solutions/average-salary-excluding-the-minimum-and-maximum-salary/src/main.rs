fn main() {
    println!("{}", Solution::average(vec![4000, 3000, 1000, 2000]));
    println!("{}", Solution::average(vec![1000, 2000, 3000]));
}

struct Solution;

impl Solution {
    pub fn average(salary: Vec<i32>) -> f64 {
        let n = salary.len();
        let mut min = salary[0];
        let mut max = salary[0];
        let mut total = salary[0];

        for s in salary.into_iter().skip(1) {
            total += s;
            min = min.min(s);
            max = max.max(s);
        }

        (total - (min + max)) as f64 / (n - 2) as f64
    }
}
