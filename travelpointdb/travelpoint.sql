CREATE DATABASE travel_point IF NOT EXISTS;
\c travel_point;

CREATE TABLE address (
    ID VARCHAR(255) SERIAL PRIMARY KEY,
    AddressLine VARCHAR(255),
    Longitude DECIMAL(10, 6),
    Latitude DECIMAL(10, 6),
    City VARCHAR(255),
    State VARCHAR(50),
    Country VARCHAR(255),
    PostalCode VARCHAR(255)
);

CREATE TABLE calendar (
    ID VARCHAR(255) SERIAL PRIMARY KEY,
    Date DATE,
    Availability BOOLEAN
);

CREATE TABLE users (
    ID VARCHAR(255) SERIAL PRIMARY KEY,
    Name VARCHAR(255),
    LastName VARCHAR(255),
    BirthDate DATE,
    Email VARCHAR(255) UNIQUE,
    Password VARCHAR(255),
    PhoneNumber VARCHAR(255),
    IsOwner BOOLEAN,
    CalendarId VARCHAR(255),
    AddressId VARCHAR(255),
    FOREIGN KEY (CalendarId) REFERENCES calendar(ID) ON DELETE SET NULL,
    FOREIGN KEY (AddressId) REFERENCES address(ID) ON DELETE SET NULL
);
