use anyhow::{anyhow, Result};
use log::debug;
use std::{fs, path::PathBuf};
use advent::challenge::Challenge;
use std::collections::HashMap;
use std::sync::OnceLock;

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
    /// let four: u32 = "4".intify()?;
    /// assert_eq!(4, 4);
    ///
    /// let four = "four".intify()?;
    /// assert_eq!(4, 4);
    ///
    /// let nope = "john smith".intify();
    /// assert!(nope.is_err());
    /// ```
fn intify(number: &str) -> Result<u32> {
    // Map has the value we are searching for
    if let Some(n) = hashmap().get(number) {
        return Ok(n.to_owned());
    }

    // Map does not have the value, so it must be an int
    number.parse::<u32>()
        .map_err(|err| anyhow!(err))
}

struct Day01;

impl Challenge for Day01 {
    fn solve(&self) -> Result<&str> {
        let exe_file = fs::canonicalize(PathBuf::from(file!()))?;
        let path = exe_file
            .ancestors()
            .nth(4)
            .ok_or(anyhow!("Cannot find parents of {:?}", exe_file))
            .map(|parent| parent.join("resources").join("input-01.txt"))?;
        debug!("path: {}", path.display());

        let _data = fs::read_to_string(path)?;
        //debug!("file contents: {}", data);

        Ok("Day 01 has not been solved yet")
    }
}

pub fn main() {
    let day = Day01 {};
    day.execute();
}
