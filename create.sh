docker run --rm \
  -v ${PWD}/../youless:/output \
  -v ${PWD}:/local openapitools/openapi-generator-cli generate \
   --git-repo-id youless \
   --git-user-id foarsitter \
  -i /local/Jelmert-Engineering-Youless-0.1.3-swagger.json \
  -g go \
  -o /output \
  --additional-properties=packageName=youless \
  --additional-properties=packageVersion=0.1.3