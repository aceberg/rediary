[![Main-Docker](https://github.com/aceberg/rediary/actions/workflows/main-docker.yml/badge.svg)](https://github.com/aceberg/rediary/actions/workflows/main-docker.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/aceberg/rediary)](https://goreportcard.com/report/github.com/aceberg/rediary)
[![Maintainability](https://api.codeclimate.com/v1/badges/e8f67994120fc7936aeb/maintainability)](https://codeclimate.com/github/aceberg/rediary/maintainability)
![Docker Image Size (latest semver)](https://img.shields.io/docker/image-size/aceberg/rediary)

<h1><a href="https://github.com/aceberg/rediary">
    <img src="https://raw.githubusercontent.com/aceberg/rediary/main/assets/logo.png" width="35" />
</a>Resource Diary</h1>

Emotional resource diary

- [Quick start](https://github.com/aceberg/rediary#quick-start)
- [Usage](https://github.com/aceberg/rediary#usage)
- [Config](https://github.com/aceberg/rediary#config)
- [Options](https://github.com/aceberg/rediary#options)
- [Thanks](https://github.com/aceberg/rediary#thanks)


![Screenshot](https://raw.githubusercontent.com/aceberg/rediary/main/assets/Screenshot%202023-04-02%20at%2022-27-40%20Resource%20Diary.png)

## Quick start

```sh
docker run --name rediary \
-e "TZ=Asia/Novosibirsk" \
-v ~/.dockerdata/rediary:/data/rediary \
-p 8847:8847 \
aceberg/rediary
```
Or use [docker-compose.yml](docker-compose.yml)

## Usage
Add your own Tags and Actions on Config page.

## Config


Configuration can be done through config file or environment variables

| Variable  | Description | Default |
| --------  | ----------- | ------- |
| DB        | Path to Database | /data/rediary/sqlite.db |
| HOST | Listen address | 0.0.0.0 |
| PORT   | Port for web GUI | 8847 |
| THEME | Any theme name from https://bootswatch.com in lowcase | minty |
| BGCOLOR | Background color: light or dark | light |
| TZ | Set your timezone for correct time | "" |
| AUTH | Enable Session-Cookie authentication | false |
| AUTH_USER | Username | "" |
| AUTH_PASSWORD | Encrypted password (bcrypt). [How to encrypt password with bcrypt?](docs/BCRYPT.md) | "" |
| AUTH_EXPIRE | Session expiration time. A number and suffix: **m, h, d** or **M**. | 7d |

## Options

| Key  | Description | Default | 
| --------  | ----------- | ------- | 
| -c | Path to config file | /data/rediary/config.yaml |  

## Thanks
- All go packages listed in [dependencies](https://github.com/aceberg/rediary/network/dependencies)
- [Bootstrap](https://getbootstrap.com/)
- Themes: [Free themes for Bootstrap](https://bootswatch.com)
- Favicon and logo: [Flaticon](https://www.flaticon.com/icons/)