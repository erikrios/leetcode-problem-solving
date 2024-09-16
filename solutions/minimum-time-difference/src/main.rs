fn main() {
    println!(
        "{}",
        Solution::find_min_difference(vec!["23:59".to_string(), "00:00".to_string()])
    );
    println!(
        "{}",
        Solution::find_min_difference(vec![
            "00:00".to_string(),
            "23:59".to_string(),
            "00:00".to_string()
        ])
    );
    println!(
        "{}",
        Solution::find_min_difference(vec![
            "00:00".to_string(),
            "04:00".to_string(),
            "22:00".to_string()
        ])
    );
}

struct Solution;

impl Solution {
    pub fn find_min_difference(time_points: Vec<String>) -> i32 {
        let mut b_tree_set = std::collections::BTreeSet::new();

        for time_point in time_points {
            let minute = Self::hours_to_minutes(time_point);

            if b_tree_set.contains(&minute) {
                return 0;
            }

            b_tree_set.insert(minute);
        }

        let minutes: Vec<usize> = b_tree_set.into_iter().collect();

        let mut min_diff = 24 * 60;
        for i in 0..minutes.len() {
            let (cur_minute, next_minute) = if i < minutes.len() - 1 {
                (minutes[i], minutes[i + 1])
            } else {
                (minutes[i], minutes[0])
            };

            let greater;
            let smaller;
            if cur_minute > next_minute {
                greater = cur_minute;
                smaller = next_minute;
            } else {
                greater = next_minute;
                smaller = cur_minute;
            }

            let min_a = greater - smaller;
            let min_b = 24 * 60 - greater + smaller;

            let min = min_a.min(min_b);

            min_diff = min_diff.min(min);
        }

        min_diff as i32
    }

    fn hours_to_minutes(hours: String) -> usize {
        let hours: Vec<&str> = hours.split(':').collect();
        let (hour, minute) = (hours[0], hours[1]);
        let hour = if hour.starts_with('0') {
            &hour[1..]
        } else {
            hour
        };
        let minute = if minute.starts_with('0') {
            &minute[1..]
        } else {
            minute
        };

        let hour: usize = hour.parse().unwrap();
        let minute: usize = minute.parse().unwrap();

        hour * 60 + minute
    }
}
