// better version:
pub fn sum_of_multiples(limit: u32, factors: &[u32]) -> u32 {
    (1..limit).filter(|i|
        factors.iter()
            .filter(|&f| *f != 0)
            .any(|f| i % f == 0)
    ).sum()
}

// original:
pub fn sum_of_multiples2(limit: u32, factors: &[u32]) -> u32 {
    let mut ret: Vec<u32> = Vec::new();
    for &f in factors {
        if f == 0 {
            continue;
        }
        for x in 1u32.. {
            if f * x >= limit {
                break;
            }
            ret.push(f * x)
        }
    }
    ret.sort();
    ret.dedup();
    ret.iter().sum()
}

