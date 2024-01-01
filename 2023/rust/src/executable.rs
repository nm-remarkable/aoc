use clap::Parser;
use log::{debug, info};
use simplelog::{Config, LevelFilter, SimpleLogger};

/// Simple program to greet a person
#[derive(Parser, Debug)]
#[command(author, version, about, long_about = None)]
struct Args {
    /// Number of times to greet
    #[arg(short, long, default_value_t = false)]
    debug: bool,
}

pub fn configure(day: &str) {
    let args = Args::parse();

    let log_level = if args.debug {
        LevelFilter::Debug
    } else {
        LevelFilter::Info
    };
    let _logger = SimpleLogger::init(log_level, Config::default());

    debug!("[DEBUG]");
    info!("Running advent challenge {}!", day);
}
