# This is your quotes starter!

The canvas on which to paint your art

# Short term urgent needs
- Implement a go HTTP server using git

- It should listen on /quotes and respond to GET with a single random quote

- Those quotes should be from the following sites. Add a minimum of 5 and a maximum of 20.[Go Proverbs](https://go-proverbs.github.io) [Go Idioms](https://dmitri.shuralyov.com/idiomatic-go)

- The json response from your API should match this syntax.
```json
{
  "quote": "YOUR RANDOM QUOTE HERE",
  "author": "Jessie Gorrell"
}
```

- Your random quotes can be in memory, hard-coded, pulled in from a file, or in a database. We suggest starting with hard-coded to keep things simple.

- Ensure you are authed against GCP
```shell
gcloud auth login
```

- Ensure you have set your active project
```shell
gcloud config set project name-apprentice
```

- Deploy your service to cloud run by running and following the prompts. Name your service FIRSTNAME-service
```shell
gcloud run deploy
```
