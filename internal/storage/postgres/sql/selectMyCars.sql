SELECT
    c.id,
    c.mark,
    c.model,
    COALESCE(c.kilometre, 0),
    COALESCE(cp.price, 0),
    c.images
FROM cars.user_cars AS uc
JOIN cars.car AS c ON c.id = uc.car_id
LEFT JOIN cars.car_price AS cp ON cp.car_id = c.id
WHERE uc.user_id = $1
ORDER BY uc.created_at DESC, c.id;
