pub fn reply(message: &str) -> &str {
    let msg = message.trim();
    let empty = msg.len() == 0;
    let has_alpha = msg.to_uppercase() != msg.to_lowercase();
    let yell = has_alpha && msg == msg.to_uppercase();
    let question = msg.ends_with('?');

    if empty {
        "Fine. Be that way!"
    } else if yell && question {
        "Calm down, I know what I'm doing!"
    } else if yell {
        "Whoa, chill out!"
    } else if question {
        "Sure."
    } else {
        "Whatever."
    }
}
