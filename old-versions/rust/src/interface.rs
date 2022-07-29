// Main interface for RedForest

/*
 * In the event of a new module being added either by the original devs or by contributors to
 * RedForest, an entry should be added here, in the corresponding menu. Any further CLI interfaces
 * for specific tools should not be included here. This file should not be extended beyond the
 * point of including launchers for each module or new sections -> This allows for all the modules
 * to be decentralized and the modules to be written each in their own language. It is recomended
 * to separe your own modules into files and to have a main or interface file to call with your
 * launcher here -> Dont build your launcher here -> try not to extend launching of modules more
 * than 5 to 10 lines of rust code here. 
 */

/*
 * Please add any missing modules to this list:
 * MODULES LIST:
 *
 * none still
 *
 */

// include specific files for each module interface here

mod ascii_utils;

use std::process;


// Behaviour

const CLEAR_AFTER_MODULE_SELECT: bool = false;

fn clear_screen() {
    // pass
}

fn banner() {                                                                                                          
    ascii_utils::cprintln("                                                                                                         
 ██▀███  ▓█████ ▓█████▄   █████▒▒█████   ██▀███  ▓█████   ██████ ▄▄▄█████▓
▓██ ▒ ██▒▓█   ▀ ▒██▀ ██▌▓██   ▒▒██▒  ██▒▓██ ▒ ██▒▓█   ▀ ▒██    ▒ ▓  ██▒ ▓▒
▓██ ░▄█ ▒▒███   ░██   █▌▒████ ░▒██░  ██▒▓██ ░▄█ ▒▒███   ░ ▓██▄   ▒ ▓██░ ▒░
▒██▀▀█▄  ▒▓█  ▄ ░▓█▄   ▌░▓█▒  ░▒██   ██░▒██▀▀█▄  ▒▓█  ▄   ▒   ██▒░ ▓██▓ ░ 
░██▓ ▒██▒░▒████▒░▒████▓ ░▒█░   ░ ████▓▒░░██▓ ▒██▒░▒████▒▒██████▒▒  ▒██▒ ░ 
░ ▒▓ ░▒▓░░░ ▒░ ░ ▒▒▓  ▒  ▒ ░   ░ ▒░▒░▒░ ░ ▒▓ ░▒▓░░░ ▒░ ░▒ ▒▓▒ ▒ ░  ▒ ░░   
  ░▒ ░ ▒░ ░ ░  ░ ░ ▒  ▒  ░       ░ ▒ ▒░   ░▒ ░ ▒░ ░ ░  ░░ ░▒  ░ ░    ░    
  ░░   ░    ░    ░ ░  ░  ░ ░   ░ ░ ░ ▒    ░░   ░    ░   ░  ░  ░    ░      
   ░        ░  ░   ░               ░ ░     ░        ░  ░      ░           
                 ░          

\n\n", 196);

        //                   >>-------------------<<
 //                           created by: lae-laps
          //                 >>-------------------<<

//\n\n", 196);                                                                                                                   
}

// Tool Launchers

fn account_lookup_launcher() {
    // pass
}

// Main menus

fn social_engineering_menu() {
    let list = &mut ["back", "Account Searcher", "Infectious binary creator", "Spear phishing"];
    let option: u32 = ascii_utils::display_option_list(list, 196, 196, 196);
    
    if CLEAR_AFTER_MODULE_SELECT {
        clear_screen();
    }

    if option == 0 {
        main_menu();
    } else if option == 1 {
        account_lookup_launcher();
    } else if option == 2 {
        ascii_utils::cprintln("You have selected option 1 -> not available rn -> dev, please change this...", 196); // change this
    }
}

fn osint_data_collection_menu() {
    let list = &mut ["back", "Username Searcher", "IP trackers", "Telephone Trackers", "Data Trackers", "Extend this"];
    let option: u32 = ascii_utils::display_option_list(list, 196, 196, 196);
    
    if CLEAR_AFTER_MODULE_SELECT {
        clear_screen();
    }

    if option == 0 {
        main_menu();
    } else if option == 1 {
        ascii_utils::cprintln("You have selected option x -> not available rn -> dev please change this...", 196);
    } else if option == 2 {
        ascii_utils::cprintln("You have selected option x -> not available rn -> dev please change this...", 196);
    } else if option == 3 {
        ascii_utils::cprintln("You have selected option x -> not available rn -> dev please change this...", 196);
    } else if option == 4 {
        ascii_utils::cprintln("You have selected option x -> not available rn -> dev please change this...", 196);
    } else if option == 5 {
        ascii_utils::cprintln("You have selected option x -> not available rn -> dev please change this...", 196);
    }
}

fn pentesting_web_menu() {
    let  list = &mut ["back", "PortScanner", "Vuln Scanner", "SQL Injections", "Extend this"];
    let option: u32 = ascii_utils::display_option_list(list, 196, 196, 196);
    
    if CLEAR_AFTER_MODULE_SELECT {
        clear_screen();
    }
    
    if option == 0 {
        main_menu();
    } else if option == 1 {
        ascii_utils::cprintln("You have selected option x -> not available rn -> dev please change this...", 196);
    } else if option == 2 {
        ascii_utils::cprintln("You have selected option x -> not available rn -> dev please change this...", 196);
    } else if option == 3 {
        ascii_utils::cprintln("You have selected option x -> not available rn -> dev please change this...", 196);
    } else if option == 4 {
        ascii_utils::cprintln("You have selected option x -> not available rn -> dev please change this...", 196);
    }
}

fn reverse_engineering_algorithms_menu() {
    let list = &mut ["back", "hash_cracking", "reverse binary"];
    let option: u32 = ascii_utils::display_option_list(list, 196, 196, 196);
    
    if CLEAR_AFTER_MODULE_SELECT {
        clear_screen();
    }

    if option == 0 {
        main_menu();
    } else if option == 1 {
        ascii_utils::cprintln("You have selected option x -> not available rn -> dev please change this...", 196);
    } else if option == 2 {
        ascii_utils::cprintln("You have selected option x -> not available rn -> dev please change this...", 196);
    }
}

fn main_menu() {
    let list = &mut ["exit", "Social Engineering", "OSINT & Data Collection", "Pentesting & Web", "Reverse Engineering & Algorithms"];
    let option: u32 = ascii_utils::display_option_list(list, 196, 196, 196);
    if option == 0 {
        ascii_utils::cprintln("Exiting...\n", 8);
        process::exit(0x0100);
    } else if option == 1 {
        social_engineering_menu();
    } else if option == 2 {
        osint_data_collection_menu();
    } else if option == 3 {
        pentesting_web_menu();
    } else if option == 4 {
        reverse_engineering_algorithms_menu();
    }
}

pub fn start() {
    banner();
    main_menu();
}

