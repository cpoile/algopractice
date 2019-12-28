use std::cmp;
use std::fmt;

const min_per_hour: i32 = 60;
const min_per_day: i32 = min_per_hour * 24;

#[derive(Debug)]
pub struct Clock {
    minutes: i32,
}

impl fmt::Display for Clock {
    fn fmt(&self, f: &mut fmt::Formatter) -> fmt::Result {
        // note: rustfmt insists on this formatting:
        write!(
            f,
            "{:02}:{:02}",
            self.minutes / min_per_hour,
            self.minutes % min_per_hour
        )
    }
}

impl cmp::PartialEq for Clock {
    fn eq(&self, other: &Self) -> bool {
        self.minutes == other.minutes
    }
}

impl Clock {
    pub fn new(hours: i32, minutes: i32) -> Self {
        let mins = ((minutes + hours * min_per_hour) % min_per_day + min_per_day) % min_per_day;
        Clock { minutes: mins }
    }

    pub fn add_minutes(&self, minutes: i32) -> Self {
        let mins = ((self.minutes + minutes) % min_per_day + min_per_day) % min_per_day;
        Clock { minutes: mins }
    }
}
