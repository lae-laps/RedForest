// CLI utils custom library for RedForest

use std::io;
use std::io::Write;

pub fn cprintln(content: &str, color: u8) {
    let to_print: String = format!("\x1b[38;5;{}m{}\x1b[0m", color, content);
    print!("{}", to_print);
    std::io::stdout().flush().unwrap();
}

pub fn cprintln_string(content: String, color: u8) {
    let to_print: String = format!("\x1b[38;5;{}m{}\x1b[0m", color, content);
    print!("{}", to_print);
    std::io::stdout().flush().unwrap();
}

pub fn display_option_list(arr: &mut [&str], number_color: u8, name_color: u8, prompt_color: u8) -> u32 {
    for i in 0..arr.len() {
        let content: &str = arr[i];
        let number_string: String = format!("[{}] ~ ", i.to_string());
        cprintln_string(number_string, number_color);
        cprintln(&content, name_color);
        print!("\n");
    }
    loop {
        let mut option: String = String::new();
        cprintln(">> ", prompt_color);
        io::stdin()
            .read_line(&mut option)
            .expect("[#] ~ Couldn not read user input");
        let option: u32 = match option.trim().parse() {
            Ok(num) => num,
            Err(_) => 4294967295,
        };

        if option as usize <= arr.len() - 1 {
            return option
        } else {
            let print_string: String = format!("[#] ~ Invalid input -> Please enter a number between 0 and {}\n", arr.len() - 1);
            cprintln_string(print_string, 196);
            continue
        }
    }
    
    /* Example use :
    let list = &mut ["option1", "option2", "option3"];
    ascii_utils::display_option_list(list, 196, 196, 196);
    */
}
