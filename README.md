# services.email

## I need to make this a script somehow...

make build-handler-email-post
make zip-handler-email-post
dockerup

This updates the lambda if changes are made and the Make file is used to rebuild and zip handler.

```bash
awslocal lambda update-function-code \
 --function-name services-email-email \
 --zip-file fileb://handler_email_post.zip
```
