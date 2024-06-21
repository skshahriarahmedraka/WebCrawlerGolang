
# WebCrawlerGolang

## Overview

WebCrawlerGolang is a powerful and efficient web crawler built with Go, designed to scrape product information from websites. This tool is perfect for extracting data for analysis, research, or integration into other systems.

## Features

- High-performance web scraping
- Customizable parameters for specific use cases
- Easy to use and configure

## Prerequisites

Before you begin, ensure you have met the following requirements:

- Go installed on your local machine
- Git installed on your local machine

## Getting Started

Follow these steps to set up and run the web crawler:

### Clone the Repository

First, clone the repository from GitHub:

```bash
git clone git@github.com:skshahriarahmedraka/WebCrawlerGolang.git
```

### Initialize the Go Module

Navigate to the project directory and initialize the Go module to install all dependencies:

```bash
cd WebCrawlerGolang
go mod tidy
```

### Run the Crawler

You can now run the crawler with the desired parameters. The following example demonstrates how to scrape data for a specific product:

```bash
go run ./main.go crawler --productid=F36640 --count=300 --base-url=https://shop.adidas.jp/products
```

### Command-Line Arguments

- `--productid`: Specifies the product ID to be scraped (e.g., F36640).
- `--count`: Specifies the number of products to scrape.
- `--base-url`: Specifies the base URL of the website from which to scrape product information.

## Contributing

Contributions are always welcome! If you have suggestions for improvements or new features, feel free to open an issue or submit a pull request.

## License

This project is licensed under the MIT License.



