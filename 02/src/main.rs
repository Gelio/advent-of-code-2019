use std::io;

fn main() {
    let mut input = String::new();
    io::stdin().read_line(&mut input).expect("Cannot read line");

    let memory: Vec<usize> = input
        .trim()
        .split(",")
        .map(|x| x.parse().expect("Cannot parse number"))
        .collect();

    let result = perform_computation(&memory, 12, 2);

    println!("At first memory address: {}", result);
}

fn step(memory: &mut Vec<usize>, instruction_pointer: &mut usize) -> bool {
    let instruction = memory.get(*instruction_pointer);

    match instruction {
        Some(1) => {
            *instruction_pointer += 1;
            let values = get_3_values(&memory, *instruction_pointer);
            if values.is_none() {
                println!("Out of memory");
                return false;
            }

            let (a_ptr, b_ptr, res_ptr) = values.unwrap();

            let a = memory.get(a_ptr).unwrap();
            let b = memory.get(b_ptr).unwrap();
            memory[res_ptr] = a + b;
            *instruction_pointer += 3;

            return true;
        }
        Some(2) => {
            *instruction_pointer += 1;
            let values = get_3_values(&memory, *instruction_pointer);
            if values.is_none() {
                println!("Out of memory");
                return false;
            }

            let (a_ptr, b_ptr, res_ptr) = values.unwrap();

            let a = memory.get(a_ptr).unwrap();
            let b = memory.get(b_ptr).unwrap();
            memory[res_ptr] = a * b;
            *instruction_pointer += 3;

            return true;
        }
        Some(99) => {
            // Intended exit
        }
        Some(value) => {
            println!(
                "Invalid instruction {} at IP {}",
                value, *instruction_pointer
            );
        }
        None => {
            println!("Out of memory. IP: {}", *instruction_pointer);
        }
    }

    false
}

fn get_3_values(memory: &Vec<usize>, index: usize) -> Option<(usize, usize, usize)> {
    let memory_slice = memory.get(index..(index + 3));
    if memory_slice.is_none() {
        return None;
    }

    let memory_slice = memory_slice.unwrap();
    Some((
        *memory_slice.get(0).unwrap(),
        *memory_slice.get(1).unwrap(),
        *memory_slice.get(2).unwrap(),
    ))
}

fn perform_computation(initial_memory: &Vec<usize>, noun: usize, verb: usize) -> usize {
    let mut instruction_pointer: usize = 0;
    let mut memory = initial_memory.clone();

    memory[1] = noun;
    memory[2] = verb;

    loop {
        let should_continue = step(&mut memory, &mut instruction_pointer);

        if !should_continue {
            break;
        }
    }

    *memory.get(0).unwrap()
}
