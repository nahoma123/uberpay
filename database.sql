CREATE TABLE "EmailNotification"(
                                    "id" UUID DEFAULT uuid_generate_v4 () NOT NULL,
                                    "body" VARCHAR(255) NOT NULL,
                                    "from" VARCHAR(255) NOT NULL,
                                    "to"   VARCHAR(255) NOT NULL,
                                    "subject" VARCHAR(255) NOT NULL,
                                    "status" VARCHAR(255) NOT NULL,
                                    "created_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL,
                                    "updated_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL
);

CREATE TABLE "PushedNotification"(
                                     "id" UUID DEFAULT uuid_generate_v4 () NOT NULL,
                                     "api_key" VARCHAR(255) NOT NULL,
                                     "token" VARCHAR(255) NOT NULL,
                                     "Title" VARCHAR(255) NOT NULL,
                                     "body" VARCHAR(255) NOT NULL,
                                     "data" VARCHAR(255) NOT NULL,
                                     "status" VARCHAR(255) NOT NULL,
                                     "created_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL,
                                     "updated_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL
);
CREATE TABLE "sms"(
                                     "id" UUID DEFAULT uuid_generate_v4 () NOT NULL,
                                     "password" VARCHAR(255) NOT NULL,
                                     "user" VARCHAR(255) NOT NULL,
                                     "sender_id" VARCHAR(255) NOT NULL,
                                     "api_gate_way" VARCHAR(255) NOT NULL,
                                     "call_back_url" VARCHAR(255) NOT NULL,
                                     "body" VARCHAR(255) NOT NULL,
                                     "receiver_phone" VARCHAR(255) NOT NULL,
                                     "status" VARCHAR(255) NOT NULL,
                                     "created_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL,
                                     "updated_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL
);


INSERT INTO "users"(
    "id",
    "username",
    "password",
    "phone",
    "first_name",
    "middle_name",
    "last_name",
    "email",
    "role_name",
    "created_at",
    "updated_at"

)VALUES(
           '9fe70d3f-bd46-4417-b932-ac69b769464e',
           'superadmin',
           '$2a$12$JNEnCqlxXaMH/rAKGooof.TkQWpfAIcW.FKLBdZ/llhkUR6RQdjKe',
           '0912345678',
           'SUPER',
           'ADMIN',
           'INSPECTION',
           'superadmin@gmail.com',
           'SUPER-ADMIN',
           '2021-06-17 12:16:35',
           '2021-06-17 12:16:35'
       ),(
           '479565db-32f0-41ab-9b0f-2c5a42132857',
    'RIDEADMIN',
    '$2a$12$JNEnCqlxXaMH/rAKGooof.TkQWpfAIcW.FKLBdZ/llhkUR6RQdjKe',
    '0912345454',
    'RIDE',
    'ADMIN',
    'INSPECTION',
    'rideadmin@gmail.com',
    'COMPANY-ADMIN',
    '2021-06-17 12:16:35',
    '2021-06-17 12:16:35'
),(
    'cb6ed203-aa05-43b5-9341-97b6fe830fbc',
    'RIDECLERK',
    '$2a$12$JNEnCqlxXaMH/rAKGooof.TkQWpfAIcW.FKLBdZ/llhkUR6RQdjKe',
    '0932145454',
    'RIDE',
    'CLERK',
    'INSPECTION',
    'rideclerk@gmail.com',
    'COMPANY-CLERK',
    '2021-06-17 12:16:35',
    '2021-06-17 12:16:35'
),
(
    '43df07bc-2383-453a-856e-3c9e45e8f96f',
    'OTHERADMIN',
    '$2a$12$JNEnCqlxXaMH/rAKGooof.TkQWpfAIcW.FKLBdZ/llhkUR6RQdjKe',
    '0912345455',
    'OTHER',
    'ADMIN',
    'INSPECTION',
    'otheradmin@gmail.com',
    'COMPANY-ADMIN',
    '2021-06-17 12:16:35',
    '2021-06-17 12:16:35'
),(
    '2ac0caff-1f50-449b-b7f0-3e2a10584cf5',
    'OTHERCLERK',
    '$2a$12$JNEnCqlxXaMH/rAKGooof.TkQWpfAIcW.FKLBdZ/llhkUR6RQdjKe',
    '0932145456',
    'OTHER',
    'CLERK',
    'INSPECTION',
    'otherclerk@gmail.com',
    'COMPANY-CLERK',
    '2021-06-17 12:16:35',
    '2021-06-17 12:16:35'
),
(
    '791c4939-4848-41b4-9b8e-8a186942d31b',
    'GARAGEADMIN',
    '$2a$12$JNEnCqlxXaMH/rAKGooof.TkQWpfAIcW.FKLBdZ/llhkUR6RQdjKe',
    '0939995454',
    'RIDE',
    'GARAGE',
    'ADMIN',
    'garageadmin@gmail.com',
    'GARAGE-ADMIN',
    '2021-06-17 12:16:35',
    '2021-06-17 12:16:35'
),(
    '791c4939-4848-41b4-9b8e-8a186909d31b',
    'INSPECTOR',
    '$2a$12$JNEnCqlxXaMH/rAKGooof.TkQWpfAIcW.FKLBdZ/llhkUR6RQdjKe',
    '0987654321',
    'RIDE',
    'GARAGE',
    'INSPECTOR',
    'garageinspector@gmail.com',
    'INSPECTOR',
    '2021-06-17 12:16:35',
    '2021-06-17 12:16:35'
),(
    '0a1b0652-5b08-4eb0-9699-c23411657778',
    'OTHERGARAGEADMIN',
    '$2a$12$JNEnCqlxXaMH/rAKGooof.TkQWpfAIcW.FKLBdZ/llhkUR6RQdjKe',
    '0987654322',
    'OTHER',
    'GARAGE',
    'ADMIN',
    'othergarageadmin@gmail.com',
    'GARAGE-ADMIN',
    '2021-06-17 12:16:35',
    '2021-06-17 12:16:35'
),(
    '47482fcb-7fda-430b-ac96-903311ed3288',
    'OTHER-INSPECTOR',
    '$2a$12$JNEnCqlxXaMH/rAKGooof.TkQWpfAIcW.FKLBdZ/llhkUR6RQdjKe',
    '0987654323',
    'OTHER',
    'GARAGE',
    'INSPECTOR',
    'otherinspector@gmail.com',
    'INSPECTOR',
    '2021-06-17 12:16:35',
    '2021-06-17 12:16:35'
);