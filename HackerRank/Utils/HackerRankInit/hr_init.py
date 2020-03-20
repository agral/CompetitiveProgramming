#!/usr/bin/env python3

# Name:          HackerRankInit.py
# Description:   Downloads a challenge description and its testcases
#                from hackerrank.com website to a proper location, depending
#                on the challenge's category.
# Options:       None
# Created on:    20.03.2020
# Last modified: 20.03.2020
# Author:        Adam Graliński (adam@gralin.ski)
# License:       MIT


import argparse
import os
import re
import sys
import zipfile

from datetime import date
from bs4 import BeautifulSoup
import requests

HACKERRANK_SOLUTIONS_DIR = os.path.join(
    os.path.expanduser("~"),
    "Repos/agral/CompetitiveProgramming/HackerRank"
)
DOMAIN_URL = "https://www.hackerrank.com"


def eprint(*args, **kwargs) -> None:
    """Prints the message to the standard error output"""
    print(*args, file=sys.stderr, **kwargs)


def abort(msg: str = "", retcode: int = 1) -> None:
    """Prints out the error message,
    then terminates the program using the retcode as exit value."""
    eprint(f"Error: {msg}")
    eprint("Aborting.")
    sys.exit(retcode)


def normalize(string: str) -> str:
    """Returns a copy of the input `string` converted to PascalCase,
    dropping all the non-alphanumeric characters."""
    def to_pascal_case(string: str, separator: str = ' ') -> str:
        """Converts "some example text" to "SomeExampleText",
        splitting the input string by the given separator (default: " ")."""
        return "".join(word.capitalize() for word in string.split(separator))

    def to_alnum(string: str) -> str:
        """Returns a copy of the input `string`
        consisting of alphanumeric characters only."""
        return re.sub(r"\W+", "", string)

    return to_alnum(to_pascal_case(string))


def fix_trailing_newline(path) -> None:
    """Sanitizes a file by ensuring that it has trailing newline.
    POSIX mandates it, but the testcases at HackerRank do not have them,
    and thus need to be sanitized before using."""
    contents = open(path, "r").readlines()
    if not contents[-1].endswith("\n"):
        contents[-1] = contents[-1] + "\n"
        with open(path, "w") as outfile:
            outfile.writelines(contents)


def url_to_canonical_name(url: str) -> str:
    """Extracts a canonical name from a full link to a HackerRank challenge.
    Sample url: https://www.hackerrank.com/challenges/simple-array-sum/problem
    Expected output: simple-array-sum"""
    part_after_challenges = url.partition("/challenges/")[2]
    if "/" in part_after_challenges:
        part_after_challenges = part_after_challenges.partition("/")[0]
    return part_after_challenges


