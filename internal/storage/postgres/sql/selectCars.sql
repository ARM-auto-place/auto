SELECT
    c.id,
    c.mark,
    c.model,
    COALESCE(c.kilometre, 0),
    COALESCE(cp.price, 0),
    c.images
FROM cars.car AS c
LEFT JOIN cars.car_price AS cp ON cp.car_id = c.id
WHERE c.release = TRUE
ORDER BY c.created_at DESC, c.id;
