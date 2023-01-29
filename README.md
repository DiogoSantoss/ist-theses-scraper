# IST Theses Scraper

A web scraper tool developed to gather thesis information from the Fenix IST website. This tool enables users to search for theses by advisor, providing a convenient and efficient way to access this information.

## Usage

```bash
go build .
./scraper <Optional Advisor Name>
```

## Example

```
$ go build .
$ ./scraper Ricardo Chaves
Visiting https://fenix.tecnico.ulisboa.pt/cursos/meic-a/dissertacoes
Success

Title: Automated Smart Fuzzing Vulnerability Testing
Link: https://fenix.tecnico.ulisboa.pt/cursos/meic-a/dissertacao/846778572213879
Author: João Afonso Lopes de Almeida Pinto Coutinho (ist189470)
Advisor: Ricardo Jorge Fernandes Chaves (ist143817)

Title: Projecto: Kruptos 2: Cryptogaphic gateway for managing cellular network keys
Link: https://fenix.tecnico.ulisboa.pt/cursos/meic-a/dissertacao/565303595503596
Author: Afonso Silvano Paredes (ist189401)
Advisor: Ricardo Jorge Fernandes Chaves (ist143817)

[...]
```

## Contributing

Pull requests are welcome. For major changes, please open an issue first
to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License

[MIT](https://choosealicense.com/licenses/mit/)