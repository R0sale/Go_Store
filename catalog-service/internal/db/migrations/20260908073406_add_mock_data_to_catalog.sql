-- +goose Up
-- +goose StatementBegin
INSERT INTO catalog (name, price, image_url) VALUES
('Wireless Mouse', 25.99, 'https://images.unsplash.com/photo-1527864550417-7fd91fc51a46?auto=format&fit=crop&w=500&q=60'),
('Mechanical Keyboard', 89.50, 'https://images.unsplash.com/photo-1595225476474-87563907a212?auto=format&fit=crop&w=500&q=60'),
('Large Desk Pad', 15.00, 'https://images.unsplash.com/photo-1618366712010-f4ae9c647dcb?auto=format&fit=crop&w=500&q=60'),
('USB-C Hub', 42.99, 'https://images.unsplash.com/photo-1644361566696-3d442b5b482a?auto=format&fit=crop&w=500&q=60'),
('Monitor Stand', 35.25, 'https://images.unsplash.com/photo-1527443154391-507e9dc6c5cc?auto=format&fit=crop&w=500&q=60'),
('Ergonomic Office Chair', 199.99, 'https://images.unsplash.com/photo-1505843490538-5133c6c7d0e1?auto=format&fit=crop&w=500&q=60'),
('Noise-cancelling Headphones', 150.00, 'https://images.unsplash.com/photo-1505740420928-5e560c06d30e?auto=format&fit=crop&w=500&q=60'),
('1080p Web Camera', 55.75, 'https://images.unsplash.com/photo-1621259182978-fbf93132d53d?auto=format&fit=crop&w=500&q=60');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM catalog;
-- +goose StatementEnd