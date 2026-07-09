# gopt

When we are working together in a team. We share code by cloning the repository. But, we have to share the .env files manually through Whatsapp for something. I wondered if there is a way in which we can commit the .env files safely without the risk of leakage. And that's how gopt was born. Gopt is gofication of _gupt_ (hindi word to hidden). While is is also the short form of go encrypt. 

I haven't yet developed this but this is how it is supposed to work. 

```bash
gopt encrypt .env .env.crypt

```

This will encrpt `.env` file into `.env.crypt`, (Maybe it will take a password but later we will move on to using ssh). This can be commited and when someone takes a pull, this can be decryped by the same password. `.env` will remain in the gitingore so there is no risk of accidently committin it. 

To decrypt use 

```bash
go decrypt .env.crypt

```

This will generate or overwrite the `.env`. 

## Contributing

Contribution are welcome but adhere to code of conduct. AI written code is also welcome. Always remember:

- Only PR with small diffs will be reviwed more smaller PR will be better. Anything greater change an 200 lines will be deleted. 
- If it fails the tests, then automatically rejected. 
- Write clear description of changes, if not then automatically rejected. 
- You can use AI to write PRs, but please don't submit huge changes. 
