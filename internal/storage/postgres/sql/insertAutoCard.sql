WITH new_car AS (
    INSERT INTO cars.car (
        mark,
        model,
        build_year,
        power,
        type_power,
        owners,
        condition,
        steering_wheel,
        color,
        descriptions,
        created_at,
        updated_at,
        release,
        images,
        vin,
        kilometre
    )
    VALUES (
        $3,
        $4,
        $5,
        $6,
        $7,
        $8,
        $9,
        $10,
        $11,
        $12,
        COALESCE(NULLIF($13::timestamptz, '0001-01-01 00:00:00+00'::timestamptz), now()),
        COALESCE(NULLIF($14::timestamptz, '0001-01-01 00:00:00+00'::timestamptz), now()),
        $15,
        COALESCE($16::jsonb, '[]'::jsonb),
        $17,
        $18
    )
    RETURNING id
), new_price AS (
    INSERT INTO cars.car_price (car_id, price)
    SELECT id, $19
    FROM new_car
    RETURNING car_id
), new_owner AS (
    INSERT INTO cars.user_cars (car_id, user_id)
    SELECT id, $1
    FROM new_car
    RETURNING car_id
)
SELECT id
FROM new_car;
