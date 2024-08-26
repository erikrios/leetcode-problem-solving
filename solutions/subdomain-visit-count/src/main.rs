use std::collections::HashMap;

fn main() {
    println!(
        "{:?}",
        Solution::subdomain_visits(vec!["9001 discuss.leetcode.com".to_string()])
    );
    println!(
        "{:?}",
        Solution::subdomain_visits(vec![
            "900 google.mail.com".to_string(),
            "50 yahoo.com".to_string(),
            "1 intel.mail.com".to_string(),
            "5 wiki.org".to_string()
        ])
    );
}

struct Solution;

impl Solution {
    pub fn subdomain_visits(cpdomains: Vec<String>) -> Vec<String> {
        let mut sub_domain_visit_counter = HashMap::new();

        for cpdomain in cpdomains {
            let cpdomain: Vec<String> = cpdomain.split(' ').map(String::from).collect();
            let rep: i32 = cpdomain[0].parse().unwrap();
            let sub_domain = &cpdomain[1];

            let mut sub_domain: Vec<&str> = sub_domain.split('.').collect();

            while !sub_domain.is_empty() {
                let domain = sub_domain.join(".");
                sub_domain_visit_counter
                    .entry(domain.clone())
                    .and_modify(|e| *e += rep)
                    .or_insert(rep);

                sub_domain.remove(0);
            }
        }

        sub_domain_visit_counter
            .into_iter()
            .map(|(key, value)| format!("{} {}", value, key))
            .collect()
    }
}
