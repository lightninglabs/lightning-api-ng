# gcloud bucket objects metadata needs to be updated after every deployment

# never cache HTML pages
gsutil -m setmeta -r -h "Cache-Control: public, max-age=0, must-revalidate" gs://${GC_BUCKET}/${GC_PATH}**/*.html

# cache files in the assets dir forever
gsutil -m setmeta -r -h "cache-control: public, max-age=31536000, immutable" gs://${GC_BUCKET}/${GC_PATH}assets/**

# cache the search index forever
gsutil -m setmeta -r -h "cache-control: public, max-age=31536000, immutable" gs://${GC_BUCKET}/${GC_PATH}search-index.json
