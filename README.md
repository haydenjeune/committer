# Committer

Committer is a small command line utility that makes it easy to set the details
that will be used to commit to the current git repo. This is useful if you have
different github accounts for uni, work, and personal things.

## Installation

There are 3 ways to install `committer`:

1) Precompiled binaries

Binaries for Mac (AMD64), Windows(AMD64), and Linux(AMD64 and ARM64) are attached to the release. Just download them, and put them somewhere on your `PATH` to access `committer` from the command line.

2) Go get

With go installed, `committer` can be installed with:

```bash
go get github.com/haydenjeune/committer
```

3) Homebrew

On Mac, you can use Homebrew to manage your installation. 

First add my tap to your homebrew.

```brew tap haydenjeune/tap```

Then install `committer` form this tap.

```brew install haydenjeune/tap/committer```


## Usage

### Adding new author details

```
committer add <profile>
```

Committer will then prompt you for a name and email to save against the given profile name.

### Setting commit author on a repo

```
committer set <profile>
```

This will set the commit author of the repository that you are currently in to my details. `<profile>` must be a profile that has been added with the
`committer add <profile>` command.

### Showing saved profiles

```
committer list
```

This will show a list of saved profiles.

### Removing saved profiles

```
committer rm <profile>
```

This will remove the profile with the given name.
