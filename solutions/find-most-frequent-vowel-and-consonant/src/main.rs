use find_most_frequent_vowel_and_consonant::Solution;

fn main() {
    println!(
        "{}",
        Solution::max_freq_sum_two_loops("successes".to_string())
    );
    println!("{}", Solution::max_freq_sum_one_loop("aeiaeia".to_string()));
}
