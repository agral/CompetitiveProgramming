"""Sets up directory structure to solve a challenge posted on HackerRank by:
  * Fetching the challenge description (PDF file),
  * Fetching the sample testcases and answers to them,
  * Populating an answer file (.cpp) with a preconfigured text"""

import argparse
import os
import re
import sys
import zipfile
from datetime import date
import requests
from bs4 import BeautifulSoup

HACKERRANK_DIR = os.path.join(
    os.path.expanduser("~"),
    "Repos/agral/CompetitiveProgramming/HackerRank")
DOMAIN_URL = "http://www.hackerrank.com"


def url_to_canonical_name(url: str) -> str:
    """Extracts a canonical name from a full link to a challenge.
    sample url: https://www.hackerrank.com/challenges/count-luck/problem
    expected output: count-luck"""
    part_after_challenges = url.partition("/challenges/")[2]
    if "/" in part_after_challenges:
        part_after_challenges = part_after_challenges.partition("/")[0]
    return part_after_challenges


def to_pascal_case(string: str, separator: str = ' ') -> str:
    """Converts `some standard name` to SomeStandardName,
    splitting the input string by a given separator"""
    return "".join(word.capitalize() for word in string.split(separator))


def to_alnum(string: str) -> str:
    """Returns a string that is a copy of `string` with all non-alphanumeric
    characters removed."""
    return re.sub(r'\W+', '', string)


def sanitize_file(path) -> None:
    """Sanitizes a file by ensuring that it has trailing newline.
    POSIX mandates it, but the testcases at HackerRank do not have them,
    and thus need to be sanitized before using."""
    contents = open(path, "r").readlines()
    if not contents[-1].endswith("\n"):
        contents[-1] = contents[-1] + "\n"
    with open(path, "w") as outfile:
        outfile.writelines(contents)


