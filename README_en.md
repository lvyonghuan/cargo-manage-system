# Cargo Management System

![cargo_management.png](assets/cargo_management.png)

This system implements the functionality of three layers: front-end, back-end, and the underlying database. Through the front-end, administrators can conveniently manage goods, such as stocking operations. The back-end is developed using Go, supporting concurrent database operations. The underlying database uses MySQL.

Usage:

    Fill in the config.json with the MySQL account and password.
    Click cargo-manage-system.exe to run.
    Open the browser and go to 127.0.0.1:8080 to access the login page.
    Enter your username and password to enter the management system.

## Detailed Explanation
### E-R Diagram

![ER.png](assets/E-R.png)

### Implementation

The back-end interacts with the underlying database. Specific behaviors are as follows:
Purchase

Since the purchase behavior is not related to the administrator, no administrator interface is designed. The purchase behavior is implemented through direct API access. The purchase simulates the operation at the cashier, and each purchase is a whole transaction that ends and commits upon checkout.

    Start a purchase by accessing the /customer/purchase/begin route, which generates a UUID as the primary key of the purchase record.
    Add products by accessing the /customer/purchase/add route and passing in the UUID. The purchase record is written into the transaction and associated with the foreign key. Each time a product is added, a trigger is activated to automatically reduce the inventory.
    End the purchase by accessing the /customer/purchase/end route and passing in the UUID. This ends the purchase and commits the transaction.

Stocking and Automatic Restocking

The back-end management system provides an automatic restocking feature when the goods are insufficient. Automatic restocking is achieved through an event and a series of triggers. The logic of the trigger is as follows:

```sql

DECLARE existing_task INT;
SELECT COUNT(*) INTO existing_task FROM Restocking_Tasks
WHERE upc_code = NEW.upc_code AND store_id = NEW.store_id;
IF NEW.store_id = [Auto-Restock Store ID]
AND NEW.upc_code = [Auto-Restock Product UPC Code]
AND NEW.inventory_amount < [Auto-Restock Inventory Threshold] AND existing_task = 0 THEN
INSERT INTO Restocking_Tasks (upc_code, store_id, restock_amount)
VALUES (NEW.upc_code, NEW.store_id, [Restock Quantity]);
END IF;

This trigger inserts an order with a specified restock quantity into the restocking plan table when the inventory of a specified product is below a certain threshold. At midnight each day, an event scans the restocking plan table and transfers the restocking plans to the restocking table.

The SQL code for the event is as follows:

```sql

CREATE DEFINER = root@localhost EVENT restock_products ON SCHEDULE
EVERY '1' DAY
STARTS '2024-06-06 15:36:27'
ENABLE
DO
BEGIN
INSERT INTO Restocking_Record (upc_code, vendor_id,
purchase_amount, total_price, restocking_date, store_id)
SELECT upc_code,
(SELECT vendor_id FROM Vendor_Products
WHERE upc_code = Restocking_Tasks.upc_code ORDER BY purchase_price LIMIT 1),
restock_amount,
restock_amount * (SELECT purchase_price FROM
Vendor_Products WHERE upc_code = Restocking_Tasks.upc_code ORDER BY purchase_price LIMIT 1),
NOW(),
store_id
FROM Restocking_Tasks;

        DELETE FROM Restocking_Tasks;
    END;
```

Also, when there is an insert operation in the restocking plan table, a trigger is activated to automatically increase the inventory. The code for this trigger is as follows:

```sql

CREATE TRIGGER update_inventory
AFTER INSERT ON Restocking_Record
FOR EACH ROW
BEGIN
DECLARE original_inventory INT;

    -- Query the current inventory amount first
    SELECT inventory_amount INTO original_inventory
    FROM Product_Store
    WHERE upc_code = NEW.upc_code AND store_id = NEW.store_id;

    -- If no record is found, perform the insert operation
    IF original_inventory IS NULL THEN
        INSERT INTO Product_Store (upc_code, store_id, inventory_amount)
        VALUES (NEW.upc_code, NEW.store_id, NEW.purchase_amount);
    ELSE
        -- If a record is found, perform the update operation
        UPDATE Product_Store
        SET inventory_amount = original_inventory + NEW.purchase_amount
        WHERE upc_code = NEW.upc_code AND store_id = NEW.store_id;
    END IF;
END;
```


#### Some Query Codes:

Query the top three best-selling products:

```sql

SELECT
p.upc_code,
p.name,
SUM(oi.quantity) AS total_quantity_sold
FROM
Order_Items oi
JOIN
Products p ON oi.upc_code = p.upc_code
GROUP BY
p.upc_code,
p.name
ORDER BY
total_quantity_sold DESC
LIMIT 3;
```
Query the top three stores with the highest sales:

```sql

SELECT
s.store_id,
s.store_name,
SUM(o.total_price) AS total_sales
FROM
Orders o
JOIN
Stores s ON o.store_id = s.store_id
GROUP BY
s.store_id,
s.store_name
ORDER BY
total_sales DESC
LIMIT 3;
```
Query the top three customers with the highest purchase volume:

```sql

SELECT
c.customer_id,
c.name,
SUM(oi.quantity) AS total_quantity_purchased
FROM
Order_Items oi
JOIN
Orders o ON oi.order_id = o.order_id
JOIN
Customers c ON o.customer_id = c.customer_id
GROUP BY
c.customer_id,
c.name
ORDER BY
total_quantity_purchased DESC
LIMIT 3;
```