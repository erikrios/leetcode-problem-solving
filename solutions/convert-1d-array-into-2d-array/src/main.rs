fn main() {
    println!("{:?}", Solution::construct2_d_array(vec![1, 2, 3, 4], 2, 2));
    println!("{:?}", Solution::construct2_d_array(vec![1, 2, 3], 1, 3));
    println!("{:?}", Solution::construct2_d_array(vec![1, 2], 1, 1));
}

struct Solution;

impl Solution {
    pub fn construct2_d_array(original: Vec<i32>, m: i32, n: i32) -> Vec<Vec<i32>> {
        let l = original.len();
        if m * n != l as i32 {
            return vec![];
        }

        let mut results = Vec::with_capacity(m as usize);

        for i in 0..m as usize {
            let mut res = Vec::with_capacity(n as usize);

            for j in 0..n as usize {
                let num = original[j + (i * n as usize)];
                res.push(num);
            }

            results.push(res);
        }

        results
    }
}
