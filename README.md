# IP PTR Lookup Script

This is a simple Go script that performs PTR (reverse DNS) lookups for IP addresses. It can handle single IP addresses, IP addresses from a file, or IP addresses provided via stdin.

## Example

```
ptr 8.8.8.8
8.8.8.8,dns.google

ptr 185.15.59.224
185.15.59.224,text-lb.esams.wikimedia.org

echo "8.8.8.8" | ptr -
8.8.8.8,dns.google

ptr ips.txt
9.9.9.9,dns9.quad9.net
```

## Installation

To install ptr, you can use the following command:

```shell
go install github.com/topscoder/ptr@latest
```

This will install the ptr script as an executable in your Go bin directory.

## Usage

Run the script with different options:

* To look up a single IP address:

```shell
ptr 8.8.8.8
```

* To process IP addresses from a file (one IP address per line):

```shell
ptr ips.txt
```

* To provide IP addresses via stdin (press Ctrl + D to signal the end of input):

```shell
cat ips.txt | ptr -
```

Enjoy the PTR lookup results!

## License

This project is licensed under the MIT License.

Feel free to fork the repository, make improvements, and submit pull requests!

## Acknowledgements

This script was inspired by the need for a simple tool to perform PTR lookups for multiple IP addresses quickly.
