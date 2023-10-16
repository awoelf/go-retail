-- Add up migration script here
INSERT INTO Departments(Name)
	VALUES
		("Meat"),	 -- ID = 1
		("Produce"), -- ID = 2
		("Seafood"); -- ID = 3


INSERT INTO Managers(FirstName, LastName, DepartmentID)
	VALUES
		("Alexis", "Red", 1),
		("John", "Blue", 2);

INSERT INTO Aisles(DepartmentID)
	VALUES
		(1), -- ID = 1
		(1), -- ID = 2
		(2), -- ID = 3
		(2); -- ID = 4

INSERT INTO Items(Name, Price, Qty, Category, AisleID)
	VALUES
		("apple", 1.99, 500, "fruit", 2);