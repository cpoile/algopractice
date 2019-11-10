// bad version:
pub fn nth1(n: u32) -> u32 {
    let mut count = 0;
    let mut cur_num = 2;
    while count < n {
        cur_num += 1;
        let sqrtn = (cur_num as f64).sqrt() as u32;
        let mut prime = true;
        for x in 2..=sqrtn + 1 {
            if cur_num % x == 0 {
                prime = false;
                break;
            }
        }
        if prime {
            count += 1;
        }
    }
    return cur_num;
}

// better:
pub fn nth2(n: u32) -> u32 {
    let mut count = 0;
    let mut cur_num = 2;
    while count < n {
        cur_num += 1;
        if is_prime(cur_num) {
            count += 1;
        }
    }
    return cur_num;
}

// best?
pub fn nth(n: u32) -> u32 {
    (2..).filter(|&x| is_prime(x)).nth(n as usize).unwrap()
}

fn is_prime(x: u32) -> bool {
    if x == 2 {
        return true;
    }
    let sqrtx = (x as f64).sqrt() as u32;
    !(2..=sqrtx + 1).any(|i| x % i == 0)
}