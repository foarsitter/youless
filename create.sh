docker run --rm \
  -v "${PWD}"/client:/output \
  -v "${PWD}"/generator:/local openapitools/openapi-generator-cli generate \
   --git-repo-id youless \
   --git-user-id foarsitter \
  -i /local/openapi.json \
  -g go \
  -o /output \
  --additional-properties=packageName=youless \
  --additional-properties=packageVersion="${VERSION}"