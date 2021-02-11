app:
  env: "{{ ENV }}"
  key: "{{ APP_KEY }}"
  port: "{{ PORT }}"

databases:
  mongodb:
    host: "{{ MONGODB_HOST }}"
    database: "{{ MONGODB_DATABASE }}"
    uri: "{{ MONGODB_CONNECTION_STRING }}"
    authentication:
      username: "{{ MONGODB_USERNAME }}"
      password: "{{ MONGODB_PASSWORD }}"
      database: "{{ MONGODB_DATABASE_AUTH }}"

logger:
  sentry_dsn: "{{ SENTRY_DSN }}"

