-- Schema for the recordstore Granitic example
-- This file was tested against MySQL 5.5 running in ANSI mode, so should be relatively easy to adapt for other RDBMSs

DROP DATABASE IF EXISTS recordstore;
CREATE DATABASE recordstore;

USE recordstore;

-- These users and privileges are not appropriate for production configurations.
DROP USER 'api'@'localhost';
CREATE USER 'api'@'localhost' IDENTIFIED BY 'apipass';

DROP USER 'api'@'%';
CREATE USER 'api'@'%' IDENTIFIED BY 'apipass';

GRANT ALL PRIVILEGES ON recordstore.* TO 'api'@'localhost';
GRANT ALL PRIVILEGES ON recordstore.* TO 'api'@'%';

CREATE TABLE artist (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(64),
    UNIQUE(name)
);

CREATE TABLE track (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(64)
);

CREATE TABLE record (
    id INT AUTO_INCREMENT PRIMARY KEY,
    cat_ref CHAR(10) NOT NULL UNIQUE,
    name VARCHAR(128),
    artist_id INT,
    FOREIGN KEY (artist_id) REFERENCES artist(id)
);

CREATE TABLE format (
    id INT PRIMARY KEY,
    name VARCHAR(16)
);

INSERT INTO format (id, name) VALUES
(0, 'LP'),
(1, 'CD'),
(2, 'DOWNLOAD');

CREATE TABLE record_format_price (
    record_id INT,
    format_id INT,
    price INT,
    FOREIGN KEY (record_id) REFERENCES record(id),
    FOREIGN KEY (format_id) REFERENCES format(id),
    PRIMARY KEY (record_id, format_id)
);


CREATE TABLE record_track (
    track_id INT,
    record_id INT,
    track_number INT,
    FOREIGN KEY (track_id) REFERENCES track(id),
    FOREIGN KEY (record_id) REFERENCES record(id),
    UNIQUE (track_id, record_id, track_number)
);

CREATE TABLE stock (
    record_id INT,
    format_id INT,
    stock INT,
    FOREIGN KEY (record_id) REFERENCES record(id),
    FOREIGN KEY (format_id) REFERENCES format(id),
    PRIMARY KEY (record_id, format_id)
);

