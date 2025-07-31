#!/usr/bin/env python3

import os
from pathlib import Path

def modify_file(file_path):
    try:
        with open(file_path, 'r', encoding='utf-8') as file:
            content = file.readlines()

        modified_content = []
        for line in content:
            modified_content.append(line.replace("//LOG.", "LOG."))

        # Write the modified content back to the file
        with open(file_path, 'w', encoding='utf-8') as file:
            file.writelines(modified_content)

        print(f"Successfully modified: {file_path}")
    except Exception as e:
        print(f"Error processing {file_path}: {str(e)}")


def modify_files_in_directory(directory_path, extension):
    # Ensure the directory exists
    if not os.path.isdir(directory_path):
        print(f"Error: {directory_path} is not a valid directory")
        return
    # Get all files with the specified extension
    for filename in os.listdir(directory_path):
        if filename.endswith(extension):
            file_path = os.path.join(directory_path, filename)
            modify_file(file_path)


def main():
    directory = Path(__file__).resolve().parent.joinpath("rsvp")
    extension = ".go"

    modify_files_in_directory(directory, extension)


if __name__ == "__main__":
    main()
