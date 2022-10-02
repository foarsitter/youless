docker run --rm \
  -v "${PWD}"/client:/output \
  -v "${PWD}"/generator:/local openapitools/openapi-generator-cli generate \
   --git-repo-id youless \
   --git-user-id foarsitter \
  -i /local/Jelmert-Engineering-Youless-0.1.5-swagger.json \
  -g go \
  -o /output \
  --additional-properties=packageName=youless \
  --additional-properties=packageVersion="${VERSION}"