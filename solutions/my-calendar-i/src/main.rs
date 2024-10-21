use std::cell::RefCell;

fn main() {
    let my_calendar = MyCalendar::new();
    println!("{}", my_calendar.book(10, 20));
    println!("{}", my_calendar.book(15, 25));
    println!("{}", my_calendar.book(20, 30));

    println!("========");

    let my_calendar = MyCalendar::new();
    let books = [
        (47, 50),
        (33, 41),
        (39, 45),
        (33, 42),
        (25, 32),
        (26, 35),
        (19, 25),
        (3, 8),
        (8, 13),
        (18, 27),
    ];

    for book in books {
        println!("{}", my_calendar.book(book.0, book.1));
    }

    println!("========");

    let my_calendar = MyCalendar::new();
    let books = [
        (97, 100),
        (33, 51),
        (89, 100),
        (83, 100),
        (75, 92),
        (76, 95),
        (19, 30),
        (53, 63),
        (8, 23),
        (18, 37),
        (87, 100),
        (83, 100),
        (54, 67),
        (35, 48),
        (58, 75),
        (70, 89),
        (13, 32),
        (44, 63),
        (51, 62),
        (2, 15),
    ];
    for book in books {
        println!("{}", my_calendar.book(book.0, book.1));
    }
}

struct MyCalendar {
    books: RefCell<Vec<(i32, i32)>>,
}

impl MyCalendar {
    fn new() -> Self {
        Self {
            books: RefCell::new(Vec::new()),
        }
    }

    fn book(&self, start: i32, end: i32) -> bool {
        for book in self.books.borrow().iter() {
            if ((start >= book.0 && start < book.1) || (end > book.0 && end <= book.1))
                || ((book.0 >= start && book.1 < start) || (book.1 > start && book.1 <= end))
            {
                return false;
            }
        }

        self.books.borrow_mut().push((start, end));

        true
    }
}
