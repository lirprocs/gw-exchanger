CREATE TABLE exchange_rates (
                                from_currency VARCHAR(3) NOT NULL,
                                to_currency VARCHAR(3) NOT NULL,
                                rate FLOAT NOT NULL,
                                PRIMARY KEY (from_currency, to_currency)
);
INSERT INTO exchange_rates (from_currency, to_currency, rate) VALUES
                                                                  ('USD', 'EUR', 0.94),
                                                                  ('USD', 'GBP', 0.82),
                                                                  ('USD', 'RUB', 85.5),
                                                                  ('USD', 'JPY', 150.0),
                                                                  ('USD', 'AUD', 1.45),
                                                                  ('USD', 'CAD', 1.34),
                                                                  ('USD', 'CHF', 0.91),
                                                                  ('USD', 'CNY', 7.0),
                                                                  ('USD', 'INR', 82.5),
                                                                  ('EUR', 'USD', 1.06),
                                                                  ('EUR', 'GBP', 0.87),
                                                                  ('EUR', 'RUB', 90.8),
                                                                  ('EUR', 'JPY', 160.0),
                                                                  ('EUR', 'AUD', 1.55),
                                                                  ('EUR', 'CAD', 1.43),
                                                                  ('EUR', 'CHF', 0.97),
                                                                  ('EUR', 'CNY', 7.45),
                                                                  ('EUR', 'INR', 87.5),
                                                                  ('GBP', 'USD', 1.22),
                                                                  ('GBP', 'EUR', 1.15),
                                                                  ('GBP', 'RUB', 104.5),
                                                                  ('GBP', 'JPY', 185.0),
                                                                  ('GBP', 'AUD', 1.67),
                                                                  ('GBP', 'CAD', 1.63),
                                                                  ('GBP', 'CHF', 1.10),
                                                                  ('GBP', 'CNY', 8.6),
                                                                  ('GBP', 'INR', 106.5),
                                                                  ('RUB', 'USD', 0.0117),
                                                                  ('RUB', 'EUR', 0.011),
                                                                  ('RUB', 'GBP', 0.0096),
                                                                  ('RUB', 'JPY', 1.74),
                                                                  ('RUB', 'AUD', 0.018),
                                                                  ('RUB', 'CAD', 0.016),
                                                                  ('RUB', 'CHF', 0.012),
                                                                  ('RUB', 'CNY', 0.087),
                                                                  ('RUB', 'INR', 1.0),
                                                                  ('JPY', 'USD', 0.0067),
                                                                  ('JPY', 'EUR', 0.0063),
                                                                  ('JPY', 'GBP', 0.0054),
                                                                  ('JPY', 'RUB', 0.57),
                                                                  ('JPY', 'AUD', 0.0095),
                                                                  ('JPY', 'CAD', 0.0089),
                                                                  ('JPY', 'CHF', 0.0061),
                                                                  ('JPY', 'CNY', 0.051),
                                                                  ('JPY', 'INR', 7.2);
