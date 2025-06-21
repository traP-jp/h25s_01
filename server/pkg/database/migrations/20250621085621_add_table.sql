-- +goose up
-- stations テーブル
CREATE TABLE stations (
  id VARCHAR(255) PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- shops テーブル
CREATE TABLE shops (
  id VARCHAR(255) PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  post_code VARCHAR(10),
  address TEXT NOT NULL,
  latitude VARCHAR(50),
  longitude VARCHAR(50),
  registerer VARCHAR(255),
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- shop_stations 中間テーブル
CREATE TABLE shop_stations (
  shop_id VARCHAR(255) NOT NULL REFERENCES shops(id) ON DELETE CASCADE,
  station_id VARCHAR(255) NOT NULL REFERENCES stations(id) ON DELETE CASCADE,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (shop_id, station_id)
);

-- shop_payment_methods テーブル
CREATE TABLE shop_payment_methods (
  shop_id VARCHAR(255) NOT NULL REFERENCES shops(id) ON DELETE CASCADE,
  payment_method VARCHAR(255) NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (shop_id, payment_method)
);

-- reviews テーブル
CREATE TABLE reviews (
  id VARCHAR(255) PRIMARY KEY,
  author VARCHAR(255) NOT NULL,
  shop_id VARCHAR(255) NOT NULL REFERENCES shops(id) ON DELETE CASCADE,
  rating INTEGER NOT NULL CHECK (rating >= 0 AND rating <= 3),
  content TEXT,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- shop_images 中間テーブル
CREATE TABLE shop_images (
  shop_id VARCHAR(255) NOT NULL REFERENCES shops(id) ON DELETE CASCADE,
  image_id VARCHAR(255) NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (shop_id, image_id)
);

-- review_images 中間テーブル
CREATE TABLE review_images (
  review_id VARCHAR(255) NOT NULL REFERENCES reviews(id) ON DELETE CASCADE,
  image_id VARCHAR(255) NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (review_id, image_id)
);

-- +goose down
DROP TABLE IF EXISTS review_images;
DROP TABLE IF EXISTS shop_images;
DROP TABLE IF EXISTS reviews;
DROP TABLE IF EXISTS shop_payment_methods;
DROP TABLE IF EXISTS shop_stations;
DROP TABLE IF EXISTS shops;
DROP TABLE IF EXISTS stations;