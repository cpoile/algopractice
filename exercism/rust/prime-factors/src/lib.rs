pub fn factors(n: u64) -> Vec<u64> {
    let mut ret: Vec<u64> = Vec::new();
    let mut n = n;
    let mut f: u64 = 2;
    while n > 1 {
        if n % f == 0 {
            ret.push(f);
            n /= f;
        } else {
            f += 1;
        }
    }
    ret
}
