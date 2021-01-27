use crate::moon::Moon;

pub struct Simulation {
    pub moons: Vec<Moon>,
}

impl Simulation {
    pub fn new(moons: Vec<Moon>) -> Self {
        Self { moons }
    }

    pub fn run(&mut self, steps: usize) {
        for _ in 0..steps {
            self.run_single_step();
        }
    }

    pub fn run_single_step(&mut self) {
        for i in 1..self.moons.len() {
            let (adjusted_moons, moons_to_adjust) = self.moons.split_at_mut(i);
            let m1 = adjusted_moons
                .last_mut()
                .expect("cannot get moon to adjust");

            for m2 in moons_to_adjust {
                Moon::adjust_velocities(m1, m2)
            }

            m1.apply_velocity()
        }

        self.moons.last_mut().unwrap().apply_velocity();
    }

    pub fn total_energy(&self) -> i32 {
        self.moons.iter().map(Moon::total_energy).sum()
    }
}

#[cfg(test)]
mod tests {
    use crate::position::Position;

    use super::*;

    #[test]
    fn simulates_and_computes_total_energy() {
        struct TestCase<'a> {
            moon_positions: Vec<&'a str>,
            steps: usize,
            expected_total_energy: i32,
        }

        let cases = vec![
            TestCase {
                moon_positions: "<x=-1, y=0, z=2>
                <x=2, y=-10, z=-7>
                <x=4, y=-8, z=8>
                <x=3, y=5, z=-1>"
                    .split("\n")
                    .map(|l| l.trim())
                    .collect::<Vec<&str>>(),
                steps: 10,
                expected_total_energy: 179,
            },
            // TestCase {
            //     moon_positions: "<x=-8, y=-10, z=0>
            // <x=5, y=5, z=10>
            // <x=2, y=-7, z=3>
            // <x=9, y=-8, z=-3>"
            //         .split("\n")
            //         .map(|l| l.trim())
            //         .collect::<Vec<&str>>(),
            //     steps: 100,
            //     expected_total_energy: 1940,
            // },
        ];

        for case in cases {
            let moons = case
                .moon_positions
                .into_iter()
                .map(|l| Position::parse(l).unwrap())
                .map(Moon::new)
                .collect();

            let mut simulation = Simulation::new(moons);
            simulation.run(case.steps);

            assert_eq!(simulation.total_energy(), case.expected_total_energy);
        }
    }
}
