CREATE DATABASE alta_online_shop;

USE alta_online_shop;

CREATE TABLE users (
	id int(11) not null auto_increment primary key,
    nama varchar(100),
    tanggal_lahir date,
    status BOOLEAN,
    gender enum("Pria","Wanita"),
    created_at timestamp default now(),
    updated_at timestamp default now()    
);

CREATE TABLE payment_methods (
	id int(11) not null auto_increment primary key,
	nama varchar(100),
    status BOOLEAN,
    created_at timestamp default now(),
    updated_at timestamp default now()
);

CREATE TABLE transactions (
	id int(11) not null auto_increment primary key,
    user_id int(11),
    payment_id int(11),
    quantity int(11),
    total_price DECIMAL(11,2),
    created_at timestamp default now(),
    updated_at timestamp default now(),
    foreign key (user_id) references users(id),
    foreign key (payment_id) references payment_methods(id)
);

CREATE TABLE operators (
	id int(11) not null auto_increment primary key,
	nama varchar(100),
    created_at timestamp default now(),
    updated_at timestamp default now()
);

CREATE TABLE product_types (
	id int(11) not null auto_increment primary key,
    nama varchar(100),
    created_at timestamp default now(),
    updated_at timestamp default now()
);

CREATE TABLE products (
	id int(11) not null auto_increment primary key,
    product_type_id int(11),
    operator_id int(11),
    nama varchar(100),
    quantity int(11),
    price decimal(11,2),
    created_at timestamp default now(),
    updated_at timestamp default now(),
    foreign key (product_type_id) references product_types(id),
    foreign key (operator_id) references operators(id)
);

CREATE TABLE product_descriptions (
	id int(11) primary key,     
    deskripsi text,
    created_at timestamp default now(),
    updated_at timestamp default now(),
    foreign key (id) references products(id)
);

CREATE TABLE transaction_details (
	id int(11) primary key auto_increment,
    product_id int(11),
    quantity int(11),
    transaction_id int(11),
    created_at timestamp default now(),
    updated_at timestamp default now(),
    foreign key (product_id) references products(id),
    foreign key (transaction_id) references transactions(id)
);

CREATE TABLE kurir (
	id int(11) primary key auto_increment,
    nama varchar(100),
    created_at timestamp default now(),
    updated_at timestamp default now()
);

ALTER TABLE kurir ADD COLUMN ongkos_dasar int(11);

ALTER TABLE kurir RENAME to shipping;

DROP TABLE shipping;

CREATE TABLE payment_method_descriptions (
	id int(11) primary key,
    deskripsi text,
    created_at timestamp default now(),
    updated_at timestamp default now(),
    foreign key (id) references payment_methods(id)
);

CREATE TABLE alamat (
	id int(11) primary key auto_increment,
    user_id int(11),
    provinsi varchar(100),
    kota varchar(100),
    kecamatan varchar(100),
    kode_pos char(5),
    nama_jalan varchar(255),
    details varchar(255),
    created_at timestamp default now(),
    updated_at timestamp default now(),
    foreign key (user_id) references users(id)
);

CREATE TABLE user_payment_method_detail (
	id int(11) primary key auto_increment,
    user_id int(11),
    payment_method_id int(11),
    created_at timestamp default now(),
    updated_at timestamp default now(),
    foreign key (user_id) references users(id),
    foreign key (payment_method_id) references payment_methods(id)
);

Alter table products add column deleted_at timestamp;