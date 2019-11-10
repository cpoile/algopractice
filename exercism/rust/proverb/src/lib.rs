use std::iter::once;

// Learned from asaf-erlich's solution
pub fn build_proverb(list: &[&str]) -> String {
    if list.len() == 0 {
        return String::new();
    }

    list.windows(2)
        .map(|ij| format!("For want of a {} the {} was lost.\n", ij[0], ij[1]))
        .chain(once(format!("And all for the want of a {}.", list[0])))
        .collect()
}

// Simple imperative solution:
pub fn build_proverb2(list: &[&str]) -> String {
    let mut res = String::new();
    for (i, s) in list.iter().enumerate() {
        if i < list.len() - 1 {
            res += &format!("For want of a {} the {} was lost.\n", s, list[i + 1]);
        } else {
            res += &format!("And all for the want of a {}.", list[0]);
        }
    }
    res
}

// Or, a poor attempt using mapping, which isn't as simple as the above:
pub fn build_proverb3(list: &[&str]) -> String {
    let res: Vec<_> = list
        .iter()
        .enumerate()
        .map(|(i, s)| {
            if i < list.len() - 1 {
                format!("For want of a {} the {} was lost.", s, list[i + 1])
            } else {
                format!("And all for the want of a {}.", list[0])
            }
        })
        .collect();
    res.join("\n")
}
