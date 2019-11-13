use std::error::Error;
use std::fs;

pub struct Config {
    pub query: String,
    pub filename: String,
}

impl Config {
    pub fn new(args: &[String]) -> Result<Config, &'static str> {
        if args.len() != 3 {
            return Err("Expected two args: search term, filename.");
        }

        Ok(Config {
            query: args[1].clone(),
            filename: args[2].clone(),
        })
    }
}

pub fn run(config: Config) -> Result<(), Box<dyn Error>> {
    let contents = fs::read_to_string(config.filename)?;

    for l in search(&config.query, &contents) {
        println!("{}", l);
    }

    Ok(())
}

pub fn search<'a>(query: &str, contents: &'a str) -> Vec<&'a str> {
    //    let mut ret: Vec<&str> = Vec::new();
    //    for l in contents.lines() {
    //        if l.contains(query) {
    //            ret.push(l);
    //        }
    //    }
    //    ret

    contents.lines().filter(|&x| x.contains(query)).collect()
}

pub fn search_case_insensitive<'a>(query: &str, contents: &'a str) -> Vec<&'a str> {
    contents
        .lines()
        .filter(|&x| x.to_uppercase().contains(&query.to_uppercase()))
        .collect()
}
