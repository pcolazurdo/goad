package cli

const template = `[general]
# The base URL to be loadtested if this is set you don't have to pass it on
# the command-line.
;url = http://example.com

# Number of concurrent lambda functions.
concurrency = 10

# Total count of requests to be executed.
requests = 1000

# Total timelimit for the benchmark to run.
;timelimit = 3600

# timeout for individual http requests
;timeout = 15

# Store output in json
;json-output = result.json

# The HTTP method to be used
;method = GET

# The requst body passed with each requests
;body = Hello world

[regions]
# You can specify various aws region to run lambda functions.

us-east-1 ;N.Virginia
;us-east-2 ;Ohio
;us-west-1 ;N.California
;us-west-2 ;Oregon
eu-west-1 ;Ireland
;eu-central-1 ;Frankfurt
;ap-southeast-1 ;Singapore
;ap-southeast-2 ;Tokio
ap-northeast-1 ;Sydney
;ap-northeast-2 ;Seoul
;sa-east-1 ;Sao Paulo

[headers]
# These headers are used in the HTTP request header

;cache-control: no-cache
;auth-token: YOUR-SECRET-AUTH-TOKEN
;base64-header: dGV4dG8gZGUgcHJ1ZWJhIA==
`
