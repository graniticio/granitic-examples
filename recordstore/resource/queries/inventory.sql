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
