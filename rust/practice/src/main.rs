#[derive(Debug)]
struct Rect {
    len: f32,
    wid: f32,
}

impl Rect {
    fn area(&self) -> f32 {
        self.len * self.wid
    }
}

enum IpAddrKind {
    V4(String),
    V6(String)
}

fn main() {
    let mut s1 = String::from("Hello");

    let s2 = s1.clone();
    s1 += ", world!";
    println!("str: {}, str2: {}", s1, s2);

    let s3 = "語";
    println!("length of {} is {}, first char is {}",
             s3, s3.len(), s3.bytes().nth(1).expect("something"));

    let s = String::from("界,　Hello, how are you?");


    println!("s: {}", first_word(&s));

    let r1 = Rect {len: 5.0, wid: 5.25};
    println!("area of {:#?} is {}", r1, r1.area());

    let home = IpAddrKind::V4(String::from("1.2.3.4"));
    let _looopback = IpAddrKind::V6(String::from("::1"));

    let _x: Option<bool> = None;

    match home {
        IpAddrKind::V4(addr) => println!("addr is {}", addr),
        IpAddrKind::V6(addr) => println!("addr is {}", addr),
    }
}

fn first_word(s: &str) -> String {
    s.chars().take_while(|x| !x.is_whitespace()).collect()
}
