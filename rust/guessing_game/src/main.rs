use std::io;
use std::cmp::Ordering;
use rand::{thread_rng, Rng};

fn main() {
    let target: u32 = thread_rng().gen_range(1, 101);

    println!("Let's play a guessing game. \
    I've picked a number between 1 and 100, try to guess it.");

    for x in 1..6 {
        println!("This is try {}", x);
        let mut guess = String::new();
        io::stdin().read_line(&mut guess)
            .expect("Failed to read line");

        let guess: u32 = match guess.trim().parse() {
            Ok(num) => num,
            Err(_) => continue,
        };

        match guess.cmp(&target) {
            Ordering::Less => println!("Too low!"),
            Ordering::Greater => println!("Too high."),
            Ordering::Equal => {
                println!("You're right!");
                return;
            },
        };
    }
    println!("Sorry, you failed. It was {}.", target)
}
