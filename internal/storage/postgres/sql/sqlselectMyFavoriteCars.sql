SELECT
    c.id,
    c.mark,
    c.model,
    COALESCE(c.kilometre, 0),
    COALESCE(cp.price, 0),
    c.images
FROM cars.favorite_cars AS fc
JOIN cars.car AS c ON c.id = fc.car_id
LEFT JOIN cars.car_price AS cp ON cp.car_id = c.id
WHERE fc.user_id = $1
  AND c.release = TRUE
ORDER BY fc.created_at DESC, c.id;