class HackerRankScraper():
    """Parses a challenge link obtaining useful info"""
    def __init__(self, params) -> None:
        self.headers = {
            "User-Agent": "Agral's HackerRankInit script",
        }
        self.params = params
        self.response = requests.get(url=self.params.url, headers=self.headers)
        if self.response.status_code != 200:
            abort(f"server returned {self.response.status_code}")

        self.soup = BeautifulSoup(self.response.text, "html.parser")
        self.parse_breadcrumb_path()

    def parse_breadcrumb_path(self) -> None:
        """Reads a path that corresponds to the HR's classification
        of the type of this particular challenge, e.g. "Algorithms/Search".
        Updates the values of members `breadcrumb_path`, `challenge_name`
        and `challenge_dir`."""
        print("Extracting breadcrumb path... ")
        breadcrumb_div = self.soup.find("div", {
            "class": "community-header-breadcrumb-items"})
        if not breadcrumb_div:
            abort("Could not find a breadcrumb div in the response.")

        breadcrumbs = breadcrumb_div.find_all("span", {
            "class": "breadcrumb-item-text"})
        self.breadcrumb_path = "/".join(
            [normalize(b.text) for b in breadcrumbs])

        print("Creating challenge directory... ")
        self.challenge_dir = os.path.join(
            HACKERRANK_SOLUTIONS_DIR, self.breadcrumb_path)
        try:
            os.makedirs(self.challenge_dir, exist_ok=self.params.dir_ok)
        except FileExistsError:
            abort(
                f"Directory [{self.challenge_dir}] already exists.\n" +
                "Consider passing -d|--dir_ok flag to allow existing dirs.")
        self.challenge_name = breadcrumbs[-1].text
        self.normalized_name = normalize(self.challenge_name)

    def download_challenge_pdf(self):
        """Downloads the PDF file with the challenge's description"""
        print("Downloading a PDF with challenge's description...")
        pdf_tag = self.soup.find("a", {"id": "pdf-link"})
        if not pdf_tag:
            abort("Could not find a link to the PDF file")
        pdf_url = DOMAIN_URL + pdf_tag.get("href")
        pdf_path = os.path.join(
            self.challenge_dir, f"{self.normalized_name}.pdf")
        pdf_request = requests.get(
            pdf_url, allow_redirects=True, headers=self.headers)
        open(pdf_path, "wb").write(pdf_request.content)

    def download_testcases(self):
        """Downloads sample testcases associated with the challenge"""
        print("Downloading sample testcases archive...")
        testcases_tag = self.soup.find("a", {"id": "test-cases-link"})
        if not testcases_tag:
            abort("Could not find a link to the sample testcases.")
        testcases_url = DOMAIN_URL + testcases_tag.get("href")
        testcases_dir = os.path.join(self.challenge_dir, "testcases")
        zip_path = os.path.join(testcases_dir, "testcases.zip")
        os.makedirs(testcases_dir, exist_ok=self.params.dir_ok)
        testcases_request = requests.get(
            testcases_url, allow_redirects=True, headers=self.headers)
        open(zip_path, "wb").write(testcases_request.content)

        print("Unzipping the downloaded archive...")
        with zipfile.ZipFile(zip_path, "r") as zip_ref:
            zip_ref.extractall(testcases_dir)
        os.remove(zip_path)

        print("Making the extracted testcases POSIX-conformant...")
        for root, _, files in os.walk(testcases_dir):
            for filename in files:
                fix_trailing_newline(os.path.join(root, filename))

    def create_answer_file_cpp(self):
        """Creates a skeleton CPP answer file with predefined contents."""
        file_path = os.path.join(
            self.challenge_dir,
            "cpp",
            f"{self.normalized_name}.cpp")
        if os.path.exists(file_path):
            print("C++ answer file already exists, not overwriting.")
            return

        # Makes the parent directory:
        os.makedirs(os.path.dirname(file_path), exist_ok=True)

        current_date_string = f"{date.today():%d.%m.%Y}"
        print(f"Creating a C++ answer file ({self.normalized_name}.cpp)...")
        with open(file_path, "w") as file_handle:
            file_handle.writelines([
                "/**\n"
                f" * Solution to the \"{self.challenge_name}\" ",
                "challenge from HackerRank:\n",
                f" * {self.params.url}\n"
                f" * Created on: {current_date_string}\n",
                f" * Last modified: ${current_date_string}\n",
                " * Author: Adam Graliński (adam@gralin.ski)\n",
                " * License: MIT\n",
                " */\n",
                "\n",
                "#include <iostream>\n",
                "\n",
                "int main(int, char**)\n",
                "{\n",
                "  return 0;\n",
                "}\n",
                ])

    def create_answer_file_java(self):
        """Creates a skeleton JAVA answer file with predefined contents."""
        file_path = os.path.join(
            self.challenge_dir,
            "java",
            f"{self.normalized_name}.java")
        if os.path.exists(file_path):
            print("Java answer file already exists, not overwriting.")
            return

        # Makes the parent directory:
        os.makedirs(os.path.dirname(file_path), exist_ok=True)

        current_date_string = f"{date.today():%d.%m.%Y}"
        print(f"Creating a Java answer file ({self.normalized_name}.java)...")
        with open(file_path, "w") as file_handle:
            file_handle.writelines([
                "/*\n"
                f" * Solution to the \"{self.challenge_name}\" ",
                "challenge from HackerRank:\n",
                f" * {self.params.url}\n"
                f" * Created on: {current_date_string}\n",
                f" * Last modified: ${current_date_string}\n",
                " * Author: Adam Graliński (adam@gralin.ski)\n",
                " * License: MIT\n",
                " */\n",
                "\n",
                "import java.io.*;\n",
                "import java.util.*;\n",
                "\n",
                f"public class {self.normalized_name}\n",
                "{\n",
                "  private static final Scanner scanner = ",
                "new Scanner(System.in);\n",
                "\n",
                "  public static void main(String[] args) {\n",
                "    \n",
                "  }\n",
                "}\n",
                ])


def main():
    """An entry point to the application."""
    argparser = argparse.ArgumentParser(
        description="Bootstraps an answer to a HackerRank challenge")
    argparser.add_argument(
        "-d", "--dir_ok", action="store_true",
        help="append to/overwrite the challenge directory if it exists")
    argparser.add_argument(
        "url", action="store", type=str,
        help="the URL of the challenge in hackerrank.com domain")

    args = argparser.parse_args()

    if not args.url:
        abort("Challenge URL has not been provided.")

    print(f"Initializing ...")
    scraper = HackerRankScraper(args)
    scraper.download_challenge_pdf()
    scraper.download_testcases()
    scraper.create_answer_file_cpp()
    scraper.create_answer_file_java()


if __name__ == "__main__":
    main()
