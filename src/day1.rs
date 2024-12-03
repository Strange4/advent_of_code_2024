use std::collections::HashMap;

use input_lib::read_day_input;

fn main() {
    part1();
    part2();
}

fn get_lists() -> (Vec<u32>, Vec<u32>) {
    let input = read_day_input(1);
    let lines = input.lines();
    let lists: (Vec<_>, Vec<_>) = lines
        .map(|line| {
            let mut split = line.split("   ");
            let left = split.next().unwrap().parse::<u32>().unwrap();
            let right = split.next().unwrap().parse::<u32>().unwrap();
            (left, right)
        })
        .unzip();
    lists
}

fn part1() {
    let (mut left_ids, mut right_ids) = get_lists();
    left_ids.sort();
    right_ids.sort();

    let sum = left_ids
        .iter()
        .zip(right_ids.iter())
        .fold(0, |sum, (left, right)| sum + left.abs_diff(*right));
    println!("Part 1: {}", sum);
}

fn part2() {
    let (left_ids, right_ids) = get_lists();
    let mut right_number_counts: HashMap<u32, u32> = HashMap::new();
    right_ids
        .into_iter()
        .fold(&mut right_number_counts, |map, number| {
            let new_number = map.get(&number).unwrap_or(&0);
            map.insert(number, *new_number + 1);
            map
        });
    let sum = left_ids.into_iter().fold(0, |sum, number| {
        sum + (number * right_number_counts.get(&number).unwrap_or(&0))
    });
    println!("Part 2: {}", sum);
}
