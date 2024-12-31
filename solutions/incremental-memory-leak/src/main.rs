fn main() {
    println!("{:?}", Solution::mem_leak(2, 2));
    println!("{:?}", Solution::mem_leak(8, 11));
}

struct Solution;

impl Solution {
    pub fn mem_leak(memory1: i32, memory2: i32) -> Vec<i32> {
        let mut memory1 = memory1;
        let mut memory2 = memory2;
        let mut i = 0;

        loop {
            i += 1;

            if memory1 >= memory2 {
                if memory1 < i {
                    break;
                }
                memory1 -= i;
            } else {
                if memory2 < i {
                    break;
                }
                memory2 -= i;
            }
        }

        vec![i, memory1, memory2]
    }
}
