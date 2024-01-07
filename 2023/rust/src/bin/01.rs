use anyhow::{anyhow, Result};
use log::debug;
use std::{fs, path::PathBuf};

use advent::challenge::Challenge;

use std::collections::HashMap;
use std::sync::OnceLock;

fn word_to_number() -> &'static HashMap<&'static str, u32> {
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
