use std::io;

mod tls;
mod ssl;

fn main() {
    let mut address = String::new();
    let mut tls_addresses_count = 0;
    let mut ssl_addresses_count = 0;

    loop {
        address.clear();

        io::stdin()
            .read_line(&mut address)
            .expect("Error while reading line");

        let address = address.trim();
        if address.len() == 0 {
            break;
        }

        if tls::supports_tls(address) {
            tls_addresses_count += 1;
        }

        if ssl::supports_ssl(address) {
            ssl_addresses_count += 1;
        }
    }

    println!("TLS addresses count: {}", tls_addresses_count);
    println!("SSL addresses count: {}", ssl_addresses_count);
}

