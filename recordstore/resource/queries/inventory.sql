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

ID:CAT_REF_SELECT

SELECT
    id
FROM
    record
WHERE
    cat_ref = '${catRef}'


ID:ARTIST_SEARCH_BASE

SELECT
    id AS ID,
    name as Name
FROM
    artist


ID:ARTIST_DETAIL

SELECT
    name
FROM
    artist
WHERE
    id = ${id}




