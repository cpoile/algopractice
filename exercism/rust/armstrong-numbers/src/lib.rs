pub fn is_armstrong_number(num: u32) -> bool {
    let mut digits: Vec<u32> = Vec::new();
    let mut n = num;
    while n > 0 {
        let d = n % 10;
        n = n / 10;
        digits.push(d);
    }
    let sum = digits
        .iter()
        .map(|x| x.pow(digits.len() as u32))
        .sum();

    num == sum
}
