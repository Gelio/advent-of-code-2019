use crate::linked_list::List;

pub fn part_1(numbers: Vec<i32>) -> i32 {
    let decrypted_numbers = mix(numbers);
    get_grove_coordinates(decrypted_numbers)
}

fn mix(numbers: Vec<i32>) -> Vec<i32> {
    let mut list = List::new(numbers);
    for item in list.initially_ordered_list_items().clone().iter() {
        let value = item.borrow().value;
        list.move_item(item, value as isize);
    }

    list.into_iter().collect()
}

fn get_grove_coordinates(decrypted_numbers: Vec<i32>) -> i32 {
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
        assert_eq!(mix(numbers), vec![1, 2, -3, 4, 0, 3, -2]);
    }

    #[test]
    fn gets_right_grove_coordinates() {
        assert_eq!(get_grove_coordinates(vec![1, 2, -3, 4, 0, 3, -2]), 3);
    }
}
