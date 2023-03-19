use alta_online_shop;

-- Insert 5 operators pada table operators.
insert into operators (nama) VALUES ("Untckd"),("Maternal Disaster"),("Wadezing"),("Thanksinsomnia"),("Erigo");

-- Insert 3 product type.
insert into product_types (nama) VALUES ("Sweatshirt"), ("Pants"), ("T-shirt");

insert into products (product_type_id,operator_id,nama,price,quantity) VALUES
	-- Insert 2 product dengan product type id = 1, dan operators id = 3. 
    (1,3,"Black Hoodie",200000,100),(1,3,"Broken White Hoodie",200000,75), 
    -- Insert 3 product dengan product type id = 2, dan operators id = 1. 
    (2,1,"Black Ankle Pants",250000,30), (2,1,"Blue Jeans",300000,25), (2,1,"Black Jeans",300000,15),
    -- Insert 3 product dengan product type id = 3, dan operators id = 4.
    (3,4,"T-shirt Luffy",95000,30), (3,4,"T-shirt Kurohige",95000,25), (3,4,"T-shirt Shirohige",300000,35);

-- Insert product description pada setiap product.
insert into product_descriptions (id,deskripsi) values (1,"Hoodie berwana hitam kain berbahan katun flece"), (2,"Hoodie berwana broken white dengan kain berbahan katun fleece"),
	(3, "Ankle pants berwarna hitam berbahan chinos"), (4,"Jeans berwarna biru"), (5,"Jeans berwarna hitam"), (6, "T-shirt bergambarkan luffy dari serial onepiece"),
    (7, "T-shirt bergambarkan kurohige dari serial One Piece"), (8, "T-shirt bergambarkan shirohige dari serial one piece");

-- Insert 3 payment methods.     
insert into payment_methods (nama) values ("DANA"), ("BNI"), ("GOPAY");

-- Insert 5 user pada tabel user. 
insert into users (nama, tanggal_lahir, status, gender) values ("Nurhadi","2000-07-19",0,1),("Imam Mahdi","1999-09-29",1,1),
	("Hidayat","1982-06-12",1,1),("Rere Sukawati","1995-01-02",0,2),("Dina Dwi","1992-12-01",1,2);

-- Insert 3 transaksi di masing-masing user. (soal berlanjut ke soal 1.j) 
insert into transactions (user_id, payment_id, quantity, total_price) values (1,1,1,200000),(2,2,2,190000),(3,3,3,285000),(4,1,1,300000),(5,2,1,300000);

-- Insert 3 product di masing-masing transaksi. 
insert into transaction_details (transaction_id, product_id,quantity) values (1,1,1),(2,7,1),(3,6,1),(4,4,1),(7,5,1);

-- Tampilkan nama user / pelanggan dengan gender Laki-laki / M. 
select nama from users where gender = "Pria";

-- Tampilkan product dengan id = 3. 
select * from products;

-- Tampilkan data pelanggan yang created_at dalam range 7 hari kebelakang dan mempunyai nama mengandung kata ‘a’. 
select * from users where nama like "%a%" and created_at >= date_sub(now(), interval 7 day);

-- Hitung jumlah user / pelanggan dengan status gender Perempuan. 
select count(id) as jumlah from users where gender = "Wanita";

-- Tampilkan data pelanggan dengan urutan sesuai nama abjad 
select * from users order by nama asc;

-- Tampilkan 5 data pada data product 
select * from products limit 5;

-- Ubah data product id 1 dengan nama ‘product dummy’. 
update products set nama = "Product Dummy" where id = 1;

-- Update qty = 3 pada transaction detail dengan product id = 1. 
update transaction_details set quantity = 3 where product_id = 1;

-- Delete data pada tabel product dengan id = 1. 
delete from products where id = 1;

-- Delete pada pada tabel product dengan product type id 1.
delete from products where product_type_id = 1;

update products set deleted_at = now() where id =1;
update products set deleted_at = now() where product_type_id = 1;
select * from products;
 
-- Gabungkan data transaksi dari user id 1 dan user id 2.
select * from transactions where user_id in (1,2);

select * from transactions;


select user_id,  sum(total_price) as total_price from transactions group by user_id;

-- Tampilkan jumlah harga transaksi user id 1. 
select sum(total_price) as total_price from transactions where user_id = 1;

select * from products where product_type_id = 2;
select * from transaction_details;

-- Tampilkan jumlah harga transaksi user id 1. 
select count(*)  from transactions t 
join transaction_details td on t.id = td.transaction_id
join products p on p.id = td.product_id
where p.product_type_id = 2;

-- Tampilkan semua field table product dan field name table product type yang saling berhubungan. 
select p.*, pt.nama as type_product from products p
join product_types pt on pt.id = p.product_type_id;

-- Tampilkan semua field table transaction, field name table product dan field name table user. 
select t.*, p.nama as nama_product, u.nama as user from transaction_details td
join transactions t on t.id = td.transaction_id
join users u on u.id = t.user_id
join products p on p.id = td.product_id
where p.deleted_at is null

-- Buat function setelah data transaksi dihapus maka transaction detail terhapus juga dengan transaction id yang dimaksud. 
delimiter $$
create trigger delete_transaction_detail
before delete on transactions for each row
begin
declare trans_id int;
set trans_id = old.id;
delete from transaction_details where transaction_id = trans_id;
end $$

-- Buat function setelah data transaksi detail dihapus maka data total_qty terupdate berdasarkan qty data transaction id yang dihapus. 
delimiter $$
create trigger update_qty
before delete on transaction_details for each row
begin
declare t_id int;
declare qty int;
set qty = old.quantity;
set t_id = old.transaction_id;
update transactions t set quantity = t.quantity - qty;
end $$

delete from transaction_details where transaction_id = 1;

select * from transaction_details;
select * from transactions;

-- Tampilkan data products yang tidak pernah ada di tabel transaction_details dengan sub-query. 
select * from products where id not in
(	select product_id from transaction_details 
)
