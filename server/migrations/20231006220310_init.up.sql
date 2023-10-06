-- Add up migration script here
CREATE TABLE IF NOT EXISTS Departments(
    ID VARCHAR(36) default (UUID()) UNIQUE,
    Name VARCHAR (127) NOT NULL UNIQUE,
    TotalSalesWeekDept FLOAT DEFAULT (0) NOT NULL,
    
	INDEX(ID),
    
    PRIMARY KEY (ID)
);

CREATE TABLE IF NOT EXISTS Managers(
    ID VARCHAR(36) default (UUID()) UNIQUE,
    FirstName VARCHAR (127) NOT NULL,
    LastName VARCHAR (127) NOT NULL,
    DepartmentID VARCHAR(36),
    
    INDEX (ID),
    INDEX (DepartmentID),
    
    FOREIGN KEY (DepartmentID)
    REFERENCES Departments(ID)
    ON UPDATE CASCADE ON DELETE RESTRICT,
    
    PRIMARY KEY (ID)
);

CREATE TABLE IF NOT EXISTS Aisles(
    ID VARCHAR(36) default (UUID()) UNIQUE,
	DepartmentID VARCHAR(36) NOT NULL,
    
    INDEX (ID),
    INDEX (DepartmentID),
    
    FOREIGN KEY (DepartmentID)
    REFERENCES Departments(ID)
    ON UPDATE CASCADE ON DELETE RESTRICT,
    
    PRIMARY KEY (ID)
);

CREATE TABLE IF NOT EXISTS Items(
    ID VARCHAR(36) DEFAULT (UUID()) UNIQUE,
    Name VARCHAR (127) NOT NULL UNIQUE,
    Price FLOAT NOT NULL,
    Qty INT NOT NULL,
    Promotion BOOLEAN DEFAULT (FALSE) NOT NULL,
    PromotionPrice FLOAT DEFAULT (0) NOT NULL,
    Replenish BOOLEAN DEFAULT (FALSE) NOT NULL,
    TotalSalesWeekItem FLOAT DEFAULT (0) NOT NULL,
    AisleID VARCHAR(36) NOT NULL,
    DepartmentID VARCHAR(36) NOT NULL,
    
    INDEX (ID),
    INDEX (AisleID, DepartmentID),
    
    FOREIGN KEY (AisleID)
    REFERENCES Aisles(ID)
    ON UPDATE CASCADE ON DELETE RESTRICT,
    
    FOREIGN KEY (DepartmentID)
    REFERENCES Departments(ID)
    ON UPDATE CASCADE ON DELETE RESTRICT,
    
    PRIMARY KEY (ID)
);