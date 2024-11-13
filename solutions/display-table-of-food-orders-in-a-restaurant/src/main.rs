use std::collections::{BTreeMap, BTreeSet, HashMap};

fn main() {
    println!(
        "{:?}",
        Solution::display_table(vec![
            vec!["David".to_string(), "3".to_string(), "Ceviche".to_string()],
            vec![
                "Corina".to_string(),
                "10".to_string(),
                "Beef Burrito".to_string()
            ],
            vec![
                "David".to_string(),
                "3".to_string(),
                "Fried Chicken".to_string()
            ],
            vec!["Carla".to_string(), "5".to_string(), "Water".to_string()],
            vec!["Carla".to_string(), "5".to_string(), "Ceviche".to_string()],
            vec!["Rous".to_string(), "3".to_string(), "Ceviche".to_string()]
        ])
    );
    println!(
        "{:?}",
        Solution::display_table(vec![
            vec![
                "James".to_string(),
                "12".to_string(),
                "Fried Chicken".to_string()
            ],
            vec![
                "Ratesh".to_string(),
                "12".to_string(),
                "Fried Chicken".to_string()
            ],
            vec![
                "Amadeus".to_string(),
                "12".to_string(),
                "Fried Chicken".to_string()
            ],
            vec![
                "Adam".to_string(),
                "1".to_string(),
                "Canadian Waffles".to_string()
            ],
            vec![
                "Brianna".to_string(),
                "1".to_string(),
                "Canadian Waffles".to_string()
            ]
        ])
    );
    println!(
        "{:?}",
        Solution::display_table(vec![
            vec![
                "Laura".to_string(),
                "2".to_string(),
                "Bean Burrito".to_string()
            ],
            vec![
                "Jhon".to_string(),
                "2".to_string(),
                "Beef Burrito".to_string()
            ],
            vec!["Melissa".to_string(), "2".to_string(), "Soda".to_string()]
        ])
    );
}

struct Solution;

impl Solution {
    pub fn display_table(orders: Vec<Vec<String>>) -> Vec<Vec<String>> {
        let mut order_b_tree_mapper: BTreeMap<usize, HashMap<String, usize>> = BTreeMap::new();
        let mut foods_b_tree_set = BTreeSet::new();

        for order in orders {
            let table = order[1].parse().unwrap();
            let food = order[2].clone();

            foods_b_tree_set.insert(food.clone());

            order_b_tree_mapper
                .entry(table)
                .and_modify(|menus| {
                    menus
                        .entry(food.clone())
                        .and_modify(|count| *count += 1)
                        .or_insert(1);
                })
                .or_insert_with(|| {
                    let mut foods = HashMap::new();
                    foods.insert(food, 1);
                    foods
                });
        }

        let mut display_table = Vec::new();
        let mut header = vec!["Table".to_string()];
        for food in foods_b_tree_set {
            header.push(food);
        }
        display_table.push(header.clone());

        for (table, foods_mapper) in order_b_tree_mapper {
            let mut table = vec![table.to_string()];

            for food in header.iter().skip(1) {
                let count = if let Some(val) = foods_mapper.get(food) {
                    val.to_string()
                } else {
                    "0".to_string()
                };

                table.push(count);
            }

            display_table.push(table);
        }

        display_table
    }
}
