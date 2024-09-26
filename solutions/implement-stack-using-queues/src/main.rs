use std::{cell::RefCell, collections::VecDeque};

fn main() {}

struct MyStack {
    queue: RefCell<VecDeque<i32>>,
}

impl MyStack {
    fn new() -> Self {
        return Self {
            queue: RefCell::new(VecDeque::new()),
        };
    }

    fn push(&self, x: i32) {
        if self.queue.borrow().is_empty() {
            self.queue.borrow_mut().push_back(x);
        } else {
            self.queue.borrow_mut().push_back(x);
            let n = self.queue.borrow().len();
            for _ in 1..n {
                let front = self.queue.borrow_mut().pop_front().unwrap();
                self.queue.borrow_mut().push_back(front);
            }
        }
    }

    fn pop(&self) -> i32 {
        self.queue.borrow_mut().pop_front().unwrap()
    }

    fn top(&self) -> i32 {
        *self.queue.borrow().front().unwrap()
    }

    fn empty(&self) -> bool {
        self.queue.borrow().is_empty()
    }
}
