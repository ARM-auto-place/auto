SELECT
    c.id,
    COALESCE(uc.user_id, '00000000-0000-0000-0000-000000000000'::uuid),
    c.mark,
    c.model,
    c.build_year,
    c.power,
    c.type_power,
    c.owners,
    c.condition,
    c.steering_wheel,
    c.color,
    c.descriptions,
    c.created_at,
    c.updated_at,
    c.release,
    c.images,
    c.vin,
    c.kilometre,
    cp.price,
    uc.user_id
    EXISTS (
        SELECT 1
        FROM cars.favorite_cars AS fc
        WHERE fc.car_id = c.id
    )
FROM cars.car AS c
LEFT JOIN cars.user_cars AS uc ON uc.car_id = c.id
LEFT JOIN cars.car_price AS cp ON cp.car_id = c.id
WHERE c.id = $1;
