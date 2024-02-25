use advent::challenge::Challenge;
use anyhow::{anyhow, Result};
use log::debug;
use regex::Regex;
use std::collections::HashMap;
use std::sync::OnceLock;
use std::{fs, path::PathBuf};

static NUMBERS: &str = r"\d|one|two|three|four|five|six|seven|eight|nine";
static REVERSED_NUMBERS: &str = r"\d|enin|thgie|neves|xis|evif|ruof|eerht|owt|eno|\d";

fn hashmap() -> &'static HashMap<&'static str, u32> {
    static HASHMAP: OnceLock<HashMap<&str, u32>> = OnceLock::new();
    HASHMAP.get_or_init(|| {
        let hmap: HashMap<&str, u32> = [
            ("one", 1),
            ("two", 2),
            ("three", 3),
            ("four", 4),
            ("five", 5),
            ("six", 6),
            ("seven", 7),
            ("eight", 8),
            ("nine", 9),
        ]
        .iter()
        .cloned()
        .collect();
        hmap
    })
}
/// Parses a string into an int. Supports words for digits.
/// Test won't run because of https://github.com/rust-lang/cargo/issues/5477
///
/// # Examples
///
/// ```
/// let four = intify("4");
/// assert!(four.is_ok());
/// assert_eq!(four.unwrap(), 4);
///
/// let four_again = intify("four");
/// assert!(four_again.is_ok());
/// assert_eq!(four_again.unwrap(), 4);
///
/// let nope = intify("john smith");
/// assert!(nope.is_err());
/// ```
fn intify(number: &str) -> Result<u32> {
    // Map has the value we are searching for
    if let Some(n) = hashmap().get(number) {
        return Ok(n.to_owned());
    }

    // Map does not have the value, so it must be an int
    number.parse::<u32>().map_err(|err| anyhow!(err))
}

fn match_digits(line: &str, reverse: bool) -> Result<u32> {
    let (regex, input) = match reverse {
        true => (
            REVERSED_NUMBERS,
            line.to_string().chars().rev().collect::<String>(),
        ),
        false => (NUMBERS, line.to_string()),
    };
    let re = Regex::new(regex)?;

    let re_match = re.captures(&input).unwrap().get(0).unwrap().as_str();
    let sure_match = match reverse {
        true => re_match.chars().rev().collect::<String>(),
        false => re_match.to_string(),
    };
    intify(&sure_match)
}

struct Day01;

impl Challenge for Day01 {
    fn solve(&self) -> Result<String> {
        let exe_file = fs::canonicalize(PathBuf::from(file!()))?;
        let path = exe_file
            .ancestors()
            .nth(4)
            .ok_or(anyhow!("Cannot find parents of {:?}", exe_file))
            .map(|parent| parent.join("resources").join("01").join("input.txt"))?;
        debug!("path: {}", path.display());

        let mut sum = 0;
        for line in fs::read_to_string(path)?.lines() {
            let first = match_digits(line, false)?.to_string();
            let last = match_digits(line, true)?.to_string();

            let result = first + last.as_str();
            sum += result.parse::<u32>().unwrap_or(0);
        }

        Ok(format!("The sum is: {}", sum))
    }
}

pub fn main() {
    let day = Day01 {};
    day.execute();
}

#[cfg(test)]
mod test {
    use super::intify;

    #[test]
    fn test_infify() {
        let four = intify("4");
        assert!(four.is_ok());
        assert_eq!(four.unwrap(), 4);

        let four_again = intify("four");
        assert!(four_again.is_ok());
        assert_eq!(four_again.unwrap(), 4);

        let nope = intify("john smith");
        assert!(nope.is_err());
    }
}
