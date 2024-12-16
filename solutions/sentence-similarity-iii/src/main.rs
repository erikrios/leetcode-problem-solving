fn main() {
    println!(
        "{}",
        Solution::are_sentences_similar("My name is Haley".to_string(), "My Haley".to_string())
    );
    println!(
        "{}",
        Solution::are_sentences_similar("of".to_string(), "A lot of words".to_string())
    );
    println!(
        "{}",
        Solution::are_sentences_similar("Eating right now".to_string(), "Eating".to_string())
    );
    println!(
        "{}",
        Solution::are_sentences_similar("of test".to_string(), "of words".to_string())
    );
    println!(
        "{}",
        Solution::are_sentences_similar("of test wo".to_string(), "of words".to_string())
    );
    println!(
        "{}",
        Solution::are_sentences_similar("of words wo".to_string(), "of words".to_string())
    );
    println!(
        "{}",
        Solution::are_sentences_similar("of test wo".to_string(), "of test wo woso".to_string())
    );
    println!(
        "{}",
        Solution::are_sentences_similar("Frog cool".to_string(), "Frogs cool".to_string())
    );
    println!(
        "{}",
        Solution::are_sentences_similar("A".to_string(), "a A b A".to_string())
    );
}

struct Solution;

impl Solution {
    pub fn are_sentences_similar(sentence1: String, sentence2: String) -> bool {
        if sentence1 == sentence2 {
            return true;
        }

        let (mut sentence1, mut sentence2): (Vec<&str>, Vec<&str>) =
            if sentence1.len() > sentence2.len() {
                (
                    sentence1.split(' ').collect(),
                    sentence2.split(' ').collect(),
                )
            } else {
                (
                    sentence2.split(' ').collect(),
                    sentence1.split(' ').collect(),
                )
            };

        let mut flipper = true;
        let mut last_state = true;
        while !sentence1.is_empty() && !sentence2.is_empty() {
            let mut state = false;
            if flipper {
                if sentence1[0] == sentence2[0] {
                    sentence1.remove(0);
                    sentence2.remove(0);
                    state = true;
                }
            } else if sentence1[sentence1.len() - 1] == sentence2[sentence2.len() - 1] {
                sentence1.pop();
                sentence2.pop();
                state = true;
            }

            if !last_state && !state {
                return false;
            }

            last_state = state;

            flipper = !flipper
        }

        true
    }
}
