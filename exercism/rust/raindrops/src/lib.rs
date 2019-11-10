pub fn raindrops1(n: u32) -> String {
    let v = vec![(3, "Pling"), (5, "Plang"), (7, "Plong")];
    let mut res = String::new();
    for (x, s) in v {
        if n % x == 0 {
            res += s
        }
    }
    if res.len() == 0 {
        return n.to_string();
    }
    res
}

pub fn raindrops(n: u32) -> String {
    let v = vec![(3, "Pling"), (5, "Plang"), (7, "Plong")];

    let res: String = v.iter()
        .filter_map(|(x, s)| if n % *x == 0 {Some(s.to_string())} else {None})
        .collect();

    if res.is_empty() {
        return n.to_string();
    }
    res
}
