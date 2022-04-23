use ureq;

fn http_request(url: &str) -> String {
    let result = ureq::get(url).call();
    result.into_string().unwrap()
}

fn main() {
    let result = http_request("https://github.com/lae-laps");
    println!("{}", result);
}
