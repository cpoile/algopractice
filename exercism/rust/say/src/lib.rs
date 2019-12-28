pub fn encode(n: u64) -> String {
    if n == 0 {
        return as_text(0);
    }
    process_chunks(chunk(n))
}

fn as_text(n: u64) -> String {
    match n {
        0 => String::from("zero"),
        1 => String::from("one"),
        2 => String::from("two"),
        3 => String::from("three"),
        4 => String::from("four"),
        5 => String::from("five"),
        6 => String::from("six"),
        7 => String::from("seven"),
        8 => String::from("eight"),
        9 => String::from("nine"),
        10 => String::from("ten"),
        11 => String::from("eleven"),
        12 => String::from("twelve"),
        13 => String::from("thirteen"),
        14 => String::from("fourteen"),
        15 => String::from("fifteen"),
        16 => String::from("sixteen"),
        17 => String::from("seventeen"),
        18 => String::from("eighteen"),
        19 => String::from("nineteen"),
        20 => String::from("twenty"),
        21..=29 => as_text(20) + "-" + &as_text(n - 20),
        30 => String::from("thirty"),
        31..=39 => as_text(30) + "-" + &as_text(n - 30),
        40 => String::from("forty"),
        41..=49 => as_text(40) + "-" + &as_text(n - 40),
        50 => String::from("fifty"),
        51..=59 => as_text(50) + "-" + &as_text(n - 50),
        60 => String::from("sixty"),
        61..=69 => as_text(60) + "-" + &as_text(n - 60),
        70 => String::from("seventy"),
        71..=79 => as_text(70) + "-" + &as_text(n - 70),
        80 => String::from("eighty"),
        81..=89 => as_text(80) + "-" + &as_text(n - 80),
        90 => String::from("ninety"),
        91..=99 => as_text(90) + "-" + &as_text(n - 90),
        100..=999 => hundred(n),
        _ => String::from(""),
    }
}

fn hundred(n: u64) -> String {
    let postfix = if n % 100 != 0 {
        String::from(" ") + &as_text(n % 100)
    } else {
        String::from("")
    };
    as_text(n / 100) + " hundred" + &postfix
}

fn chunk(n: u64) -> Vec<u64> {
    let mut chunks: Vec<u64> = vec![];
    let mut n = n;
    for place in (0..=6u32).rev() {
        chunks.push(n / (1000u64.pow(place)));
        n = n % 1000u64.pow(place);
    }
    chunks
}

fn process_chunks(v: Vec<u64>) -> String {
    let places = vec![
        " quintillion",
        " quadrillion",
        " trillion",
        " billion",
        " million",
        " thousand",
        "",
    ];
    let mut text: Vec<String> = vec![];
    for (&val, &place) in v.iter().zip(places.iter()) {
        if val != 0 {
            text.push(as_text(val) + place)
        }
    }
    text.join(" ")
}

#[test]
fn test_chunks() {
    assert_eq!(chunk(1120), vec![0, 0, 0, 0, 0, 1, 120]);
    assert_eq!(chunk(12030), vec![0, 0, 0, 0, 0, 12, 30]);
    assert_eq!(chunk(1_000_023), vec![0, 0, 0, 0, 1, 0, 23]);
    assert_eq!(chunk(123_030_033), vec![0, 0, 0, 0, 123, 30, 33]);
    assert_eq!(chunk(999_123_030_999), vec![0, 0, 0, 999, 123, 30, 999]);
    assert_eq!(chunk(1_999_123_030_123), vec![0, 0, 1, 999, 123, 30, 123]);
    assert_eq!(
        chunk(92_001_999_123_030_559),
        vec![0, 92, 1, 999, 123, 30, 559]
    );
    assert_eq!(
        chunk(12_092_001_999_123_030_325),
        vec![12, 92, 1, 999, 123, 30, 325]
    );
    assert_eq!(
        chunk(12_000_001_000_000_000_000),
        vec![12, 0, 1, 0, 0, 0, 0]
    );
}
