# Meta Instructions - REMOVE THIS PART

After duplicating or using this template

- Remove directories not required
- Update readme
- Update LICENSE date & name
- Update makefile, and uncomment suggested commands
- Update / remove contents of api directory
- Edit .github/workflows
- Place code in src
- Use sample configs for linting and other tools in sample folder

# Project Title

Purpose and description of this project

Goals:

- Make a thing
- Do a thing

Use cases & key features:

- Something
- Something else

Supporting technologies and libraries:

- Stuff
- Things

<!-- Note! Change benc-uk/project-starter for the real repo!! -->
<!-- See https://shields.io/ for more -->

![](https://img.shields.io/github/license/benc-uk/project-starter)
![](https://img.shields.io/github/last-commit/benc-uk/project-starter)
![](https://img.shields.io/github/release/benc-uk/project-starter)
![](https://img.shields.io/github/checks-status/benc-uk/project-starter/main)
![](https://img.shields.io/github/workflow/status/benc-uk/project-starter/CI%20Build?label=ci-build)
![](https://img.shields.io/github/workflow/status/benc-uk/project-starter/Release%20Assets?label=release)

# Table Of Contents

Optional. Remove TOC for smaller projects

# Getting Started

## Installing / Deploying

- If the project can be installed (such as a command line tool or library)
- Or deployed to Kubernetes, public cloud etc

## Running as container

Notes on running the project from Docker image / container

## Running locally

Notes on running the project locally, including pre-reqs

# Architecture

Optional. Diagram or description of the overall system architecture, only where applicable.

# Screenshots

Optional. Screenshots can help convey what the project looks like when running and what it's purpose and use is.

# Configuration

Details of any configuration files, environmental variables, command line parameters, etc.

For services
| Setting / Variable | Purpose | Default |
| ------------------ | ------------------------------------------- | ------- |
| PORT | Port the server will listen on. | 8000 |
| SOMETHING | Some very important setting. **_Required_** | _None_ |
| SOMETHING_ELSE | Some less important setting | "foo" |

Example for CLI tools

```bash
./foo-tool --help

Options:
  -p, --preset <presetName>       Skip prompts and use saved or remote preset
  -d, --default                   Skip prompts and use default preset
```

# Repository Structure

A brief description of the top-level directories of this project is as follows:

```c
/api        - Details of the API specification & docs
/build      - Build configuration e.g. Dockerfiles
/charts     - Helm charts
/deploy     - Deployment and infrastructure as code, inc Kubernetes
/scripts    - Bash and other supporting scripts
/src        - Source code
/test       - Testing, mock data and API + load tests
```

# API

See the [API documentation](./api/) for full infomration about the API(s).  
Optional. Delete this section if project as no API.

# Known Issues

List any known bugs or gotchas.

# Change Log

See [complete change log](./CHANGELOG.md)

# License

This project uses the MIT software license. See [full license file](./LICENSE)

# Acknowledgements

Optional. Put acknowledgements and credits here, if any
