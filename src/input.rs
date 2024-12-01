use std::fs;

pub fn read_day_input(day: u8) -> String {
    let file_path = format!("inputs/day{}.txt", day);
    let contents = fs::read_to_string(&file_path);
    contents.expect(&file_path)
}
