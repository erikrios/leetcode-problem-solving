use std::collections::{BTreeSet, HashMap};

fn main() {
    println!(
        "{}",
        Solution::smallest_equivalent_string(
            "parker".to_string(),
            "morris".to_string(),
            "parser".to_string()
        )
    );
    println!(
        "{}",
        Solution::smallest_equivalent_string(
            "hello".to_string(),
            "world".to_string(),
            "hold".to_string()
        )
    );
    println!(
        "{}",
        Solution::smallest_equivalent_string(
            "leetcode".to_string(),
            "programs".to_string(),
            "sourcecode".to_string()
        )
    );
    println!(
        "{}",
        Solution::smallest_equivalent_string(
            "dccaccbdafgeabeeghbigbhicggfbhiccfgbechdicbhdcgahi".to_string(),
            "igfcigeciahdafgegfbeddhgbacaeehcdiehiifgbhhehhccde".to_string(),
            "sanfbzzwblekirguignnfkpzgqjmjmfokrdfuqbgyflpsfpzbo".to_string(),
        )
    );

    println!(
        "{}",
        Solution::smallest_equivalent_string(
            "adkbbjigibahbjjgdkkiighagijfdfjkcdaakdkbjcidgjjfga".to_string(),
            "hbfahdikgchbkigebgjghdhadaikhccjejafkaibdgffichcbb".to_string(),
            "hotutsrhanyvpzusrnsxbncpqhnxrvgmbrpcbhheqotadyzfyl".to_string(),
        )
    );
}

struct Solution;

impl Solution {
    pub fn smallest_equivalent_string(s1: String, s2: String, base_str: String) -> String {
        let n = s1.len();

        let s1: Vec<char> = s1.chars().collect();
        let s2: Vec<char> = s2.chars().collect();
        let base_str: Vec<char> = base_str.chars().collect();

        let mut char_groups = HashMap::new();
        let mut group_locs = HashMap::new();

        for i in 0..n {
            let s1_char = s1[i];
            let s2_char = s2[i];

            if char_groups.contains_key(&s1_char) {
                let &group = char_groups.get(&s1_char).unwrap();

                let mut need_updates = Vec::new();
                if let Some(&old_group) = char_groups.get(&s2_char) {
                    if group != old_group {
                        let old_tree_set = group_locs.get(&old_group).unwrap();
                        for &old_set in old_tree_set {
                            need_updates.push(old_set);
                        }
                        group_locs.remove(&old_group);
                    }
                }

                char_groups.insert(s1_char, group);
                char_groups.insert(s2_char, group);

                let tree_set: &mut BTreeSet<char> = group_locs.get_mut(&group).unwrap();
                tree_set.insert(s1_char);
                tree_set.insert(s2_char);
                for ch in need_updates {
                    char_groups.insert(ch, group);
                    tree_set.insert(ch);
                }
            } else if char_groups.contains_key(&s2_char) {
                let &group = char_groups.get(&s2_char).unwrap();

                let mut need_updates = Vec::new();
                if let Some(&old_group) = char_groups.get(&s1_char) {
                    if group != old_group {
                        let old_tree_set = group_locs.get(&old_group).unwrap();
                        for &old_set in old_tree_set {
                            need_updates.push(old_set);
                        }
                        group_locs.remove(&old_group);
                    }
                }

                char_groups.insert(s1_char, group);
                char_groups.insert(s2_char, group);

                let tree_set: &mut BTreeSet<char> = group_locs.get_mut(&group).unwrap();
                tree_set.insert(s1_char);
                tree_set.insert(s2_char);
                for ch in need_updates {
                    char_groups.insert(ch, group);
                    tree_set.insert(ch);
                }
            } else {
                char_groups.insert(s1_char, i);
                char_groups.insert(s2_char, i);
                let mut tree_set = BTreeSet::new();
                tree_set.insert(s1_char);
                tree_set.insert(s2_char);
                group_locs.insert(i, tree_set);
            }
        }

        let m = base_str.len();

        let mut result = String::with_capacity(m);
        for &ch in base_str.iter().take(m) {
            let smallest_ch = if let Some(group) = char_groups.get(&ch) {
                let tree_set = group_locs.get(group).unwrap();
                *tree_set.iter().next().unwrap()
            } else {
                ch
            };
            result.push(smallest_ch);
        }

        result
    }
}
