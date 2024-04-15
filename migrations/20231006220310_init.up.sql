CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS departments(
    "id" uuid PRIMARY KEY NOT NULL DEFAULT (uuid_generate_v4()),
    "name" VARCHAR NOT NULL,
    "totalSalesDept" FLOAT DEFAULT(0) NOT NULL,
	"createdAt" TIMESTAMP DEFAULT NOW(),
	"updatedAt" TIMESTAMP DEFAULT NOW(),
);

CREATE TABLE IF NOT EXISTS Managers(
    ID INT NOT NULL UNIQUE,
    FirstName VARCHAR(127) NOT NULL,
    LastName VARCHAR(127) NOT NULL,
    DepartmentID INT,
	CreatedAt TIMESTAMP DEFAULT(NOW()),
	UpdatedAt TIMESTAMP DEFAULT(NOW()),

    
    INDEX(ID),
    INDEX(DepartmentID),
    
    FOREIGN KEY(DepartmentID)
    REFERENCES Departments(ID)
    ON UPDATE CASCADE,
    
    PRIMARY KEY(ID)
);

CREATE TABLE IF NOT EXISTS Items(
    ID INT NOT NULL UNIQUE,
    Name VARCHAR(127) NOT NULL UNIQUE,
    Price FLOAT NOT NULL,
    Qty INT NOT NULL,
	QtySold INT NOT NULL DEFAULT (0),
	Category VARCHAR(127) NOT NULL DEFAULT("NONE"),
    Promotion BOOLEAN DEFAULT(FALSE) NOT NULL,
    Replenish BOOLEAN DEFAULT(FALSE) NOT NULL,
    TotalSalesItem FLOAT DEFAULT(0) NOT NULL,
	Aisle INT NOT NULL,
	DepartmentID INT NOT NULL,
	CreatedAt TIMESTAMP DEFAULT(NOW()),
	UpdatedAt TIMESTAMP DEFAULT(NOW()),

    INDEX(ID),
    INDEX(DepartmentID),
    
    FOREIGN KEY(DepartmentID)
    REFERENCES Departments(ID)
    ON UPDATE CASCADE,
    
    PRIMARY KEY(ID)
);