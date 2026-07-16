USE Game_Store;

-- CATEGORIES

INSERT INTO categories
(name, description)
VALUES
('Shooter', 'First-person & third-person shooting games'),
('Action', 'Action-based games'),
('Racing', 'Car & motorcycle racing games'),
('Adventure', 'Exploration & story-driven games'),
('Sport', 'Sports simulation games');

-- PRODUCTS

INSERT INTO products
(category_id, title, price, stock)
VALUES
(1, 'Call Of Duty', 700000, 30),
(1, 'PUBG', 550000, 23),

(2, 'GTA V', 1000000, 12),
(2, 'Red Dead Redemption 2', 925000, 34),

(3, 'Mario Kart', 400000, 52),
(3, 'Crash Team Racing', 350000, 18),

(4, 'The Legend of Zelda', 500000, 9),
(4, 'Assassin Creed', 675000, 29),

(5, 'FIFA 25', 800000, 41),
(5, 'PES 21', 735000, 54);

-- USERS

INSERT INTO users
(email, password, name, role)
VALUES
('admin@mail.com', 'Admin123', 'Administrator', 'admin'),
('hacktiv8@mail.com', 'Hacktiv8', 'Hacktiv', 'customer');

-- USER PROFILES

INSERT INTO user_profiles
(user_id, user_name, phone)
VALUES
(1, 'AdminSejati', '081818181818'),
(2, 'HacktivDelapan', '088888888888');

-- ORDERS

INSERT INTO orders
(user_id, status, total_price)
VALUES
(2, 'paid', 1400000),
(2, 'pending', 800000),
(2, 'cancelled', 550000);

-- ORDER ITEMS

INSERT INTO order_items
(order_id, product_id, qty, price)
VALUES

-- Order #1
(1, 1, 2, 1400000),

-- Order #2
(2, 9, 1, 800000),

-- Order #3
(3, 2, 1, 550000);

-- REVIEWS

INSERT INTO reviews
(user_id, product_id, rating, comment)
VALUES
(2, 1, 5, 'Amazing gameplay and graphics!'),
(2, 2, 4, 'Great battle royale experience.'),
(2, 3, 5, 'One of the best open-world games.'),
(2, 7, 5, 'Amazing story and exploration.'),
(2, 9, 4, 'Fun to play with friends.');