# Oppgaver til intervju i NRK

Som en forberedelse til intervjuet ber vi deg løse to ulike oppgaver på forhånd. Det er viktig for oss at du både prøver å løse oppgavene og kan forklare for oss hvordan du har tenkt. Veien din til målet er vel så viktig som å nå målet.

### Del 1

Github brukes av NRK til å lagre og versjonskontrollere kildekode i utvikling av applikasjoner. Se for deg at det finnes et Github-repo som bl.a. holder informasjon om brukere.

Lag et script som henter ut informasjon om et Github repo – du kan selv velge kode/scriptspråk. Du kan bruke [nrkno/terraform-registry](https://github.com/nrkno/terraform-registry)
Vi ønsker at du henter ut div info om repoet igjennom Github API.

1. Få med følgende felter fra JSON-strukturen API-et returnerer:
	- `name`: Navn på repo
	- `html_url`: URL til repo
	- `description`: Beskrivelse av repo
	- `license`: lisens brukt på repo
1. Hent ut informasjon om siste commit med nyttig metadata (feks, dato, hvem som commitet el.).
1. Det er en bonus om det som printes i et format som gjør det lett å lese

### Del 2
 1. Bruk koden du lagde over til å hente ut alle public-repos til [nrkno](https://github.com/nrkno)
 1. Pakketer koden din i et container-image
 1. Hvis container-imaget ditt skulle kjørt gjevnlig, hvordan ville du gått frem for å løse det?

#### Tips og ressurser

- Denne oppgaven går ut på å hente ut data fra [Github-API-et](https://docs.github.com/en/rest/guides/getting-started-with-the-rest-api)
- Det finnes egne biblioteker du kan bruke for [Golang](https://github.com/google/go-github) og [Python](https://github.com/PyGithub/PyGithub)
- Docker doc: [docs.docker.com](https://docs.docker.com/reference/dockerfile/)