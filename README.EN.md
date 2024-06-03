**[中文文档](README.md)**

## lwe
LWE stands for "Leave Work Early," which is a lighthearted way of saying "finish work early"! 🤣🤣🤣
It is a cross-platform command-line tool designed to help developers increase work efficiency. Of course, it's also suitable for those who want to use it as a project to learn Go!

In short, feel free to submit issues, fun features, or usage feedback. It would be even better if you could directly participate in the project through Pull Requests. Let's all work together and strive for an early finish!!! 💪💪💪

## Features

[Enhanced Git operations for multiple repositories: glog, gl, gcl, gst](#git)

[Conversion from SQL statements to Java Beans, Go structures, JSON, etc.](#fmt)

[Transformation of SQL statements into ElasticSearch query DSL language](#es)

[PDF tools: merging multiple images or PDFs, extracting specific pages from PDFs](#pdf)

[Other utilities](#other)
- Retrieving passwords from Navicat connection configurations
- Synchronizing files between two directories
- Formatting request URLs

## Installation

Download the compiled executable  files from the [release](https://github.com/yesAnd92/lwe/releases)

Usually,the more recommended approach is to configure the binary file to your environment variables, allowing you to use it anytime, anywhere.

For more installation methods and notes, please refer to the [Wiki](https://github.com/yesAnd92/lwe/wiki/0.%E5%AE%89%E8%A3%85%E3%80%81%E9%85%8D%E7%BD%AE%E4%BD%BF%E7%94%A8)

## Usage
You can input `lwe` to view the usage of the LWE commands, including the subcommands and their respective functionality descriptions.

If you're interested in a specific subcommand, you can use the `-h` flag to see usage examples, such as `lwe glog -h`.

<h3 id="git">Git Enhanced Operations for Multiple Repositories: glog, gl, gcl, gst</h3>
Here are several enhanced commands centered around Git, essentially adding cross-repository operations to the original semantics.

For detailed usage of Git enhanced features, please refer to the [Wiki](https://github.com/yesAnd92/lwe/wiki/3.Git%E5%A2%9E%E5%BC%BA%E5%8A%9F%E8%83%BD)

#### glog: Enhances Git log functionality
It allows you to view the commit logs of all Git repositories in a given directory. Developers often work on multiple Git repositories and may need to check the commit logs of several repositories at the same time. The `glog` subcommand comes in handy for such scenarios.

usage:
  ```
  lwe glog [git repo dir] [-a=yesAnd] [-n=50] [-s=2023-08-04] [-e=2023-08-04]
  ```

#### gl: Enhances the code pulling feature
Pulls the latest code from all Git repositories in a given directory (using `git pull --rebase`).

usage:
  ```
  lwe gl [git repo dir]
  ```

#### gcl: Enhances the `git clone` feature
usage:
  ```
  lwe gcl gitGroupUrl [dir for this git group] -t=yourToken
  ```

#### gst: Views the status of all Git repositories in a specified directory
usage:
  ```
  lwe gst [your git repo dir]
  ```

<h3 id="fmt">SQL Statement Generation of Java Bean Entities, Go Structures, etc.</h3>
If we already have a table structure, generating corresponding entities from the table creation statements can greatly reduce "mindless and repetitive" work. Currently supported structures include Java, Go, and JSON.

usage:
  ```
  lwe fmt sql-file-path [-t=java|go|json] [-a=yesAnd]
  ```

For detailed usage instructions, please refer to the [Wiki](https://github.com/yesAnd92/lwe/wiki/1.%E5%BB%BA%E8%A1%A8SQL%E8%AF%AD%E5%8F%A5%E7%94%9F%E6%88%90%E4%B8%8D%E7%94%A8%E8%AF%AD%E8%A8%80%E6%89%80%E9%9C%80%E5%AE%9E%E4%BD%93)


<h3 id="es">SQL Statement Generation of DSL Statements</h3>
`lwe es [optional parameters] <SQL statement>`

This command helps us escape the tedious ES query syntax by converting SQL statements into the corresponding DSL and outputting them in the form of curl commands, making it convenient for use on servers as well.

Supported SQL operations in the current version:
```
lwe es 'select * from user where age >18' [-p=true]
```

For detailed usage instructions, please refer to the [Wiki](https://github.com/yesAnd92/lwe/wiki/2.%E5%B0%86SQL%E8%AF%AD%E5%8F%A5%E8%BD%AC%E6%8D%A2%E6%88%90ElasticSearch%E6%9F%A5%E8%AF%A2%E7%9A%84DSL%E8%AF%AD%E8%A8%80)

<h3 id="pdf">PDF Tools: Merging Multiple Images or PDFs, Extracting Specific Pages from PDFs</h3>
Simple editing of PDFs is a rather common feature, such as merging several PDFs or images into one, or extracting specific pages from a PDF. While this is a paid feature in many office software, LWE provides the capability for simple PDF editing.

#### pdfm: Merge PDFs or images
Combines multiple PDF or image files into a single PDF file in a specified order.
usage:
  ```
  lwe pdfm out.pdf in1.pdf,in2.jpg,*.png,in3.pdf ...
  ```

#### pdfc: Extract specified pages from a PDF
Extracts corresponding pages from a PDF and generates a PDF file based on specified page numbers.
usage:
  ```
  lwe pdfc [-m] in.pdf outDir 2,3,5,7-9,15 ...
  ```

For detailed usage instructions, please refer to the [Wiki](https://github.com/yesAnd92/lwe/wiki/PDF%E5%B7%A5%E5%85%B7%EF%BC%9A%E5%90%88%E5%B9%B6%E5%A4%9A%E4%B8%AA%E5%9B%BE%E7%89%87%E6%88%96%E8%80%85PDF%E3%80%81%E6%88%AA%E5%8F%96PDF%E6%8C%87%E5%AE%9A%E9%A1%B5#pdfm-%E5%90%88%E5%B9%B6pdf%E6%88%96%E8%80%85%E5%9B%BE%E7%89%87)

<h3 id="other">Other utilities</h3>
Some highly practical and efficient tools

<h4> Formatting request URLs</h4>
  Sometimes the URL for a request can be very long, making it difficult to find the target parameters. The `url` command can be used to format the URL, increasing the readability of the request.

usage:
  ```
  lwe url yourUrl
  ```
For detailed, [Wiki](https://github.com/yesAnd92/lwe/wiki/%E5%85%B6%E5%AE%83%E5%B0%8F%E5%B7%A5%E5%85%B7#%E6%A0%BC%E5%BC%8F%E5%8C%96%E8%AF%B7%E6%B1%82url)


<h4>  Retrieving passwords from Navicat connection configurations</h4>
  If you want to retrieve the username/password for a corresponding database from a connection saved in Navicat, you can use the `ncx` file. The `ncx` file is a connection configuration file exported by Navicat, but the password in the `ncx` file is an encrypted hexadecimal string. The `ncx` command can retrieve the corresponding plaintext.

usage:
  ```
  lwe ncx ncx-file-path
  ```
For detailed, [Wiki](https://github.com/yesAnd92/lwe/wiki/%E5%85%B6%E5%AE%83%E5%B0%8F%E5%B7%A5%E5%85%B7#%E8%8E%B7%E5%8F%96navicat%E8%BF%9E%E6%8E%A5%E9%85%8D%E7%BD%AE%E4%B8%AD%E7%9A%84%E5%AF%86%E7%A0%81)

<h4> Synchronizing files between two directories </h4>
  If you have a habit of backing up files, this tool might help you. It can synchronize newly added files from the source directory to the backup directory, saving you the trouble of manually syncing each folder and file one by one.

usage:
  ```
  lwe fsync sourceDir targetDir [-d=true]
  ```
For detailed, [Wiki](https://github.com/yesAnd92/lwe/wiki/%E5%85%B6%E5%AE%83%E5%B0%8F%E5%B7%A5%E5%85%B7#%E5%90%8C%E6%AD%A5%E4%B8%A4%E4%B8%AA%E7%9B%AE%E5%BD%95%E4%B8%8B%E6%96%87%E4%BB%B6)


## Disclaimer
1. The spf13/cobra library is used to conveniently build command-line tools.
2. The implementation of the `es` subcommand relies on the sqlparser library to parse SQL statements, which is an excellent library for SQL parsing.
3. The conversion from SQL to DSL heavily borrows from the elasticsql project by Cao Da, which is already a mature and easy-to-use tool. The reason for not directly using this library is to practice on our own and to have more flexibility in adding or removing features later.
4. The output of Git enhanced command results uses the go-pretty library to tabulate commit information.
5. The PDF commands are encapsulated based on pdfcpu.

## RoadMap
- `fmt`: Support more types of conversions as needed.
- `es`: Add support for insert, update, delete operations as required.
- ...

## License
MIT License
