# Contributing

Thanks for your interest in contributing to btrade!

## Code of Conduct

Help us keep btrade open and inclusive. Please read and follow our [Code of Conduct](CODE_OF_CONDUCT.md).

## Getting Started

* make sure you have a [GitHub account](https://github.com/signup/free)
* submit a ticket for your issue, assuming one does not already exist
    * clearly describe the issue including steps to reproduce when it is a bug
    * identify specific versions of the binaries and client libraries
* fork the repository on GitHub

## Making Changes

* create a branch from where you want to base your work
    * we typically name branches according to the following format: `helpful_name_<issue_number>`
* make commits of logical units
* make sure your commit messages are in a clear and readable format, example:

```
btraded: fixed bug in broker x
  
* update the message pump to properly account for RDYness
* cleanup variable names
* ...
```

* if you're fixing a bug or adding functionality it probably makes sense to write a test
* make sure to run `fmt.sh` in the root of the repo to ensure that your code is
  properly formatted and that tests pass (we use GitHub Actions for continuous integration)

## Submitting Changes

* push your changes to your branch in your fork of the repository
* submit a pull request against btrade's repository
* comment in the pull request when you're ready for the changes to be reviewed: `"ready for review"`