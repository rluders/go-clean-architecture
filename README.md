# Task Manager: Clean Architecture Demonstration

Welcome to the **Task Manager** repository! This project is a hands-on demonstration of 
Clean Architecture principles applied to a simple yet versatile application—a 
Task Management system. It's designed as a companion to the article 
**"[Demystifying Clean Architecture in Go: Separating Fact from Fiction](http://github.com)"**.

## Purpose

This repository showcases how Clean Architecture can adapt to different project needs and complexities. It provides two examples:
1. **Flat Project**: A minimalistic approach for smaller projects, combining layers for simplicity while maintaining core principles.
2. **Elaborated Project**: A modular, scalable design suitable for larger projects, with clear separation of responsibilities.

By comparing these two projects, you'll understand how Clean Architecture balances clarity, scalability, and Go’s philosophy of simplicity.

## Features

- **Dual Interfaces**:
    - REST API
    - Command-Line Interface (CLI)
- **Clean Architecture Layers**:
    - Domain (Business Logic)
    - Use Cases (Application Logic)
    - Infrastructure (Delivery Mechanisms and Persistence)

## Getting Started

To explore the examples:
1. Clone this repository.
2. Navigate to the desired project (`flat-project` or `elaborated-project`).
3. Run the corresponding entry points (`cmd/cli/main.go` for CLI, `cmd/rest/main.go` for REST API).

## License

This project is licensed under the MIT License. See the `LICENSE` file for details.
