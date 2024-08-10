# Password Steps Recommendation

## Required
1. Docker
2. Make
3. Golang 1.21

## Getting started
Start docker compose (webserver, postgres, nginx)
  ```bash
  make compose
  ```
Web server will run in port `3000`

You can try to run below command in terminal to see result

  ```bash
  curl --location 'http://localhost:3000/api/strong_password_steps' \
  --header 'Content-Type: application/json' \
  --data '{
      "init_password": "aA1"
  }'
  ```

## Unit Test
Run in command line
  ```bash
  make test
  ```
