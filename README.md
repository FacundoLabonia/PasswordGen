
# PasswordGen

This is a command-line interface (CLI) application built in Go that generates secure, random passwords based on user-specified parameters. The tool leverages Go’s crypto/rand package for cryptographically secure randomness and provides flexibility through customizable flags. Key features include:
Flags:
-length: Sets the password length (default: 12).

-complexity: Defines the character set complexity with three levels:

low: Lowercase letters only.

medium: Lowercase, uppercase, and digits.

high: Lowercase, uppercase, digits, and special symbols, with advanced validation to ensure at least one character from each category (if length ≥ 4).

-file: Optionally saves the generated password to a specified file.

Functionality:
Generates passwords tailored to the chosen complexity level.

For "high" complexity, ensures a mix of character types and shuffles the result for unpredictability.

Outputs the password to the console and, if specified, writes it to a file.

Use Case: Ideal for users needing quick, secure passwords for personal or professional use, with the option to store them locally.



## How to use

This guide explains how to use the pre-built executable (PasswordGen.exe) included in the repository to generate secure passwords from the command line. No compilation is required—just download the executable and run it in your terminal or command prompt.
### Prerequisites
* A terminal or command prompt (e.g., Windows Command Prompt, PowerShell, or any Unix-like shell).

* The PasswordGen.exe file from the repository, placed in a directory of your choice.

### Basic Usage
1. Open your terminal and navigate to the directory containing PasswordGen.exe using the cd command:
```bash
  cd path\to\directory
```
2. Run the executable with optional flags to generate a password. If no flags are provided, it defaults to a 12-character password with medium complexity.
### Available Flags
* -length \<number\>: Specifies the password length (e.g., -length 8). Must be greater than 0.
* -complexity \<level\>: Sets the complexity level:
    - low: Only lowercase letters (e.g., abcdxyz).
    - medium: Lowercase, uppercase, and digits (e.g., Ab7kP9mN). Default setting.
    - high: Lowercase, uppercase, digits, and symbols, ensuring at least one of each if length ≥ 4 (e.g., kP9$mNx#).
### Examples
1. Generate a default password (12 characters, medium complexity):
```bash
  PasswordGen.exe

  Output: Xy7kP9mNx2Rt
```
2. Generate an 8-character password with low complexity:
```bash
  PasswordGen.exe -length 8 -complexity low

  Output: xkcdpqrt
```
3. Generate a 16-character password with high complexity and save it to a file:
```bash
  PasswordGen.exe -length 16 -complexity high -file password.txt

  Output: 
  Generated password: P@5kLm#9nQ$jRx&T
  Password saved to password.txt
```
4. Generate a 10-character password with high complexity:
```bash
  PasswordGen.exe -length 10 -complexity high

  Output: kP9$mNx#2j
```
