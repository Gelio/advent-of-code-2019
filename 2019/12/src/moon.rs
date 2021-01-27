use std::hash::Hash;

use crate::position::Position;

#[derive(Clone, PartialEq, Eq, Default, Debug, Hash)]
pub struct Moon {
    position: Position,
    velocity: Position,
}

impl Moon {
    pub fn new(position: Position) -> Self {
        Moon {
            position,
            ..Moon::default()
        }
    }

    pub fn adjust_velocities(m1: &mut Moon, m2: &mut Moon) {
        Self::adjust_single_direction_velocity(
            m1.position.x,
            m2.position.x,
            &mut m1.velocity.x,
            &mut m2.velocity.x,
        );
        Self::adjust_single_direction_velocity(
            m1.position.y,
            m2.position.y,
            &mut m1.velocity.y,
            &mut m2.velocity.y,
        );
        Self::adjust_single_direction_velocity(
            m1.position.z,
            m2.position.z,
            &mut m1.velocity.z,
            &mut m2.velocity.z,
        );
    }

    fn adjust_single_direction_velocity(p1: i32, p2: i32, v1: &mut i32, v2: &mut i32) {
        if p1 > p2 {
            *v1 -= 1;
            *v2 += 1;
        } else if p1 < p2 {
            *v1 += 1;
            *v2 -= 1;
        }
    }

    pub fn apply_velocity(&mut self) {
        self.position = Position::add(&self.position, &self.velocity);
    }

    pub fn kinetic_energy(&self) -> i32 {
        self.velocity.x.abs() + self.velocity.y.abs() + self.velocity.z.abs()
    }

    pub fn potential_energy(&self) -> i32 {
        self.position.x.abs() + self.position.y.abs() + self.position.z.abs()
    }

    pub fn total_energy(&self) -> i32 {
        self.kinetic_energy() * self.potential_energy()
    }
}

pub fn parse_moons(input: &str) -> Vec<Moon> {
    input
        .split("\n")
        .map(|l| l.trim())
        .filter(|l| !l.is_empty())
        .map(|l| Position::parse(l).unwrap())
        .map(Moon::new)
        .collect()
}
