<!-- markdownlint-disable MD041 MD033 -->
[![CircleCI](https://circleci.com/gh/ZupIT/ritchie-formulas/tree/ritchie-2.0.0.svg?style=shield)](https://circleci.com/gh/ZupIT/ritchie-formulas)
[![pre-commit](https://img.shields.io/badge/pre--commit-enabled-brightgreen?logo=pre-commit&logoColor=white)](https://github.com/pre-commit/pre-commit)

<img class="special-img-class" src="/docs/img/ritchie-banner.png" />

# Ritchie's commons formula repository

This repository contains the community formulas which can be executed by the [ritchie-cli](https://github.com/ZupIT/ritchie-cli).

## Create a new formula

- Fork the repository.
- Create a branch:

```bash
git checkout -b <branch_name>`
```

- Create a new formula, using the
 forked repository as a Ritchie workspace:

```bash
 rit create formula
```

 if you need help please visit [how to create formulas on Ritchie](https://docs.ritchiecli.io/getting-started/creating-formulas)

- Build and use the new formula:

```bash
rit build formula
```

 or use --watch to watch changes on formula code:

```bash
rit build formula --watch
```

- Commit your implementation.
- Push your branch.
- Open a pull request on the repository for analysis.

## Add support to other languages on create formula command

The rit create formula command use the folder `/templates/create_formula`
to list the languages options. If you like to edit some language template
or to add more language to create formula command please access
the following tutorial:
[Languages Template Tutorial](https://github.com/ZupIT/ritchie-formulas/tree/ritchie-2.0.0/templates/create_formula)

## Full Documentation

- [Gitbook](https://docs.ritchiecli.io)

## Contributing

[Contribute to the Ritchie community](https://github.com/ZupIT/ritchie-cli/blob/master/CONTRIBUTING.md)

## Zup Products

- [Zup open source](https://opensource.zup.com.br)
