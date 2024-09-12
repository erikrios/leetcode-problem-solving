use std::collections::HashSet;

fn main() {
    let node = Some(Box::new(ListNode {
        val: 1,
        next: Some(Box::new(ListNode {
            val: 2,
            next: Some(Box::new(ListNode {
                val: 3,
                next: Some(Box::new(ListNode {
                    val: 4,
                    next: Some(Box::new(ListNode { val: 5, next: None })),
                })),
            })),
        })),
    }));
    println!("{:?}", Solution::modified_list(vec![1, 2, 3], node));

    let node = Some(Box::new(ListNode {
        val: 1,
        next: Some(Box::new(ListNode {
            val: 2,
            next: Some(Box::new(ListNode {
                val: 1,
                next: Some(Box::new(ListNode {
                    val: 2,
                    next: Some(Box::new(ListNode {
                        val: 1,
                        next: Some(Box::new(ListNode { val: 2, next: None })),
                    })),
                })),
            })),
        })),
    }));
    println!("{:?}", Solution::modified_list(vec![1], node));

    let node = Some(Box::new(ListNode {
        val: 1,
        next: Some(Box::new(ListNode {
            val: 2,
            next: Some(Box::new(ListNode {
                val: 3,
                next: Some(Box::new(ListNode { val: 4, next: None })),
            })),
        })),
    }));
    println!("{:?}", Solution::modified_list(vec![5], node));
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
    pub fn modified_list(nums: Vec<i32>, head: Option<Box<ListNode>>) -> Option<Box<ListNode>> {
        let nums_set: HashSet<_> = nums.into_iter().collect();
        let mut dummy = ListNode { next: head, val: 0 };

        let mut cur = &mut dummy;
        while let Some(next) = cur.next.as_mut() {
            if nums_set.contains(&next.val) {
                cur.next = next.next.take();
            } else {
                cur = cur.next.as_mut().unwrap();
            }
        }

        dummy.next
    }
}
