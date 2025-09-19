use std::collections::BTreeMap;

fn main() {
    println!(
        "{:?}",
        Solution::spellchecker(
            vec![
                "KiTe".to_string(),
                "kite".to_string(),
                "hare".to_string(),
                "Hare".to_string()
            ],
            vec![
                "kite".to_string(),
                "Kite".to_string(),
                "KiTe".to_string(),
                "Hare".to_string(),
                "HARE".to_string(),
                "Hear".to_string(),
                "hear".to_string(),
                "keti".to_string(),
                "keet".to_string(),
                "keto".to_string(),
            ],
        ),
    );

    println!(
        "{:?}",
        Solution::spellchecker(vec!["yellow".to_string()], vec!["YellOw".to_string()])
    );
}

struct Solution;

impl Solution {
    pub fn spellchecker(wordlist: Vec<String>, queries: Vec<String>) -> Vec<String> {
        let mut match_mapper = BTreeMap::new();
        let mut caps_mapper = BTreeMap::new();
        let mut vow_err_mapper = BTreeMap::new();

        for (i, word) in wordlist.iter().enumerate() {
            match_mapper.entry(word).or_insert(i);

            let caps = word.to_lowercase();
            caps_mapper.entry(caps).or_insert(i);

            let mut vow_err = String::with_capacity(word.len());

            for &ch in word.as_bytes() {
                if ch == b'a'
                    || ch == b'i'
                    || ch == b'u'
                    || ch == b'e'
                    || ch == b'o'
                    || ch == b'A'
                    || ch == b'I'
                    || ch == b'U'
                    || ch == b'E'
                    || ch == b'O'
                {
                    vow_err.push(' ');
                } else {
                    vow_err.push(ch as char);
                }
            }

            vow_err = vow_err.to_lowercase();
            vow_err_mapper.entry(vow_err).or_insert(i);
        }

        let mut results = Vec::with_capacity(queries.len());

        for query in queries {
            if match_mapper.contains_key(&query) {
                let &i = match_mapper.get(&query).unwrap();
                results.push(wordlist[i].clone());
                continue;
            }

            let caps_query = query.to_lowercase();
            if caps_mapper.contains_key(&caps_query) {
                let &i = caps_mapper.get(&caps_query).unwrap();
                results.push(wordlist[i].clone());
                continue;
            }

            let mut vow_err_query = String::with_capacity(query.len());

            for &ch in query.as_bytes() {
                if ch == b'a'
                    || ch == b'i'
                    || ch == b'u'
                    || ch == b'e'
                    || ch == b'o'
                    || ch == b'A'
                    || ch == b'I'
                    || ch == b'U'
                    || ch == b'E'
                    || ch == b'O'
                {
                    vow_err_query.push(' ');
                } else {
                    vow_err_query.push(ch as char);
                }
            }

            vow_err_query = vow_err_query.to_lowercase();
            if vow_err_mapper.contains_key(&vow_err_query) {
                let &i = vow_err_mapper.get(&vow_err_query).unwrap();
                results.push(wordlist[i].clone());
                continue;
            }

            results.push("".to_string());
        }

        results
    }
}
