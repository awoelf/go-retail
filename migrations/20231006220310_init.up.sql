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

INSERT INTO Items(Name, Price, Qty, Category, DepartmentIDID)
	VALUES
		("apple", 1.99, 500, "fruit", 2);