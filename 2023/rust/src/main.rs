mod days;

use clap::Parser;
use log::{debug, info};
use simplelog::{Config, LevelFilter, SimpleLogger};

/// Simple program to greet a person
#[derive(Parser, Debug)]
#[command(author, version, about, long_about = None)]
struct Args {
    /// Name of the person to greet
    #[arg(value_name = "day", index = 1)]
    day: String,

    /// Number of times to greet
    #[arg(short, long, default_value_t = false)]
    debug: bool,
}

fn main() {
    let args = Args::parse();

    let log_level = if args.debug {
        LevelFilter::Debug
    } else {
        LevelFilter::Info
    };
    let _logger = SimpleLogger::init(log_level, Config::default());

    debug!("[DEBUG]");
    info!("Running advent challenge {}!", args.day);
    days::r01::solve();
}
