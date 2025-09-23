# 🛠️ Developing the ASIST Extension

Welcome to the development guide for the **ASIST** VS Code extension! This document will walk you through the steps required to get started with the development environment, build the extension, and contribute to the project.

- The source code is written in [Go Lang](https://go.dev/).
- The extension directory contains the source code for the VS Code Extension.

## 📚 Table of Contents

1. [Prerequisites](#-prerequisites)
2. [Setting Up the Development Environment](#-setting-up-the-development-environment)
3. [Building and Bundling](#%EF%B8%8F-building-and-bundling)
4. [Running the Extension Locally](#-running-the-extension-locally)
5. [Packaging the Extension](#-packaging-the-extension)

## 🔧 Prerequisites

Before you start developing, make sure you have the following tools installed:

- **Go** 1.23.0 or later: [Install GO](https://go.dev/)
- **Node.js** v22 or above: [Install Node.js](https://nodejs.org/en/)
- **[npm](https://www.npmjs.com/)**: This package manager will be used for installing dependencies
- [VS Code](https://code.visualstudio.com/) : 1.90.0 or later

Once you’ve got these ready, you’re all set to get started! 🚀

## 👨‍💻 Setting Up the Development Environment

To get started, clone this repository and install the necessary dependencies.

**Create a fork of the repository first**

**Clone the forked repository:** 
```zsh
git clone https://github.com/certinia/asist.git
cd <repo-name>
```
* Set up your local environment as described in the README. This usually involves dependencies.

* **Create a Branch:** Create a new branch for your changes:
```zsh
git checkout -b my-feature-branch
```

**Install dependencies:**

   Use [npm](https://www.npmjs.com/) to install project dependencies:

```zsh
npm i
```
- Keep the version field from package.json and stash commit tag updated with the same value.

## ⚙️ Building and Bundling

You can build the extension and prepare it for local development, run the watcher to re-build automatically for production use. Here's how:

1. **For Binary:**
```zsh
make build-binaries
```

2. **For extension:**
```zsh
make build-binaries-for-extension
```

## 🏁 Running the Extension Locally

Once you’ve built the extension or run the watcher, you can run it inside a local VS Code instance for testing and development.

1. **Start the extension host:**

   - Go to extension directory and run `npm i`
   - Compile the extension code `npx tsc -b` to compile ts code into js
   - Open the **Run and Debug** panel in VS Code (CMD/CTRL + Shift + D).
   - Select **Run Extension** from the dropdown.
   - Click the green play button to launch a new VS Code window (the extension host).

### Testing create config file

In order to use create config, we need to run `make copy-config-file` from the root of the repo.
This makes the template available to the extension.

## 🧪 Testing Your Changes

Make sure your changes don’t break anything. If you’re working on a feature or bug fix that requires tests, be sure to add or update the relevant tests.

Run Tests Locally:
If you have added or modified tests, you can run them with:

```zsh
npm run test
```

or run the tests from the test explorer in VScode

Ensure all tests pass before submitting your pull request.

## 📦 Packaging the Extension

This is for information only. Packaging and releasing are handled in GitHub.
Once you're ready to package the extension for distribution:

Package the extension:

```zsh
make build-vscode-extension-prepublish
```
This command will create a .vsix package file that you can install locally.

# Common Developer pitfalls when testing

Setting is already registered when running in debug mode. This is because your release ASIST extension is conflicting with your test version to resolve this you need to uninstall ASIST temporarily and remove it from \~/.vscode/extensions when testing is complete, simply re-install ASIST.

# 🚀 Making a release (For the maintainers)

- Ensure that the latest SBOM report is available at the root of repository before releasing.
- A release will be made by the ASIST maintainers for a specific version. We are releasing ASIST Binary and Extension separately using two different type of tags i.e. `v*.*.*` for Binary and `vsix-v*.*.*` for Extension. There are certain scenarios which need to be considered while publishing:
   - If code changes are done only in binary release, First, successfully push the binary tag format to publish binaries. Once the binary gets released successfully, push the extension format tag to publish extension's new version with latest binary code.
   - If code changes are done only in extension, push extension tag format only.<br>
   Note: Ensure that the publish build is executed successfully.
- Once the version has been successfully released, it needs to be updated in the .env file (this file is used by Makefile for providing versions to local builds) only when the changes are done in binary code.