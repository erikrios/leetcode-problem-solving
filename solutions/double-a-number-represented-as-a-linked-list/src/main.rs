fn main() {
    let node = Some(Box::new(ListNode {
        val: 1,
        next: Some(Box::new(ListNode {
            val: 8,
            next: Some(Box::new(ListNode { val: 9, next: None })),
        })),
    }));
    println!("{:?}", Solution::double_it(node));

    let node = Some(Box::new(ListNode {
        val: 9,
        next: Some(Box::new(ListNode {
            val: 9,
            next: Some(Box::new(ListNode { val: 9, next: None })),
        })),
    }));
    println!("{:?}", Solution::double_it(node));
}

#[derive(PartialEq, Eq, Clone, Debug)]
pub struct ListNode {
    pub val: i32,
    pub next: Option<Box<ListNode>>,
}

impl ListNode {
    #[inline]
    fn new(val: i32) -> Self {
        ListNode { val, next: None }
    }
}

struct Solution;

impl Solution {
    pub fn double_it(head: Option<Box<ListNode>>) -> Option<Box<ListNode>> {
        let mut head = if let Some(node) = head {
            if node.val > 4 {
                let mut new_head = ListNode::new(0);
                new_head.next = Some(node);
                Some(Box::new(new_head))
            } else {
                Some(node)
            }
        } else {
            head
        };

        let mut cur = head.as_mut();

        while let Some(node) = cur {
            node.val = (node.val * 2) % 10;
            if let Some(next_node) = &node.next.as_mut() {
                if next_node.val > 4 {
                    node.val += 1;
                }
            }
            cur = node.next.as_mut();
        }

        head
    }
}
