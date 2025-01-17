use std::collections::HashMap;

fn main() {}

struct UndergroundSystem {
    average_times: HashMap<(String, String), (f64, usize)>,
    checks: HashMap<i32, (String, i32)>,
}

/**
 * `&self` means the method takes an immutable reference.
 * If you need a mutable reference, change it to `&mut self` instead.
 */
impl UndergroundSystem {
    fn new() -> Self {
        Self {
            average_times: HashMap::new(),
            checks: HashMap::new(),
        }
    }

    fn check_in(&mut self, id: i32, station_name: String, t: i32) {
        self.checks.insert(id, (station_name, t));
    }

    fn check_out(&mut self, id: i32, station_name: String, t: i32) {
        let departure = self.checks.remove(&id).unwrap();
        let time = t - departure.1;

        if let Some(avg_time) = self
            .average_times
            .get_mut(&(departure.0.clone(), station_name.clone()))
        {
            avg_time.0 += time as f64;
            avg_time.1 += 1;
        } else {
            self.average_times
                .insert((departure.0, station_name), (time as f64, 1));
        }
    }

    fn get_average_time(&self, start_station: String, end_station: String) -> f64 {
        let avg = *self
            .average_times
            .get(&(start_station, end_station))
            .unwrap();
        avg.0 / (avg.1 as f64)
    }
}
