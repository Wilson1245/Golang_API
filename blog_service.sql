DROP TABLE `blog_service`. `buylists`;
DROP TABLE `blog_service`.`users`;
DROP TABLE `blog_service`.`products`;

Create table blog_service.admins (
    admin_id int Auto_increment NOT NULL,
    admin_name varchar(255) NOT NULL,
    admin_password varchar(255) NOT NULL,
    admin_identity varchar(255) NOT NULL,
    PRIMARY KEY (admin_id)
);

Create table blog_service.users (
    user_id int Auto_increment NOT NULL,
    user_name varchar(255) NOT NULL,
    user_password varchar(255) NOT NULL,
    user_email varchar(255) NOT NULL,
    user_identity varchar(1) NOT NULL,
    PRIMARY KEY (user_id)
);

Create table blog_service.products (
    product_id int AUTO_INCREMENT Not null,
    product_name varchar(255) NOT NULL,
    product_qty int NOT NULL,
    product_user int NOT NULL,
    primary key (product_id)
);

Create table blog_service.buylists (
    buy_list_id int AUTO_INCREMENT NOT NULL,
    buy_user_id int NOT NULL,
    buy_product_id int NOT NULL,
    primary key (buy_list_id)
);
-- T 為 Admin , U 為 User
insert into blog_service.users (user_name, user_password, user_email, user_identity) VALUES
('test001', '123456', 'wilson052864@gmail.com', 'T'),
('test002', '123456', 'wilson0528@ymail.com','U');

insert into blog_service.products (product_name, product_qty, product_user) values
('product001', '10', '3'),('product002', '20', '4'), ('product003', '12', '5');

insert into blog_service.buylists (buy_user_id, buy_product_id) values
('1','2'), ('2', '1'), ('1', '3'), ('2', '2');

insert into blog_service.admins (admin_name, admin_password, admin_identity) values
('admin001', '123456','T'),('admin002', '123456','T');