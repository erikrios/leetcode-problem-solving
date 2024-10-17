use std::{cell::RefCell, rc::Rc};

fn main() {
    let node = Some(Rc::new(RefCell::new(TreeNode {
        val: 1,
        left: Some(Rc::new(RefCell::new(TreeNode {
            val: 2,
            left: None,
            right: Some(Rc::new(RefCell::new(TreeNode::new(5)))),
        }))),
        right: Some(Rc::new(RefCell::new(TreeNode::new(3)))),
    })));

    let paths = Solution::binary_tree_paths(node.clone());
    println!("{:?}", paths);

    let single_node = Some(Rc::new(RefCell::new(TreeNode::new(1))));
    let single_node_paths = Solution::binary_tree_paths(single_node);
    println!("{:?}", single_node_paths);
}

#[derive(Debug, PartialEq, Eq)]
pub struct TreeNode {
    pub val: i32,
    pub left: Option<Rc<RefCell<TreeNode>>>,
    pub right: Option<Rc<RefCell<TreeNode>>>,
}

impl TreeNode {
    pub fn new(val: i32) -> Self {
        TreeNode {
            val,
            left: None,
            right: None,
        }
    }
}

struct Solution;

impl Solution {
    pub fn binary_tree_paths(root: Option<Rc<RefCell<TreeNode>>>) -> Vec<String> {
        Self::dfs(root, String::new())
    }

    fn dfs(root: Option<Rc<RefCell<TreeNode>>>, mut res: String) -> Vec<String> {
        if let Some(node) = root {
            let node_ref = node.borrow();
            let val = node_ref.val;
            if Self::is_leaf(node.clone()) {
                res += &val.to_string();
                return vec![res];
            }

            res += &format!("{}->", val);

            let mut results = Vec::new();

            if let Some(ref left_node) = node_ref.left {
                results.append(&mut Self::dfs(Some(left_node.clone()), res.clone()))
            }

            if let Some(ref right_node) = node_ref.right {
                results.append(&mut Self::dfs(Some(right_node.clone()), res.clone()))
            }

            results
        } else {
            vec![]
        }
    }

    fn is_leaf(node: Rc<RefCell<TreeNode>>) -> bool {
        let node_ref = node.borrow();
        node_ref.left.is_none() && node_ref.right.is_none()
    }
}
