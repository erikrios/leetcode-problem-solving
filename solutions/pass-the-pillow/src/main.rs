fn main() {
    println!("{}", Solution::pass_the_pillow(4, 1));
    println!("{}", Solution::pass_the_pillow(4, 2));
    println!("{}", Solution::pass_the_pillow(4, 3));
    println!("{}", Solution::pass_the_pillow(4, 4));
    println!("{}", Solution::pass_the_pillow(4, 5));
    println!("{}", Solution::pass_the_pillow(4, 6));
    println!("{}", Solution::pass_the_pillow(4, 7));
    println!("{}", Solution::pass_the_pillow(4, 8));
    println!("{}", Solution::pass_the_pillow(4, 9));
    println!("{}", Solution::pass_the_pillow(3, 2));
}

struct Solution;

impl Solution {
    pub fn pass_the_pillow(n: i32, time: i32) -> i32 {
        let mut div_res = time / (n - 1);
        let mut mod_res = time % (n - 1);
        if mod_res == 0 {
            div_res -= 1;
            mod_res = n - 1;
        }
        let is_from_front = div_res % 2 == 0;
        if is_from_front {
            1 + mod_res
        } else {
            n - mod_res
        }
    }
}
