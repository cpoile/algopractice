use minigrep::Config;
use std::env;
use std::process;

fn main() {
    let args: Vec<String> = env::args().collect();

    let cfg = Config::new(&args).unwrap_or_else(|err| {
        println!("Error parsing args: {}", err);
        process::exit(1);
    });

    if let Err(e) = minigrep::run(cfg) {
        println!("Application error: {}", e);
        process::exit(1);
    }
}

#[cfg(test)]
mod tests {
    use minigrep::*;

    #[test]
    fn one_result() {
        let query = "duct";
        let contents = "\
Rust:
safe, fast, productive.
Pick three.";

        assert_eq!(vec!["safe, fast, productive."], search(query, contents));
    }

    #[test]
    fn case_insensitive() {
        let query = "DucT";
        let contents = "\
Rust:
safe, fast, productive.
Pick three.";

        assert_eq!(search(query, contents), Vec::<&str>::new());
        assert_eq!(
            search_case_insensitive(query, contents),
            vec!["safe, fast, productive."]
        );
    }
}
