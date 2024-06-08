create database cargo_management;
use cargo_management;

#管理员
CREATE TABLE Admins (
                        # 管理员ID
                        admin_id INT PRIMARY KEY AUTO_INCREMENT,
                        # 管理员姓名
                        username VARCHAR(64),
                        # 管理员密码
                        password VARCHAR(64)
);

# 品牌
CREATE TABLE Brands (
                        # 品牌ID
                        brand_id INT PRIMARY KEY AUTO_INCREMENT,
                        # 品牌名称
                        brand_name VARCHAR(64)
);

# 产品类型
CREATE TABLE Product_Types (
                                # 类型ID
                              type_id INT PRIMARY KEY AUTO_INCREMENT,
                                # 类型名称
                              type_name VARCHAR(32),
                                # 父类型ID
                              parent_type_id INT,
                                # 外键关联父类型
                              FOREIGN KEY (parent_type_id) REFERENCES Product_Types(type_id)
);

# 产品
CREATE TABLE Products (
                        # 产品UPC码
                          upc_code VARCHAR(32) PRIMARY KEY,
                        # 产品名称
                          name VARCHAR(255),
                        # 产品价格
                          price DECIMAL(10, 2),
                        # 产品品牌ID
                          brand_id INT,
                        # 产品类型ID
                          type_id INT,
                          FOREIGN KEY (brand_id) REFERENCES Brands(brand_id),
                          FOREIGN KEY (type_id) REFERENCES Product_Types(type_id)
);

# 供应商
CREATE TABLE Vendors (
                        # 供应商ID
                         vendor_id INT PRIMARY KEY AUTO_INCREMENT,
                        # 供应商名称
                         vendor_name VARCHAR(255)
);

# 门店
CREATE TABLE Stores (
                        # 门店ID
                        store_id INT PRIMARY KEY AUTO_INCREMENT,
                        # 门店名称
                        store_name VARCHAR(255),
                        # 门店地址
                        location VARCHAR(255)
);

# 顾客
CREATE TABLE Customers (
                        # 顾客ID
                           customer_id INT PRIMARY KEY AUTO_INCREMENT,
                        # 顾客姓名
                           name VARCHAR(64),
                        # 顾客会员状态
                           membership_status int,
                        # 顾客联系方式
                           contact_info VARCHAR(255)
);

# 销售记录
CREATE TABLE Orders (
                    # 订单ID
                        order_id varchar(36) PRIMARY KEY,#36位UUID
                    # 门店ID
                        store_id INT,
                    # 顾客ID
                        customer_id INT,
                    # 订单日期
                        order_date datetime,
                    # 订单总金额
                        total_price DECIMAL(10, 2),
                        FOREIGN KEY (store_id) REFERENCES Stores(store_id),
                        FOREIGN KEY (customer_id) REFERENCES Customers(customer_id)
);

# 订单商品(销售记录相当于购物车, 订单商品相当于购物车中的商品)
CREATE TABLE Order_Items (
                        # 订单商品ID
                            order_item_id INT PRIMARY KEY AUTO_INCREMENT,
                        # 订单ID
                            order_id varchar(36),
                        # 商品ID
                            upc_code VARCHAR(32),
                        # 商品数量
                            quantity INT,
                        # 商品价格
                            FOREIGN KEY (order_id) REFERENCES Orders(order_id),
                            FOREIGN KEY (upc_code) REFERENCES Products(upc_code)
);

# 进货记录表
CREATE TABLE Restocking_Record (
                        # 进货记录ID
                            restocking_id INT PRIMARY KEY AUTO_INCREMENT,
                        # 产品ID
                               upc_code VARCHAR(32),
                        # 供应商ID
                               vendor_id INT,
                        # 进货量
                                purchase_amount INT,
                        # 总价
                                total_price DECIMAL(10, 2),
                        # 进货日期
                                restocking_date DATE,
                        #门店
                                store_id INT,

                               FOREIGN KEY (upc_code) REFERENCES Products(upc_code),
                               FOREIGN KEY (vendor_id) REFERENCES Vendors(vendor_id),
                            FOREIGN KEY (store_id) REFERENCES Stores(store_id)

);

# 供应商产品
CREATE TABLE Vendor_Products (
                            # 产品UPC码
                                 upc_code VARCHAR(32),
                            # 供应商ID
                                 vendor_id INT,
                            # 进货价格
                                 purchase_price DECIMAL(10, 2),
                                 PRIMARY KEY (upc_code, vendor_id),
                                 FOREIGN KEY (upc_code) REFERENCES Products(upc_code),
                                 FOREIGN KEY (vendor_id) REFERENCES Vendors(vendor_id)
);

