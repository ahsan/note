# note

Easily add new and read back old notes on the CLI.

## Installation

1. Clone repository locally.
2. Make the `note` script executable and add its location to path.
3. Create a repository to house all of your notes, e.g `~/.notes/`
```
mkdir ~/.notes/
```
4. Export the notes directory to your `bashrc` as `NOTES_DIR` environment variable.
```
# .bashrc
export NOTES_DIR=~/projects/note/.notes/
```

## Usage

1. Run `note` to see the help menu.
2. You can sync your notes across multiple computers using git with a remote repository (e.g on GitHub) for free. 
```
# Initializing git repo
cd $NOTES_DIR
git init
git remote add origin <remote git repo url>

# Pushing new notes to origin
cd $NOTES_DIR
git add -u
git commit -m "Added notes about <topic 1>, <topic 2>, ..."
git push

# Syncing notes
cd $NOTES_DIR
git pull
```
