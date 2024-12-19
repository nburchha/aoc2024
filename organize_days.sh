#!/bin/bash

# Define the source directory
src_dir="./src"

# Create folders day01 to day19
for i in $(seq -w 1 19); do
    folder_name="day$i"
    mkdir -p "$folder_name"
done

# Move corresponding .go files from the src directory into their respective folders
for file in "$src_dir"/day*.go; do
    # Extract the base name without the extension
    base_name="${file%.go}"
    base_name="${base_name##*/}" # Remove the path to get just the filename
    # Check if the corresponding folder exists
    if [ -d "$base_name" ]; then
        mv "$file" "$base_name/"
    else
        echo "No matching folder for $file. Skipping."
    fi
done

echo "Folders created and .go files moved successfully."
