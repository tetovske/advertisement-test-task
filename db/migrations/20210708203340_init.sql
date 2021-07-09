-- +goose Up
-- +goose StatementBegin
CREATE TABLE ADVERTISEMENTS(
    Id          SERIAL PRIMARY KEY,
    Title       VARCHAR(200) NOT NULL,
    Description VARCHAR(1000),
    Price       INTEGER NOT NULL
);

CREATE TABLE PHOTOS(
    Id                  SERIAL PRIMARY KEY,
    Link                VARCHAR(70) NOT NULL,
    Tag                 INTEGER NOT NULL DEFAULT 0,
    AdvertisementId     INTEGER,
    FOREIGN KEY (AdvertisementId) REFERENCES ADVERTISEMENTS (Id) ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE PHOTOS;

DROP TABLE ADVERTISEMENTS;
-- +goose StatementEnd
