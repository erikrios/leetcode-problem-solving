use std::collections::HashMap;

fn main() {
    println!("{:?}", Solution::find_relative_ranks(vec![5,4,3,2,1]));
    println!("{:?}", Solution::find_relative_ranks(vec![10,3,8,9,4]));
}

struct Solution;

impl Solution {
    pub fn find_relative_ranks(score: Vec<i32>) -> Vec<String> {
        let mut sorted_score = score.clone();
        sorted_score.sort_by(|a, b| b.cmp(a));

        let mut mapper: HashMap<i32, usize> = HashMap::new();


        for i in 0..sorted_score.len() {
            mapper.insert(sorted_score[i], i);
        }

        let mut results: Vec<String> = Vec::new();
        for i in 0..score.len() {
            let val = score[i];
            match mapper.get(&val) {
                None => continue,
                Some(index) => {
                    match index {
                        0 => results.push(String::from("Gold Medal")),
                        1 => results.push(String::from("Silver Medal")),
                        2 => results.push(String::from("Bronze Medal")),
                        _ => results.push((index+1).to_string())
                    }
                }
            }
            
        }

        results
    }
}
