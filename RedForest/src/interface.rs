// Main interface for RedForest

mod ascii_utils;

pub fn banner() {                                                                                                          
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
                           >>-------------------<<
                            created by: lae-laps
                            and wans... some day
                           >>-------------------<<

\n\n", 196);                                                                                                                   
}

pub fn start() {
    banner();
    let list = &mut ["exit", "Social Engineering", "OSINT & Data Collection", "Pentesting & Web", "Reverse engineering & Decryption"];
    ascii_utils::display_option_list(list, 196, 196, 196);
}
