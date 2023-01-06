use crate::linked_list::List;

pub fn part_1(numbers: Vec<i64>) -> i64 {
    let mut list = List::new(numbers);
    mix(&mut list);
    let decrypted_numbers = list.into_iter().collect();
    get_grove_coordinates(decrypted_numbers)
}

const DECRYPTION_KEY: i64 = 811589153;

pub fn part_2(numbers: Vec<i64>) -> i64 {
    let mut list = List::new(
        numbers
            .into_iter()
            .map(|num| num * DECRYPTION_KEY)
            .collect(),
    );
    for _ in 0..10 {
        mix(&mut list);
    }

    let decrypted_numbers = list.into_iter().collect();
    get_grove_coordinates(decrypted_numbers)
}

fn mix(list: &mut List<i64>) {
    for item in list.initially_ordered_list_items().clone().iter() {
        let value = item.borrow().value;
        list.move_item(item, value as isize);
    }
}

fn get_grove_coordinates(decrypted_numbers: Vec<i64>) -> i64 {
    let (first_zero_index, _) = decrypted_numbers
        .iter()
        .enumerate()
        .find(|(_, number)| **number == 0)
        .expect("Numbers did not contain any zeros");

    vec![1000, 2000, 3000]
        .into_iter()
        .map(|index| decrypted_numbers[(first_zero_index + index) % decrypted_numbers.len()])
        .sum()
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn mixes() {
        let numbers = vec![1, 2, -3, 3, -2, 0, 4];
        let mut list = List::new(numbers);
        mix(&mut list);
        assert_eq!(
            list.into_iter().collect::<Vec<_>>(),
            vec![1, 2, -3, 4, 0, 3, -2]
        );
    }

    #[test]
    fn gets_right_grove_coordinates() {
        assert_eq!(get_grove_coordinates(vec![1, 2, -3, 4, 0, 3, -2]), 3);
    }
}
