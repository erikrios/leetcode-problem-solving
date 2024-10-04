fn main() {
    println!("{}", Solution::divide_players(vec![3, 2, 5, 1, 3, 4]));
    println!("{}", Solution::divide_players(vec![3, 4]));
    println!("{}", Solution::divide_players(vec![1, 1, 2, 3]));
}

struct Solution;

impl Solution {
    pub fn divide_players(skill: Vec<i32>) -> i64 {
        let n = skill.len();
        let mut skill = skill;
        skill.sort();

        let mut chemistry = 0;
        let mut cur_total_skill = 0;
        for i in 0..n / 2 {
            let skill_player_one = skill[i];
            let skill_player_two = skill[n - 1 - i];
            let sum = skill_player_one + skill_player_two;
            if i == 0 {
                cur_total_skill = sum;
            }
            if sum != cur_total_skill {
                return -1;
            }

            let product = skill_player_one * skill_player_two;
            chemistry += product as i64;
        }

        chemistry
    }
}
