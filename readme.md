# resistance

A game of deceit

## Setup

```sh
$ git clone https://github.com/azdavis/resistance.git
$ cd resistance/server
$ go get
$ cd ../client
$ npm install
```

## Development

In one terminal:

```sh
$ cd server
$ go build
$ ./resistance
```

In another terminal:

```sh
$ cd client
$ npm run start
```

## Deploy

See [the deploy script][1] and [some setup instructions][2].

[1]: https://github.com/azdavis/azdavis.xyz/blob/master/deploy.sh
[2]: ./deploy-setup.md
