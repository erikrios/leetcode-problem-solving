use std::cell::RefCell;

fn main() {
    let my_circular_dequeue = MyCircularDeque::new(3);
    println!("{}", my_circular_dequeue.insert_last(1));
    println!("{}", my_circular_dequeue.insert_last(2));
    println!("{}", my_circular_dequeue.insert_front(3));
    println!("{}", my_circular_dequeue.insert_front(4));
    println!("{}", my_circular_dequeue.get_rear());
    println!("{}", my_circular_dequeue.is_full());
    println!("{}", my_circular_dequeue.delete_last());
    println!("{}", my_circular_dequeue.insert_front(4));
    println!("{}", my_circular_dequeue.get_front());

    println!("=====");

    let my_circular_dequeue = MyCircularDeque::new(4);
    println!("{}", my_circular_dequeue.insert_front(9));
    println!("{}", my_circular_dequeue.delete_last());
    println!("{}", my_circular_dequeue.get_rear());
    println!("{}", my_circular_dequeue.get_front());
    println!("{}", my_circular_dequeue.get_front());
    println!("{}", my_circular_dequeue.delete_front());
    println!("{}", my_circular_dequeue.insert_front(6));
    println!("{}", my_circular_dequeue.insert_last(5));
    println!("{}", my_circular_dequeue.insert_front(9));
    println!("{}", my_circular_dequeue.get_front());
    println!("{}", my_circular_dequeue.insert_front(6));
}

struct MyCircularDeque {
    nums: RefCell<Vec<i32>>,
    current_size: RefCell<usize>,
    front: RefCell<usize>,
    rear: RefCell<usize>,
}

impl MyCircularDeque {
    fn new(k: i32) -> Self {
        Self {
            nums: RefCell::new(vec![0; k as usize]),
            current_size: RefCell::new(0),
            front: RefCell::new(0),
            rear: RefCell::new(0),
        }
    }

    fn insert_front(&self, value: i32) -> bool {
        if self.is_full() {
            return false;
        }

        if !self.is_empty() {
            if *self.front.borrow() == 0 {
                *self.front.borrow_mut() = self.nums.borrow().len() - 1;
            } else {
                *self.front.borrow_mut() -= 1;
            }
        }

        self.nums.borrow_mut()[*self.front.borrow()] = value;
        *self.current_size.borrow_mut() += 1;

        true
    }

    fn insert_last(&self, value: i32) -> bool {
        if self.is_full() {
            return false;
        }

        if !self.is_empty() {
            let new_rear = (*self.rear.borrow() + 1) % self.nums.borrow().len();
            *self.rear.borrow_mut() = new_rear;
        }

        self.nums.borrow_mut()[*self.rear.borrow()] = value;
        *self.current_size.borrow_mut() += 1;

        true
    }

    fn delete_front(&self) -> bool {
        if self.is_empty() {
            return false;
        } else if *self.current_size.borrow() == 1 {
            *self.front.borrow_mut() = 0;
            *self.rear.borrow_mut() = 0;
        } else {
            let new_front = (*self.front.borrow() + 1) % self.nums.borrow().len();
            *self.front.borrow_mut() = new_front;
        }

        *self.current_size.borrow_mut() -= 1;

        true
    }

    fn delete_last(&self) -> bool {
        if self.is_empty() {
            return false;
        } else if *self.current_size.borrow() == 1 {
            *self.rear.borrow_mut() = 0;
            *self.front.borrow_mut() = 0;
        } else if *self.rear.borrow() == 0 {
            *self.rear.borrow_mut() = self.nums.borrow().len() - 1;
        } else {
            *self.rear.borrow_mut() -= 1;
        }

        *self.current_size.borrow_mut() -= 1;

        true
    }

    fn get_front(&self) -> i32 {
        if self.is_empty() {
            return -1;
        }
        self.nums.borrow_mut()[*self.front.borrow()]
    }

    fn get_rear(&self) -> i32 {
        if self.is_empty() {
            return -1;
        }
        self.nums.borrow_mut()[*self.rear.borrow()]
    }

    fn is_empty(&self) -> bool {
        *self.current_size.borrow() == 0
    }

    fn is_full(&self) -> bool {
        *self.current_size.borrow() == self.nums.borrow().len()
    }
}
