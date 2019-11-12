use std::cmp;
use std::fmt;

#[derive(Debug)]
pub struct Clock {
    minutes: i32,
}

impl fmt::Display for Clock {
    fn fmt(&self, f: &mut fmt::Formatter) -> fmt::Result {
        write!(f, "{:02}:{:02}", self.minutes / 60, self.minutes % 60)
    }
}

impl cmp::PartialEq for Clock {
    fn eq(&self, other: &Self) -> bool {
        self.minutes == other.minutes
    }
}

impl Clock {
    pub fn new(hours: i32, minutes: i32) -> Self {
        let hrs = (hours % 24 + 24) % 24;
        let mins = ((minutes + hrs * 60) % 1440 + 1440) % 1440;
        Clock { minutes: mins }
    }

    pub fn add_minutes(&self, minutes: i32) -> Self {
        let mins = ((self.minutes + minutes) % 1440 + 1440) % 1440;
        Clock { minutes: mins }
    }
}
