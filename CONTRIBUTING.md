# **Contributing to ASIST**

Thank you for your interest in contributing to ASIST\! We welcome contributions from the community to help make this tool even better.

## **Code of Conduct**

Please note that this project is released with a [Contributor Covenant Code of Conduct](https://github.com/certinia/asist/blob/main/CODE_OF_CONDUCT.md). By participating in this project, you agree to abide by its terms.

## **How Can I Contribute?**

Here are many ways you can contribute:

* **Reporting Bugs:** Help us identify and fix issues by reporting them clearly and in detail.  
* **Suggesting Enhancements:** Propose new features or improvements to existing ones.  
* **Writing Code:** Contribute new rules, improve the engine, or add features.  
* **Improving Documentation:** Help us make the documentation clearer, more concise, and more helpful.  
* **Reviewing Code:** Review pull requests from other contributors.

## **Getting Started**
Refer [DEVELOPING.md](https://github.com/certinia/asist/blob/main/DEVELOPING.md)

## **Contributing Code**

1. **Write Your Code:** Implement your feature, bug fix, or enhancement.  
2. **Follow Go Coding Standards:** Ensure your code adheres to the [Go Code Review Comments](https://go.dev/wiki/CodeReviewComments) and is formatted with gofmt.  
3. Naming conventions: Variable and method names should be in Camel case, and constants should be in uppercase.  
4. Comment when necessary, but avoid over-commenting, preferring code that is easily interpreted   
5. Avoid global access modifier usage.  
   * These will be locked into managed packages forever.  
6. When declaring variables with var, they should be defined at the start of the method instead of being defined during execution  
7. Follow Go Idioms \-   
   * Use ":=" for local variable declaration  
   * Use short, descriptive variable names  
8. Error Handling  
   * Return values and errors from the method instead of logging errors inside the method.  
9. Segregation of Functions   
   * Segregate complex logic into multiple functions  
10. Use Structs and Interfaces to encapsulate behavior  
11. Check to use Go’s Standard Library Effectively   
    * Try to use standard library functions instead of creating custom ones wherever possible  
    * Avoid introducing additional third-party libraries unless explicitly needed  
12. **Write Tests:** Write multiple unit-test cases for methods to cover all the scenarios  
13. **Run Tests:** Make sure all tests pass:  
    go test ./...
14. **Build and Run:** Build and run ASIST locally to verify your changes. Refer [DEVELOPING.md](https://github.com/certinia/asist/blob/main/DEVELOPING.md)
15. **Commit Your Changes:**  
    git add .  
    git commit \-m "feat: Add new feature"

    * Use clear and descriptive commit messages, following the [Conventional Commits](https://www.conventionalcommits.org/en/v1.0.0/) specification.  
16. **Push to Your Fork:**  
    git push origin my-feature-branch

17. **Create a Pull Request:** Submit a pull request (PR) from your branch to the main branch of the original repository.  
    * Provide a clear title and description of your changes in the PR.  
    * Reference any related issues.

## **Reporting Bugs**

1. **Check Existing Issues:** Before submitting a new issue, please check if a similar issue already exists.  
2. **Create a New Issue:** If you've confirmed it's a new issue:  
   * Go to the repository's Issues page.  
   * Click "New issue".  
   * Provide a clear and descriptive title.  
   * Include the following information in your report:  
     * Version of Go and the ASIST.  
     * Operating system and CPU architecture.  
     * Code snippet that reproduces the bug.  
     * The expected behavior and the actual behavior.  
     * Any error messages or logs.  
     * Steps to reproduce the issue.  
3. **Be Responsive:** If a maintainer asks for more information, please respond promptly.

## **Suggesting Enhancements**

1. **Check Existing Issues:** See if your suggestion has already been proposed.  
2. **Create a New Issue:** If it's a new suggestion:  
   * Go to the repository's Issues page.  
   * Click "New issue".  
   * Provide a clear and descriptive title (e.g., "Change request: Support for XYZ").  
   * Describe the proposed enhancement in detail, including:  
     * The problem it solves.  
     * The proposed solution.  
     * Any potential benefits and drawbacks.  
     * Use cases.  
   * Be open to discussion and feedback.

## **Improving Documentation**

Contributions to the documentation are highly valued\! You can help by:

* Fixing typos and grammatical errors.  
* Improving clarity and organization.  
* Adding examples and use cases.  
* Translating documentation.  
* Creating new documentation pages.

To contribute to the documentation, follow the same workflow as contributing code (fork, branch, PR).

## **Code Review Process**

* All code contributions will be reviewed by project maintainers.  
* Code reviews focus on code quality, correctness, test coverage, and adherence to coding standards.  
* You may be asked to make revisions to your code before it is merged.  
* Be patient and responsive during the review process.

## **Thank You**

Thank you for contributing to ASIST! Your contributions help make this tool better for everyone.