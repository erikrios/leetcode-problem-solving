fn main() {
    let mut product_of_numbers = ProductOfNumbers::new();
    product_of_numbers.add(3);
    product_of_numbers.add(0);
    product_of_numbers.add(2);
    product_of_numbers.add(5);
    product_of_numbers.add(4);
    println!("{}", product_of_numbers.get_product(2));
    println!("{}", product_of_numbers.get_product(3));
    println!("{}", product_of_numbers.get_product(4));
    product_of_numbers.add(8);
    println!("{}", product_of_numbers.get_product(2));
}

struct ProductOfNumbers {
    products: Vec<i32>,
}

impl ProductOfNumbers {
    fn new() -> Self {
        Self { products: vec![] }
    }

    fn add(&mut self, num: i32) {
        for number in self.products.iter_mut().rev() {
            if *number == 0 {
                break;
            }

            *number *= num;
        }
        self.products.push(num);
    }

    fn get_product(&self, k: i32) -> i32 {
        let n = self.products.len();
        self.products[n - k as usize]
    }
}
