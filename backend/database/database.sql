CREATE TABLE IF NOT EXISTS users (user_id SERIAL PRIMARY KEY, user_name TEXT NOT NULL, balance INT DEFAULT 0, membership INT DEFAULT 0);

CREATE TABLE IF NOT EXISTS banners (banner_id SERIAL PRIMARY KEY, banner_name TEXT NOT NULL, banner_image TEXT NOT NULL, next_url TEXT, date_start DATE NOT NULL, date_end DATE NOT NULL );

CREATE TABLE IF NOT EXISTS user_x_banner (user_id INT, banner_id INT, primary key (user_id, banner_id));

INSERT INTO users (user_name, balance, membership) VALUES('Tono', 0, 0);

INSERT INTO users (user_name, balance, membership) VALUES('Uddin', 50000, 0);

INSERT INTO users (user_name, balance, membership) VALUES('Blangkon', 100000, 0);

INSERT INTO users (user_name, balance, membership) VALUES('John Doe', 1000000, 1);

INSERT INTO users (user_name, balance, membership) VALUES('Jane Doe', 50, 2);

INSERT INTO users (user_name, balance, membership) VALUES('Andy Smith', 1000, 0);

INSERT INTO banners (banner_name, banner_image, next_url, date_start, date_end) VALUES('Diskon Makanan', 'diskon_makanan.png', 'next.com', '20211002', '20211130');

INSERT INTO banners (banner_name, banner_image, next_url, date_start, date_end) VALUES('Diskon Gadget', 'diskon_gadget.png', 'next.com', '20211002', '20211130');

INSERT INTO banners (banner_name, banner_image, next_url, date_start, date_end) VALUES('Diskon Baju', 'diskon_baju.png', 'next.com', '20211002', '20211130');

INSERT INTO banners (banner_name, banner_image, next_url, date_start, date_end) VALUES('Upgrade Membership', 'upgrade_membership.png', 'https://tokopedia.com', '20211003', '20211230');

INSERT INTO banners (banner_name, banner_image, next_url, date_start, date_end) VALUES('Top Up Saldo Tokopedia', 'topup_saldo.png', 'https://tokopedia.com', '20211003', '20211103');

INSERT INTO banners (banner_name, banner_image, next_url, date_start, date_end) VALUES('Kupon WIB', 'kupon_wib.png', 'https://tokopedia.com', '20211003', '20211103');