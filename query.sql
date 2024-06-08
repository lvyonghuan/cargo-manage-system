use cargo_management;

#查询

#查询销售量最高的前三项商品
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

#查询销售额最高的前三家商店
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

#查询购买商品数量最多的前三位顾客
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