fn main() {
    let mut comb = CombinationIterator::new("abc".to_string(), 2);
    println!("{}", comb.next());
    println!("{}", comb.has_next());
    println!("{}", comb.next());
    println!("{}", comb.has_next());
    println!("{}", comb.next());
    println!("{}", comb.has_next());
}

struct CombinationIterator {
    i: usize,
    res: Vec<String>,
}

impl CombinationIterator {
    fn new(characters: String, combination_length: i32) -> Self {
        let mut combination = Vec::new();
        let mut res = Vec::new();
        Solution::combination(
            characters.as_bytes(),
            combination_length as usize,
            0,
            &mut combination,
            &mut res,
        );
        Self { i: 0, res }
    }

    fn next(&mut self) -> String {
        let res = self.res[self.i].clone();
        self.i += 1;
        res
    }

    fn has_next(&self) -> bool {
        self.i < self.res.len()
    }
}

struct Solution;

impl Solution {
    fn combination(
        words: &[u8],
        n: usize,
        i: usize,
        combination: &mut Vec<u8>,
        res: &mut Vec<String>,
    ) {
        if combination.len() == n {
            res.push(String::from_utf8(combination.to_vec()).unwrap());
            return;
        }

        for j in i..words.len() {
            let ch = words[j];
            combination.push(ch);
            Self::combination(words, n, j + 1, combination, res);
            combination.pop();
        }
    }
}
