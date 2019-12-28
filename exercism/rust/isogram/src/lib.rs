use std::collections::HashSet;

// "clever" version
pub fn check(candidate: &str) -> bool {
    let mut set: HashSet<_> = HashSet::new();

    candidate
        .to_lowercase()
        .chars()
        .filter(char::is_ascii_alphabetic)
        .all(|c| set.insert(c))
}

// original version:
pub fn check1(candidate: &str) -> bool {
    let mut set: HashSet<_> = HashSet::new();

    for c in candidate
        .to_lowercase()
        .chars()
        .filter(|&x| x != ' ' && x != '-')
    {
        if set.contains(&c) {
            return false;
        }
        set.insert(c);
    }
    true
}
