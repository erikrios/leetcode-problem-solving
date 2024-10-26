use std::collections::HashSet;

fn main() {
    println!(
        "{:?}",
        Solution::remove_subfolders(vec![
            "/a".to_string(),
            "/a/b".to_string(),
            "/c/d".to_string(),
            "/c/d/e".to_string(),
            "/c/f".to_string()
        ])
    );
    println!(
        "{:?}",
        Solution::remove_subfolders(vec![
            "/a".to_string(),
            "/a/b/c".to_string(),
            "/a/b/d".to_string()
        ])
    );
    println!(
        "{:?}",
        Solution::remove_subfolders(vec![
            "/a/b/c".to_string(),
            "/a/b/ca".to_string(),
            "/a/b/d".to_string()
        ])
    );
}

struct Solution;

impl Solution {
    pub fn remove_subfolders(folder: Vec<String>) -> Vec<String> {
        let mut folder = folder;
        folder.sort();
        let mut folder_tracker = HashSet::new();

        'outer: for sub_folder in folder.iter() {
            for i in 0..sub_folder.len() {
                if i == sub_folder.len() - 1 || sub_folder.as_bytes()[i + 1] == b'/' {
                    let path = &sub_folder[0..i + 1];
                    if folder_tracker.contains(path) {
                        continue 'outer;
                    }
                }
            }

            folder_tracker.insert(sub_folder.as_str());
        }

        folder_tracker
            .into_iter()
            .map(|sub| sub.to_string())
            .collect()
    }
}
