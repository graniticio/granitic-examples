ID:ARTIST_ID_SELECT

SELECT
    id
FROM
    artist
WHERE
    name = '${artistName}'

ID:ARTIST_INSERT

INSERT INTO artist (
    name
) VALUES (
    '${artistName}'
)


ID:RECORD_INSERT

INSERT INTO record (
    cat_ref,
    name,
    artist_id
) VALUES (
    '${catRef}',
    '${recordName}',
    ${artistID}
)

ID:TRACK_INSERT

INSERT INTO record_track (
    record_id,
    track_number,
    name
) VALUES (
    ${recordId},
    ${trackNumber},
    '${name}'
)

