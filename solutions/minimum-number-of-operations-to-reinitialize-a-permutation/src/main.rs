fn main() {
    println!("{}", Solution::reinitialize_permutation(2));
    println!("{}", Solution::reinitialize_permutation(4));
    println!("{}", Solution::reinitialize_permutation(6));
    println!("{}", Solution::reinitialize_permutation(1000));
}

struct Solution;

impl Solution {
    pub fn reinitialize_permutation(n: i32) -> i32 {
        let mut number_of_operations = 0;
        let mut arr = Vec::with_capacity(n as usize);
        let mut tmp_arr = Vec::with_capacity(n as usize);
        for i in 0..n as usize {
            arr.push(i);
            tmp_arr.push(i);
        }

        loop {
            number_of_operations += 1;
            let mut is_initial = true;

            for i in 0..n as usize {
                tmp_arr[i] = if i & 1 != 1 {
                    arr[i / 2]
                } else {
                    arr[n as usize / 2 + (i - 1) / 2]
                };

                if tmp_arr[i] != i {
                    is_initial = false;
                }
            }

            arr = tmp_arr.clone();

            if is_initial {
                return number_of_operations;
            }
        }
    }
}