class ChallengeParser():
    """Parses a challenge link obtaining files and useful info"""
    def __init__(self, args) -> None:
        """Creates a new ChallengeParser instance"""
        self.args = args
        self.url_challenge = args.url_challenge
        print("Inspecting problem's webpage...")
        self.response_challenge = requests.get(self.url_challenge)
        self.soup_challenge = BeautifulSoup(
            self.response_challenge.text, "html.parser")
        self.read_breadcrumb_path()
        self.create_challenge_directory()

    def read_breadcrumb_path(self) -> None:
        """Reads a path that corresponds to the HR's classification
        of the type of this particular challenge, e.g. "Algorithms/Search".
        Updates the values of members `dir_challenge` and `breadcrumb_path`.
        """
        print("Extracting path from breadcrumb items...")
        breadcrumb_div = self.soup_challenge.find("div", {
            "class": "community-header-breadcrumb-items"})
        breadcrumbs = breadcrumb_div.find_all("span", {
            "class": "breadcrumb-item-text"})
        self.breadcrumb_path = "/".join(
            [to_alnum(to_pascal_case(x.text, " ")) for x in breadcrumbs])
        self.name_challenge = breadcrumbs[-1].text
        self.name_challenge_pascalcase = to_alnum(
            to_pascal_case(self.name_challenge))
        self.dir_challenge = os.path.join(HACKERRANK_DIR, self.breadcrumb_path)
        print(self.dir_challenge)

    def create_challenge_directory(self):
        """Creates a root directory for the answer to the challenge"""
        print(f"Creating a challenge directory (path: {self.breadcrumb_path})")
        try:
            os.makedirs(
                self.dir_challenge, exist_ok=self.args.allow_existing_dir)
        except FileExistsError:
            print(f"Error: directory {self.breadcrumb_path} already exists.")
            print(f"Full path: {self.dir_challenge}")
            print("Aborting")
            sys.exit(1)

    def download_challenge_pdf(self):
        """Downloads the PDF file with the challenge's description"""
        print("Searching for a link to PDF with challenge description...")
        pdf_url_tag = self.soup_challenge.find("a", {"id": "pdf-link"})
        if not pdf_url_tag:
            print("Error: Could not find a link to the PDF")
            print("Aborting.")
            sys.exit(1)

        pdf_url = DOMAIN_URL + pdf_url_tag.get("href")
        print(f"Downloading {pdf_url}")
        pdf_request = requests.get(pdf_url, allow_redirects=True)
        pdf_path = os.path.join(
            self.dir_challenge, f"{self.name_challenge_pascalcase}.pdf")
        open(pdf_path, "wb").write(pdf_request.content)

    def download_testcases(self):
        """Downloads the sample testcases associated with the challenge"""
        print("Searching for a link to the archive with testcases... ", end="")
        testcases_url_tag = self.soup_challenge.find(
            "a", {"id": "test-cases-link"})
        if not testcases_url_tag:
            print("FAILED")
            print("Error: Could not find a link to the sample testcases")
            print("Aborting.")
            sys.exit(1)
        print("OK")

        testcases_url = DOMAIN_URL + testcases_url_tag.get("href")
        testcases_path = os.path.join(self.dir_challenge, "testcases")
        os.makedirs(testcases_path, exist_ok=True)
        print(f"Downloading {testcases_url}... ", end="")
        testcases_request = requests.get(testcases_url, allow_redirects=True)
        zip_path = os.path.join(testcases_path, "testcases.zip")
        open(zip_path, "wb").write(testcases_request.content)
        print("OK")

        print("Unzipping the downloaded archive... ", end="")
        with zipfile.ZipFile(zip_path, "r") as zip_ref:
            zip_ref.extractall(testcases_path)
        print("OK")

        print("Removing the archive file... ", end="")
        os.remove(zip_path)
        print("OK")

        print("Sanitizing the testcase files...")
        for root, _, files in os.walk(testcases_path):
            for filename in files:
                sanitize_file(os.path.join(root, filename))

    def create_answer_file_cpp(self):
        """Creates a skeleton answer file (ans.cpp) with predefined contents"""
        absolute_path = os.path.join(self.dir_challenge, "ans.cpp")
        if os.path.exists(absolute_path):
            print("C++ answer file already exists, not overwriting.")
            return

        current_date_string = f"{date.today():%d.%m.%Y}"
        print("Creating a C++ answer file (ans.cpp)...")
        with open(absolute_path, "w") as file_handle:
            file_handle.writelines([
                "/**\n",
                f" * Solution to the \"{self.name_challenge}\"",
                " challenge form HackerRank:\n",
                f" * {self.url_challenge}\n",
                f" * Created on: {current_date_string}\n",
                f" * Last modified: {current_date_string}\n",
                " * Author: Adam Gralinski (adam@gralin.ski)\n",
                " * License: MIT\n",
                " */\n",
                "\n",
                "#include <iostream>\n",
                "\n",
                "int main(int, char**)\n",
                "{\n",
                "  testcases\n",
                "  return 0;\n",
                "}\n",
                ])

    def create_makefile(self):
        """Creates a simple Makefile with options to compile, run and clean
        the answer file. Also contains an option to verify the correctness
        of the answers given to the sample testcases."""
        makefile_path = os.path.join(self.dir_challenge, "Makefile")
        if os.path.exists(makefile_path):
            print("Makefile already exists, not overwriting.")
            return
        print("Creating a Makefile...")
        with open(makefile_path, "w") as makefile:
            makefile.writelines([
                ".PHONY: all clean run verify\n\n",
                "CXX=clang++\n",
                "CXXFLAGS=-g -Wall -Wextra -Wshadow -Wnon-virtual-dtor ",
                "-pedantic -D_REENTRANT -std=c++14 -O0\n",
                "DIR_BIN=bin\n",
                "DIR_TESTCASES_IN=testcases/input\n",
                "DIR_TESTCASES_OUT=testcases/output\n",
                "EXE_APP=exe\n\n",
                "all: $(DIR_BIN)/$(EXE_APP)\n\n",
                "clean:\n\trm -rf $(DIR_BIN)\n\n",
                "run: $(DIR_BIN)/$(EXE_APP)\n\t@$(DIR_BIN)/$(EXE_APP)\n\n",
                "verify: $(DIR_BIN)/$(EXE_APP)\n\t@VerifyHrAnswer.sh\n\n",
                "$(DIR_BIN)/$(EXE_APP): $(DIR_BIN)/ans.o\n",
                "\t@mkdir -p $(@D)\n\t$(CXX) $(CXXFLAGS) $< -o $@\n\n",
                "$(DIR_BIN)/ans.o: ans.cpp\n",
                "\t@mkdir -p $(@D)\n\t$(CXX) -c $(CXXFLAGS) $< -o $@\n"
            ])


def main():
    """A main method"""
    argparser = argparse.ArgumentParser(
        description="Bootstraps an answer to a HackerRank challenge")
    argparser.add_argument(
        "url_challenge", action="store", type=str,
        help="The URL of the challenge in hackerrank.com domain")
    argparser.add_argument(
        "--allow_existing_dir", "-d", action="store_true",
        help="Update the challenge directory if it already exists")

    args = argparser.parse_args()
    parser = ChallengeParser(args)
    parser.download_challenge_pdf()
    parser.download_testcases()
    parser.create_answer_file_cpp()
    parser.create_makefile()
    print("To enter the challenge's directory use the following command:")
    print(f"cd {parser.dir_challenge}")


if __name__ == "__main__":
    main()
