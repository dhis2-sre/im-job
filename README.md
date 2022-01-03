# Instance Jobs


# Token valid by key found in jwks.json

```sh
export ACCESS_TOKEN=eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjQ3OTYzODQwNzcsImlhdCI6MTY0MDY5MTQ3NywidXNlciI6eyJJRCI6NCwiQ3JlYXRlZEF0IjoiMjAyMS0xMi0yOFQxMDo0NjoxNi44NTEzMzlaIiwiVXBkYXRlZEF0IjoiMjAyMS0xMi0yOFQxMDo0NjoxNi44NTEzMzlaIiwiRGVsZXRlZEF0IjpudWxsLCJFbWFpbCI6InNvbWVvbmVAc29tZXRoaW5nLm9yZyIsIkdyb3VwcyI6bnVsbCwiQWRtaW5Hcm91cHMiOm51bGx9fQ.FnPIu36kV1T-Jix5Wy-HsZeqxQI6Q_7HQ14C1DWKHETIBSk-vLQ_sCMHVPKA42utEDFI3Xpmf6Gyzv9aPU_Cvg-JDazRprfrZBqn4LzSmT6K3HGoKoQ0b5G8exxz0Ote8NQDB1NBZmYvD1gpVVisCvzaewJTRAvRA3DS0n_O4kU5QENdLNPfWFo0rXOC83sLBsEIe2Ce4TiRrepOCSQKE-_rwQQSA3w30MhFmhAY7Ozcd9i69mtfcvqjORdNJ-zREgiw8B2g9oh7byE1h2oxjvoKC3WRfPeSYoRY6GuMHSSWJdzFKIswlZHdWU1GicPJASBbkKGbP5n5O6FXyeo0bw
```

# Backup... save, saveas
Is backup just another job?
Create container with collection of bash scripts
Create secret with input variables
Launch job with above container and command pointing to a specific script and env from secret

backup... pg_dump, aws s3 cp, eod
input: db host/user/pass/etc, backup name (s3 key: group/whatever.pgc)

create im-jobs service?
post /jobs      - launch new job
    payload: script name + key/value for environment variables
get /jobs/:id   - get job status
