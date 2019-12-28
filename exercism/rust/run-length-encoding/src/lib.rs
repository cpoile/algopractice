pub fn encode(source: &str) -> String {
    let mut ret = String::new();
    if source.len() == 0 {
        return ret;
    }
    let mut count = 1;
    let mut cur_char = source.chars().next().unwrap();
    for c in source.chars().skip(1) {
        if c == cur_char {
            count += 1;
        } else {
            add_char(&mut ret, count, cur_char);
            count = 1;
            cur_char = c;
        }
    }
    add_char(&mut ret, count, cur_char);
    ret
}

fn add_char(ret: &mut String, count: i32, cur_char: char) {
    if count == 1 {
        ret.push(cur_char);
    } else {
        ret.push_str(&(count.to_string() + &cur_char.to_string()));
    }
}

pub fn decode(source: &str) -> String {
    let mut ret = String::new();
    let mut cur_count = String::from("0");
    for c in source.chars() {
        if c.is_numeric() {
            cur_count += &c.to_string();
        } else if cur_count == "0" {
            ret.push(c);
        } else {
            ret += &c.to_string().repeat(cur_count.parse().unwrap());
            cur_count = String::from("0");
        }
    }
    ret
}
