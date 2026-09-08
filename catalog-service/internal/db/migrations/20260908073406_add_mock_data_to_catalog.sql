-- +goose Up
-- +goose StatementBegin
INSERT INTO catalog (name, price, image_url) VALUES
('Wireless Mouse', 25.99, 'wireless_mouse.png'),
('Mechanical Keyboard', 89.50, 'mech_keyboard.jpg'),
('Large Desk Pad', 15.00, 'desk_pad.png'),
('USB-C Hub', 42.99, 'usbc_hub.jpg'),
('Monitor Stand', 35.25, 'monitor_stand.png'),
('Ergonomic Office Chair', 199.99, 'office_chair.jpg'),
('Noise-cancelling Headphones', 150.00, 'headphones.png'),
('1080p Web Camera', 55.75, 'webcam.jpg');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM catalog;
-- +goose StatementEnd