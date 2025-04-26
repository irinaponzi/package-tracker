-- This script creates the tables for the package tracking system

-- Database: package-tracker
DROP DATABASE IF EXISTS package_tracker;
CREATE DATABASE package_tracker;
USE package_tracker;

-- Create tables
CREATE TABLE tracking_codes (
                                id CHAR(36) PRIMARY KEY,
                                country CHAR(2) NOT NULL,
                                date DATE NOT NULL,
                                sequence INT NOT NULL,
                                code VARCHAR(50) NOT NULL UNIQUE,
                                status ENUM('generated', 'assigned', 'cancelled') NOT NULL DEFAULT 'generated',
                                created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
                                updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE packages (
                          id CHAR(36) PRIMARY KEY,
                          tracking_code_id CHAR(36) NOT NULL,
                          status ENUM('created', 'in_transit', 'delivered', 'lost', 'returned') NOT NULL DEFAULT 'created',
                          size ENUM('small', 'medium', 'large') NOT NULL,
                          weight_kg DECIMAL(6,2) NOT NULL,
                          destination VARCHAR(255) NOT NULL,
                          created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
                          updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                          FOREIGN KEY (tracking_code_id) REFERENCES tracking_codes(id)
);

CREATE TABLE batches (
                         id CHAR(36) PRIMARY KEY,
                         quantity INT NOT NULL,
                         transport_company VARCHAR(255) NULL,
                         status ENUM('created', 'done', 'fail') NOT NULL DEFAULT 'created',
                         error_details VARCHAR(255),
                         created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE batch_packages (
                                id CHAR(36) PRIMARY KEY,
                                batch_id CHAR(36) NOT NULL,
                                package_id CHAR(36) NOT NULL,
                                FOREIGN KEY (batch_id) REFERENCES batches(id),
                                FOREIGN KEY (package_id) REFERENCES packages(id),
                                UNIQUE (batch_id, package_id)
);

-- Some insert statements to populate the tables with dummy data
INSERT INTO tracking_codes (id, country, date, sequence, code) VALUES
('tracking-uuid-001', 'AR', '2025-04-26', 1, 'AR-20250426-000001'),
('tracking-uuid-002', 'BR', '2025-04-26', 2, 'BR-20250426-000002'),
('tracking-uuid-003', 'US', '2025-04-26', 3, 'US-20250426-000003'),
('tracking-uuid-004', 'MX', '2025-04-26', 4, 'MX-20250426-000004'),
('tracking-uuid-005', 'ES', '2025-04-26', 5, 'ES-20250426-000005'),
('tracking-uuid-006', 'AR', '2025-04-26', 6, 'AR-20250426-000006'),
('tracking-uuid-007', 'BR', '2025-04-26', 7, 'BR-20250426-000007'),
('tracking-uuid-008', 'US', '2025-04-26', 8, 'US-20250426-000008'),
('tracking-uuid-009', 'MX', '2025-04-26', 9, 'MX-20250426-000009'),
('tracking-uuid-010', 'ES', '2025-04-26', 10, 'ES-20250426-000010');

INSERT INTO packages (id, tracking_code_id, size, weight_kg, destination)
VALUES
('package-uuid-001', 'tracking-uuid-001', 'small', 1.25, 'Buenos Aires'),
('package-uuid-002', 'tracking-uuid-002', 'medium', 3.75, 'Cordoba'),
('package-uuid-003', 'tracking-uuid-003', 'large', 8.20, 'Santa Fe'),
('package-uuid-004', 'tracking-uuid-004', 'small', 0.95, 'Mendoza'),
('package-uuid-005', 'tracking-uuid-005', 'medium', 4.10, 'La Plata'),
('package-uuid-006', 'tracking-uuid-006', 'large', 9.50, 'Tucuman'),
('package-uuid-007', 'tracking-uuid-007', 'small', 1.10, 'Rosario'),
('package-uuid-008', 'tracking-uuid-008', 'medium', 3.20, 'Salta'),
('package-uuid-009', 'tracking-uuid-009', 'large', 10.00, 'Neuquen'),
('package-uuid-010', 'tracking-uuid-010', 'small', 0.85, 'Rio Cuarto');

INSERT INTO batches (id, quantity, transport_company)
VALUES
('batch-uuid-001', 2, 'FastDelivery'),
('batch-uuid-002', 2, 'ExpressLogistics'),
('batch-uuid-003', 2, 'QuickShip'),
('batch-uuid-004', 2, 'SpeedyTransport'),
('batch-uuid-005', 2, 'AirCargo');

INSERT INTO batch_packages (id, batch_id, package_id)
VALUES
('batchpackage-uuid-001', 'batch-uuid-001', 'package-uuid-001'),
('batchpackage-uuid-002', 'batch-uuid-001', 'package-uuid-002'),
('batchpackage-uuid-003', 'batch-uuid-002', 'package-uuid-003'),
('batchpackage-uuid-004', 'batch-uuid-002', 'package-uuid-004'),
('batchpackage-uuid-005', 'batch-uuid-003', 'package-uuid-005'),
('batchpackage-uuid-006', 'batch-uuid-003', 'package-uuid-006'),
('batchpackage-uuid-007', 'batch-uuid-004', 'package-uuid-007'),
('batchpackage-uuid-008', 'batch-uuid-004', 'package-uuid-008'),
('batchpackage-uuid-009', 'batch-uuid-005', 'package-uuid-009'),
('batchpackage-uuid-010', 'batch-uuid-005', 'package-uuid-010');
