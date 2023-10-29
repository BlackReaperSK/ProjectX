# ProjectX: Your own Data Leak Repository

[![License](https://img.shields.io/badge/license-MIT-blue.svg)](https://opensource.org/licenses/MIT)

## Overview

ProjectX is an educational repository that aims to teach developers how to prevent credential leaks and create their own XKeyscore (Intelligence/Exploitation tool). This project emphasizes the importance of security and responsible programming practices to protect sensitive information. Depending on your region, it may be illegal to access this type of data. In the USA, it's not illegal to access leaked data. Consider learning about your country's jurisdiction before embarking on this project. 


## Table of Contents

- [ProjectX: Your own Data Leak Repository](#projectx-your-own-data-leak-repository)
  - [Overview](#overview)
  - [Table of Contents](#table-of-contents)
  - [Features](#features)
  - [Getting Started](#getting-started)
  - [Usage](#usage)
     - [Add Data](#add-data)
     - [Search Portal](#search-portal)
     - [Export](#export-data)

## Features

- Easy big data search and visualization.
- Understand how data leaks works and how to protect sensitive information.
- Create your own XKeyScore or IntelX with easly.
- Selected Data Export

## Getting Started

These instructions will guide you on setting up the project on your local machine with some example data. 

1. Clone the repository:

```bash
git clone https://github.com/BlackReaperSK/ProjectX.git
cd ProjectX
```
2. Build and run the server

```bash
go build
./projectx
```

## Usage
### Add Data 
To get started, obtain your dataset. For initial testing, I've provided generated examples in the "leaks" directory. Ensure that your data is in a readable file format.

The platform utilizes multiple routines to concurrently read all files within the "leaks" directory, searching for occurrences of your specified query. This process is file-type agnostic, offering a versatile and comprehensive search experience.

When adding data to the platform, you can use the cp command:

```bash
cp -r your_data_directory ./ProjectX/leaks
```
### Search Portal 
After starting the server, access it by navigating to http://localhost:8080/portal/ (You can customize the port by modifying the configuration file conf.go). Once there, you will be greeted with a user interface similar to the one below.

Simply input your query and set the limit for results (It's not recommended to exceed 1000; for larger datasets, consider using the export feature instead).

![image](https://github.com/BlackReaperSK/ProjectX/assets/82294569/3294c758-db5e-45be-87c7-39e1bcc709b3)

For optimal performance, fine-tune the number of routines based on your dataset size and server specifications to ensure a response time below 2000ms. Exceeding this threshold may lead to significant slowdowns over time. Upon completing the search, the platform provides detailed results, including the file path and line number where the occurrence was found. Additionally, it offers a contextual view with a count of five lines both above and below the identified line.

![image](https://github.com/BlackReaperSK/ProjectX/assets/82294569/e04b2d19-0b5f-419c-9a4b-fc804a4c8c2c)

### Export Data

The export route is crafted to cater to the daily needs of Red Team professionals, offering swift and effortless access to critical information. It proves invaluable for Data Leak report projects, enabling the seamless bulk download and analysis of selected data. Leveraging tools like Curl or WGET, you can efficiently perform searches and fetch files with ease.

```bash
wget -O leaks.txt "https://localhost:8080/export?q=example.com&limit=500"
or
curl "https://localhost:8080/export?q=example.com&limit=500" -O leaks.txt
```
