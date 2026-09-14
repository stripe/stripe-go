This release **doesn't** change the pinned API version; it still uses `2026-05-27.dahlia`.

We're doing an out-of-band-major to update a field type that changed. If you're not using `tax_details`, this is a no-op release when compared with the last one. If you _are_ using `tax_details` its type has changed slightly and you'll have to update your code when upgrading.
