use clap::Parser;
use log::{debug, info};
use simplelog::{Config, LevelFilter, SimpleLogger};
use std::panic::Location;

#[derive(Parser, Debug)]
struct Args {
    #[arg(short, long, default_value_t = false)]
    debug: bool,
}

pub trait Challenge {
    fn execute(&self, file_name: &str) {
        self.setup(file_name);
        info!("{}", self.solve())
    #[track_caller]
    fn execute(&self) {
        self.setup(Location::caller().file());
    }
    fn setup(&self, file_name: &str) {
        let args = Args::parse();

        let log_level = if args.debug {
            LevelFilter::Debug
        } else {
            LevelFilter::Info
        };
        let _logger = SimpleLogger::init(log_level, Config::default());

        debug!("[DEBUG]");
        info!("Running advent challenge {}!", file_name);
    }
    fn solve(&self) -> &str;
}
