import os


def flatten_files(output_file):
    with open(output_file, "w") as outfile:
        for root, _, files in os.walk("."):
            for file in files:
                if file.endswith(".go"):
                    file_path = os.path.join(root, file)
                    with open(file_path, "r") as infile:
                        outfile.write(f"// {file_path}\n")
                        outfile.write(infile.read())
                        outfile.write("\n\n")


if __name__ == "__main__":
    flatten_files("flattened.go")
