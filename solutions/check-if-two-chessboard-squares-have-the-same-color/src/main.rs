fn main() {
    println!(
        "{}",
        Solution::check_two_chessboards("a1".to_string(), "c3".to_string())
    );
    println!(
        "{}",
        Solution::check_two_chessboards("a1".to_string(), "h3".to_string())
    );
    println!(
        "{}",
        Solution::check_two_chessboards("d1".to_string(), "h4".to_string())
    );
}

struct Solution;

impl Solution {
    pub fn check_two_chessboards(coordinate1: String, coordinate2: String) -> bool {
        let coordinate1 = coordinate1.as_bytes();
        let coordinate2 = coordinate2.as_bytes();

        !((((coordinate1[0] - b'a') % 2 == 0) ^ ((coordinate1[1] - b'1') % 2 == 0))
            ^ (((coordinate2[0] - b'a') % 2 == 0) ^ ((coordinate2[1] - b'1') % 2 == 0)))
    }
}
