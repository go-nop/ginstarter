#!/bin/sh

read -p "Enter new go package name (i.e. github.com/{username}/{repo_name}): " pkg_name

if [ -z "$pkg_name" ]; then
    echo "Error: package name cannot be empty"
    exit 1
fi

repo_name=$(basename "$pkg_name")

# Update go.mod
echo "Updating go.mod"
go mod edit -module "$pkg_name"

# Convert all imports in code
echo "Converting all imports in code"
escaped_pkg_name=$(echo "$pkg_name" | sed 's/\//\\\//g')
old_pkg_name=$(echo 'github.com/go-nop/ginstarter' | sed 's/\//\\\//g')
find . -name "*.go" -print0 | xargs -0 sed -i -e "s/${old_pkg_name}/${escaped_pkg_name}/g"

# Delete backup files created by sed
echo "Deleting backup files"
find . -type f -name "*.go-e" -delete

# Change git origin
echo "Changing git origin"
git remote remove origin
git remote add origin "https://$pkg_name"

# Move the repository folder
current_dir=$(basename "$PWD")
parent_dir=$(dirname "$PWD")
new_dir="$parent_dir/$repo_name"

echo "Moving $current_dir to $new_dir"
mv "$PWD" "$new_dir"

echo "New Folder: $new_dir"
