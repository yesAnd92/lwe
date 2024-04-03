**[中文文档](README.zh-CN.md)**

## lwe
LWE stands for "Leave Work Early," which is a lighthearted way of saying "finish work early"! 🤣🤣🤣
It is a cross-platform command-line tool designed to help developers increase work efficiency. Of course, it's also suitable for those who want to use it as a project to learn Go!

In short, feel free to submit issues, fun features, or usage feedback. It would be even better if you could directly participate in the project through Pull Requests. Let's all work together and strive for an early finish!!! 💪💪💪

## Features
- [Enhanced Git operations for multiple repositories: glog, gl, gcl, gst](#git)
- Conversion from SQL statements to Java Beans, Go structures, JSON, etc.
- Transformation of SQL statements into ElasticSearch query DSL language
- PDF tools: merging multiple images or PDFs, extracting specific pages from PDFs
- Other small tools
- Retrieving passwords from Navicat connection configurations
- Synchronizing files between two directories
- Formatting request URLs

## Installation
Download the compiled executable files from the release page to use the binary on your platform!

However, a more recommended approach is to configure the binary file to your environment variables, allowing you to use it anytime, anywhere.

For more installation methods and notes, please refer to the Wiki.

## Usage
You can input `lwe` to view the usage of the LWE commands, including the subcommands and their respective functionality descriptions.

If you're interested in a specific subcommand, you can use the `-h` flag to see usage examples, such as `lwe glog -h`.

<h3 id="git">Git Enhanced Operations for Multiple Repositories: glog, gl, gcl, gst</h3>
Here are several enhanced commands centered around Git, essentially adding cross-repository operations to the original semantics.

For detailed usage of Git enhanced features, please refer to the Wiki.

#### glog: Enhances Git log functionality
  It allows you to view the commit logs of all Git repositories in a given directory. Developers often work on multiple Git repositories and may need to check the commit logs of several repositories at the same time. The `glog` subcommand comes in handy for such scenarios.

  usage:
  ```
  lwe glog [git repo dir] [-a=yesAnd] [-n=50] [-s=2023-08-04] [-e=2023-08-04]
  ```

- **gl**: Enhances the code pulling feature
  Pulls the latest code from all Git repositories in a given directory (using `git pull --rebase`).

  usage:
  ```
  lwe gl [git repo dir]
  ```

- **gcl**: Enhances the `git clone` feature
  usage:
  ```
  lwe gcl gitGroupUrl [dir for this git group] -t=yourToken
  ```

- **gst**: Views the status of all Git repositories in a specified directory
  usage:
  ```
  lwe gst [your git repo dir]
  ```

**SQL Statement Generation of Java Bean Entities, Go Structures, etc.**
If we already have a table structure, generating corresponding entities from the table creation statements can greatly reduce "mindless and repetitive" work. Currently supported structures include Java, Go, and JSON.

usage:
  ```
  lwe fmt sql-file-path [-t=java|go|json] [-a=yesAnd]
  ```

For detailed usage instructions, please refer to the Wiki.

**SQL Statement Generation of DSL Statements**
`lwe es [optional parameters] <SQL statement>`

This command helps us escape the tedious ES query syntax by converting SQL statements into the corresponding DSL and outputting them in the form of curl commands, making it convenient for use on servers as well.

Supported SQL operations in the current version:
```
lwe es 'select * from user where age >18' [-p=true]
```

For detailed usage instructions, please refer to the Wiki.

**PDF Tools: Merging Multiple Images or PDFs, Extracting Specific Pages from PDFs**
Simple editing of PDFs is a rather common feature, such as merging several PDFs or images into one, or extracting specific pages from a PDF. While this is a paid feature in many office software, LWE provides the capability for simple PDF editing.

- **pdfm**: Merge PDFs or images
  Combines multiple PDF or image files into a single PDF file in a specified order.
  usage:
  ```
  lwe pdfm out.pdf in1.pdf,in2.jpg,*.png,in3.pdf ...
  ```

- **pdfc**: Extract specified pages from a PDF
  Extracts corresponding pages from a PDF and generates a PDF file based on specified page numbers.
  usage:
  ```
  lwe pdfc [-m] in.pdf outDir 2,3,5,7-9,15 ...
  ```

For detailed usage instructions, please refer to the Wiki.

**Other Small Tools**
Some very practical features:
- Formatting request URLs
  Sometimes the URL for a request can be very long, making it difficult to find the target parameters. The `url` command can be used to format the URL, increasing the readability of the request.
  usage:
  ```
  lwe url yourUrl
  ```

- Retrieving passwords from Navicat connection configurations
  If you want to retrieve the username/password for a corresponding database from a connection saved in Navicat, you can use the `ncx` file. The `ncx` file is a connection configuration file exported by Navicat, but the password in the `ncx` file is an encrypted hexadecimal string. The `ncx` command can retrieve the corresponding plaintext.
  usage:
  ```
  lwe ncx ncx-file-path
  ```

- Synchronizing files between two directories
  If you have a habit of backing up files, this tool might help you. It can synchronize newly added files from the source directory to the backup directory, saving you the trouble of manually syncing each folder and file one by one.
  usage:
  ```
  lwe fsync sourceDir targetDir [-d=true]
  ```

For detailed usage instructions, please refer to the Wiki.

**Disclaimer**
1. The spf13/cobra library is used to conveniently build command-line tools.
2. The implementation of the `es` subcommand relies on the sqlparser library to parse SQL statements, which is an excellent library for SQL parsing.
3. The conversion from SQL to DSL heavily borrows from the elasticsql project by Cao Da, which is already a mature and easy-to-use tool. The reason for not directly using this library is to practice on our own and to have more flexibility in adding or removing features later.
4. The output of Git enhanced command results uses the go-pretty library to tabulate commit information.
5. The PDF commands are encapsulated based on pdfcpu.

**RoadMap**
- `fmt`: Support more types of conversions as needed.
- `es`: Add support for insert, update, delete operations as required.
- ...

**Open Source License**
MIT License