# 产品门店
CREATE TABLE Product_Store (
                            # 产品ID
                              upc_code VARCHAR(32),
                            # 门店ID
                              store_id INT,
                            # 库存数量
                              inventory_amount INT,
                            # 产品门店关系表的主键
                              PRIMARY KEY (upc_code, store_id),
                              FOREIGN KEY (upc_code) REFERENCES Products(upc_code),
                              FOREIGN KEY (store_id) REFERENCES Stores(store_id)
);

# 进货计划表
CREATE TABLE Restocking_Tasks (
                                  id INT AUTO_INCREMENT PRIMARY KEY,
                                  upc_code VARCHAR(255),
                                  store_id INT,
                                  restock_amount INT,
                                  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

# ——————————————————————
-- 插入品牌数据
INSERT INTO Brands (brand_id, brand_name) VALUES
                                            (1, '品牌一'),
                                            (2, '品牌二'),
                                            (3, '品牌三');

-- 插入产品类型数据
INSERT INTO Product_Types (type_id, type_name, parent_type_id) VALUES
                                                              (1, '类型一', NULL),
                                                              (2, '类型二', 1),
                                                              (3, '类型三', 1);

-- 插入产品数据
INSERT INTO Products (upc_code, name, price, brand_id, type_id) VALUES
                                                                ('0001', '产品一', 1.0, 1, 1),
                                                                ('0002', '产品二', 2.0, 2, 2),
                                                                ('0003', '产品三', 3.0, 3, 3);

-- 插入供应商数据
INSERT INTO Vendors (vendor_id, vendor_name) VALUES
                                               (1, '供应商一'),
                                               (2, '供应商二'),
                                               (3, '供应商三');

-- 插入门店数据
INSERT INTO Stores (store_id, store_name, location) VALUES
                                                      (1, '门店一', '北京'),
                                                      (2, '门店二', '上海'),
                                                      (3, '门店三', '广州');

-- 插入顾客数据
INSERT INTO Customers (customer_id, name, membership_status, contact_info) VALUES
                                                                            (1, '顾客一', 1, '13800138000'),
                                                                            (2, '顾客二', 0, '13900139000'),
                                                                            (3, '顾客三', 1, '13700137000');

-- 插入产品门店数据
INSERT INTO Product_Store (upc_code, store_id, inventory_amount) VALUES
                                                                 ('0001', 1, 100),
                                                                 ('0002', 2, 200),
                                                                 ('0003', 3, 300);

-- 插入更多的品牌数据
INSERT INTO Brands (brand_id, brand_name) VALUES
                                            (4, '品牌四'),
                                            (5, '品牌五'),
                                            (6, '品牌六');

-- 插入更多的产品类型数据
INSERT INTO Product_Types (type_id, type_name, parent_type_id) VALUES
                                                              (4, '类型四', NULL),
                                                              (5, '类型五', 4),
                                                              (6, '类型六', 4);

-- 插入更多的产品数据
INSERT INTO Products (upc_code, name, price, brand_id, type_id) VALUES
                                                                ('0004', '产品四', 4.0, 4, 4),
                                                                ('0005', '产品五', 5.0, 5, 5),
                                                                ('0006', '产品六', 6.0, 6, 6);

-- 插入更多的供应商数据
INSERT INTO Vendors (vendor_id, vendor_name) VALUES
                                               (4, '供应商四'),
                                               (5, '供应商五'),
                                               (6, '供应商六');

-- 插入更多的门店数据
INSERT INTO Stores (store_id, store_name, location) VALUES
                                                      (4, '门店四', '深圳'),
                                                      (5, '门店五', '杭州'),
                                                      (6, '门店六', '成都');

-- 插入更多的顾客数据
INSERT INTO Customers (customer_id, name, membership_status, contact_info) VALUES
                                                                            (4, '顾客四', 0, '13600136000'),
                                                                            (5, '顾客五', 1, '13500135000'),
                                                                            (6, '顾客六', 0, '13400134000');

-- 插入更多的产品门店数据
INSERT INTO Product_Store (upc_code, store_id, inventory_amount) VALUES
                                                                 ('0004', 4, 400),
                                                                 ('0005', 5, 500),
                                                                 ('0006', 6, 600);

INSERT INTO Vendor_Products (upc_code, vendor_id, purchase_price) VALUES
                                                                      ('0001', 1, 10.0),
                                                                      ('0001', 2, 1.0),
                                                                      ('0001', 3, 1.0),
                                                                      ('0001', 4, 1.0),
                                                                      ('0001', 5, 1.0),
                                                                      ('0001', 6, 12.0),
                                                                      ('0002', 1, 2.0),
                                                                      ('0002', 2, 2.0),
                                                                      ('0002', 3, 2.0),
                                                                      ('0002', 4, 2.0),
                                                                      ('0002', 5, 2.0),
                                                                      ('0002', 6, 2.0),
                                                                      ('0003', 1, 3.0),
                                                                      ('0003', 2, 3.0),
                                                                      ('0003', 3, 3.0),
                                                                      ('0003', 4, 3.0),
                                                                      ('0003', 5, 3.0),
                                                                      ('0003', 6, 3.0),
                                                                      ('0004', 1, 15.0),
                                                                      ('0004', 2, 1.0),
                                                                      ('0004', 3, 2.0),
                                                                      ('0004', 4, 3.0),
                                                                      ('0004', 5, 4.0),
                                                                      ('0004', 6, 5.0),
                                                                      ('0005', 1, 50.0),
                                                                      ('0005', 2, 51.0),
                                                                      ('0005', 3, 5.0),
                                                                      ('0005', 4, 5.0),
                                                                      ('0005', 5, 5.0),
                                                                      ('0005', 6, 5.0),
                                                                      ('0006', 1, 10.0),
                                                                      ('0006', 2, 1.0),
                                                                      ('0006', 3, 1.0),
                                                                      ('0006', 4, 1.0),
                                                                      ('0006', 5, 1.0),
                                                                      ('0006', 6, 5.0);

# —————————————————————— 数据库用户管理
CREATE TABLE users (
                       username VARCHAR(64) PRIMARY KEY,
                       password VARCHAR(64)
);

insert into users (username, password) values ('test', 'test');
create user 'test'@'%' identified by 'test';

# —————————————————————— 数据库自动程式
# 创建进货记录时，更新产品门店库存
DELIMITER //
CREATE TRIGGER update_inventory
    AFTER INSERT ON Restocking_Record
    FOR EACH ROW
BEGIN
    DECLARE original_inventory INT;

    -- 先查询当前库存量
    SELECT inventory_amount INTO original_inventory
    FROM Product_Store
    WHERE upc_code = NEW.upc_code AND store_id = NEW.store_id;

    -- 如果没有找到记录，则进行插入操作
    IF original_inventory IS NULL THEN
        INSERT INTO Product_Store (upc_code, store_id, inventory_amount)
        VALUES (NEW.upc_code, NEW.store_id, NEW.purchase_amount);
    ELSE
        -- 如果找到了记录，则进行更新操作
        UPDATE Product_Store
        SET inventory_amount = original_inventory + NEW.purchase_amount
        WHERE upc_code = NEW.upc_code AND store_id = NEW.store_id;
    END IF;
END;
//
DELIMITER ;

# 测试触发器
INSERT INTO Restocking_Record (upc_code, vendor_id, purchase_amount, total_price, restocking_date, store_id) VALUES
                                                                                                        ('1145', 7, 10, 10.0, '2021-01-01', 2);

# 创建销售记录时，更新产品门店库存
DELIMITER //
CREATE TRIGGER decrease_inventory
    AFTER INSERT ON Order_Items
    FOR EACH ROW
BEGIN
    DECLARE original_inventory INT;
    DECLARE store_id_for_order INT;

    SELECT store_id INTO store_id_for_order
    FROM Orders
    WHERE Orders.order_id = NEW.order_id;

    SELECT inventory_amount INTO original_inventory
    FROM Product_Store
    WHERE Product_Store.upc_code = NEW.upc_code AND Product_Store.store_id = store_id_for_order
        FOR UPDATE;

    UPDATE Product_Store
    SET inventory_amount = original_inventory - NEW.quantity
    WHERE upc_code = NEW.upc_code AND store_id = store_id_for_order;
END;
//
DELIMITER ;

DELIMITER //
CREATE EVENT restock_products
    ON SCHEDULE EVERY 1 DAY
        STARTS CURRENT_TIMESTAMP
    DO
    BEGIN
        INSERT INTO Restocking_Record (upc_code, vendor_id, purchase_amount, total_price, restocking_date, store_id)
        SELECT upc_code,
               (SELECT vendor_id FROM Vendor_Products WHERE upc_code = Restocking_Tasks.upc_code ORDER BY purchase_price LIMIT 1),
               restock_amount,
               restock_amount * (SELECT purchase_price FROM Vendor_Products WHERE upc_code = Restocking_Tasks.upc_code ORDER BY purchase_price LIMIT 1),
               NOW(),
               store_id
        FROM Restocking_Tasks;

        DELETE FROM Restocking_Tasks;
    END;
//
DELIMITER ;

show triggers;