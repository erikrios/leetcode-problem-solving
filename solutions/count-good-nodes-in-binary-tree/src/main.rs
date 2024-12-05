use std::{cell::RefCell, collections::BTreeMap, rc::Rc};

fn main() {
    let root = Rc::new(RefCell::new(TreeNode::new(3)));
    let node1 = Rc::new(RefCell::new(TreeNode::new(1)));
    let node4 = Rc::new(RefCell::new(TreeNode::new(4)));
    let node3 = Rc::new(RefCell::new(TreeNode::new(3)));
    let node1_2 = Rc::new(RefCell::new(TreeNode::new(1)));
    let node5 = Rc::new(RefCell::new(TreeNode::new(5)));

    root.borrow_mut().left = Some(node1.clone());
    root.borrow_mut().right = Some(node4.clone());

    node1.borrow_mut().left = Some(node3.clone());

    node4.borrow_mut().left = Some(node1_2.clone());
    node4.borrow_mut().right = Some(node5.clone());

    println!("{}", Solution::good_nodes(Some(root)));

    let root = Rc::new(RefCell::new(TreeNode::new(3)));
    let node3 = Rc::new(RefCell::new(TreeNode::new(3)));
    let node4 = Rc::new(RefCell::new(TreeNode::new(4)));
    let node2 = Rc::new(RefCell::new(TreeNode::new(2)));

    // Build the tree
    root.borrow_mut().left = Some(node3.clone());

    node3.borrow_mut().left = Some(node4.clone());
    node3.borrow_mut().right = Some(node2.clone());

    println!("{}", Solution::good_nodes(Some(root)));

    let root = Rc::new(RefCell::new(TreeNode::new(1)));

    println!("{}", Solution::good_nodes(Some(root)));
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
    pub fn good_nodes(root: Option<Rc<RefCell<TreeNode>>>) -> i32 {
        let mut mapper = BTreeMap::new();
        Self::count(root, &mut mapper)
    }

    fn count(root: Option<Rc<RefCell<TreeNode>>>, mapper: &mut BTreeMap<i32, usize>) -> i32 {
        if let Some(node) = root {
            let val = node.borrow().val;
            let is_good = if let Some(entry) = mapper.last_entry() {
                val >= *entry.key()
            } else {
                true
            };

            mapper
                .entry(val)
                .and_modify(|count| *count += 1)
                .or_insert(1);

            let mut res = Self::count(node.borrow().left.clone(), mapper)
                + Self::count(node.borrow().right.clone(), mapper);

            mapper.entry(val).and_modify(|count| *count -= 1);

            if let Some(&0) = mapper.get(&val) {
                mapper.remove(&val);
            }

            if is_good {
                res += 1;
            }

            res
        } else {
            0
        }
    }
}
