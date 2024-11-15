fn main() {
    println!(
        "{}",
        Solution::min_groups(vec![
            vec![5, 10],
            vec![6, 8],
            vec![1, 5],
            vec![2, 3],
            vec![1, 10]
        ])
    );
    println!(
        "{}",
        Solution::min_groups(vec![vec![1, 3], vec![5, 6], vec![8, 10], vec![11, 13]])
    );
}

struct Solution;

impl Solution {
    pub fn min_groups(intervals: Vec<Vec<i32>>) -> i32 {
        let n = intervals.len();
        let mut starts = Vec::with_capacity(n / 2);
        let mut ends = Vec::with_capacity(n / 2);

        for interval in intervals {
            starts.push(interval[0]);
            ends.push(interval[1]);
        }

        starts.sort();
        ends.sort();

        let mut res = 0;
        let mut groups = 0;
        let mut i = 0;
        let mut j = 0;

        while i < n {
            if starts[i] <= ends[j] {
                i += 1;
                groups += 1;
            } else {
                j += 1;
                groups -= 1;
            }
            res = res.max(groups);
        }

        res
    }
}
