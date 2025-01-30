CREATE TABLE court_facilities (
    court_id INT NOT NULL,
    facility_id INT NOT NULL,
    PRIMARY KEY (court_id, facility_id),
    FOREIGN KEY (court_id) REFERENCES courts(id) ON DELETE CASCADE,
    FOREIGN KEY (facility_id) REFERENCES facilities(id) ON DELETE CASCADE
);
