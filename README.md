# Package Tracker

**Package Tracker** is a backend simulation system for managing package shipments in batches. It allows tracking the state of packages (`created`, `in_transit`, `delivered`, etc.), generating custom tracking codes, and handling batch operations with multi-database storage. This system is designed to help with learning Go architecture and transactional flows.

## Features

- **Batch Shipment Management**: Create, track, and manage shipments in bulk.
- **Tracking Codes**: Automatically generate unique tracking codes with a simple and readable format.
- **State Management**: Transition packages through various states like `created`, `in_transit`, `delivered`, and more.
- **Multi-Database Support**: Store data in both relational and fast-access databases for efficient tracking.

## Getting Started

To get started with the project:

1. Clone the repository:
   ```
   git clone https://github.com/irinaponzi/package-tracker.git
   ```
   
2. Install dependencies:
   ```
   go mod tidy
   ```
   
3. Run the application:
   ```
   go run main.go
   ```

4. Access the package tracking features and start managing shipments.

## Technologies

- **Go**: Main language used for backend logic.
- **MySQL / PostgreSQL**: Relational databases for managing batch and package data.
- **Redis / Memcached**: For caching package status and fast access to frequently updated data.

## License

This project is licensed under the GNU General Public License v3.0 (GPL-3.0).
