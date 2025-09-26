fn main() {
    println!(
        "{}",
        Solution::compare_version("1.2".to_string(), "1.10".to_string())
    );
    println!(
        "{}",
        Solution::compare_version("1.01".to_string(), "1.001".to_string())
    );
    println!(
        "{}",
        Solution::compare_version("1.0".to_string(), "1.0.0.0".to_string())
    );
    println!(
        "{}",
        Solution::compare_version("0.1".to_string(), "1.0".to_string())
    );
}

struct Solution;

impl Solution {
    pub fn compare_version(version1: String, version2: String) -> i32 {
        let m = version1.len();
        let n = version2.len();

        let mut version1_revisions = Vec::new();
        let mut revision = String::new();
        let mut is_leading_zero = true;
        for (i, ch) in version1.chars().enumerate() {
            if i == m - 1 {
                if !(is_leading_zero && ch == '0') {
                    revision.push(ch);
                } else {
                    revision.push('0');
                }
                version1_revisions.push(revision);
                break;
            }

            if ch == '.' {
                if !revision.is_empty() {
                    version1_revisions.push(revision);
                } else {
                    version1_revisions.push("0".to_string());
                }
                is_leading_zero = true;
                revision = String::new();
                continue;
            }

            if ch == '0' && is_leading_zero {
                continue;
            }

            is_leading_zero = false;
            revision.push(ch);
        }

        let mut version2_revisions = Vec::new();
        let mut revision = String::new();
        for (i, ch) in version2.chars().enumerate() {
            if i == n - 1 {
                if !(is_leading_zero && ch == '0') {
                    revision.push(ch);
                } else {
                    revision.push('0');
                }
                version2_revisions.push(revision);
                break;
            }
            if ch == '.' {
                if !revision.is_empty() {
                    version2_revisions.push(revision);
                } else {
                    version2_revisions.push("0".to_string());
                }
                is_leading_zero = true;
                revision = String::new();
                continue;
            }

            if ch == '0' && is_leading_zero {
                continue;
            }

            is_leading_zero = false;
            revision.push(ch);
        }

        let n = if version2_revisions.len() > version1_revisions.len() {
            version2_revisions.len()
        } else {
            version1_revisions.len()
        };

        for i in 0..n {
            let ver1_num: i32 = if i < version1_revisions.len() {
                version1_revisions[i].parse().unwrap()
            } else {
                0
            };

            let ver2_num: i32 = if i < version2_revisions.len() {
                version2_revisions[i].parse().unwrap()
            } else {
                0
            };

            if ver1_num < ver2_num {
                return -1;
            } else if ver1_num > ver2_num {
                return 1;
            }
        }

        0
    }
}
