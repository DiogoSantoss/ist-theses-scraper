package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func newTestServer() *httptest.Server {

	mux := http.NewServeMux()

	mux.HandleFunc("/dissertacoes", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte(`<!DOCTYPE html>
<html>
<head>
</head>
<body>
<ul>
<li>
<div year="2443836393482" state="EVALUATED" degree="2761663971475">
<div class="thesis-title">
<h5>
<a href="dissertacao/2353642462449">Syntactic REAP.PT: Exercises on Word Formation</a>
<span class="label label-success">EVALUATED</span>
</h5>
<h5>
<small>Author: Pedro Henrique Santos Figueirinha (ist158123)</small>
</h5>
<h5>
<small>
" Coordenação: "
<span class="orientator">Jorge Manuel Evangelista Baptista (ist406262)</span>
<span class="orientator">Nuno João Neves Mamede (ist12099)</span>
</small>
</h5>
</div>
<hr>
</div>
</li>
<li>
<div year="846868766523393" state="DRAFT" degree="2761663971475" style="display: block;">
<div class="thesis-title">
<h5>
<a href="dissertacao/1128253548923306">(Projeto) Análise de dados do CERT.PT</a>
<span class="label label-default">DRAFT</span>
</h5>
<h5>
<small>Author: Clara Gil da Silva Pereira (ist197070)</small>
</h5>
<h5>
<small>
" Coordenação: "
<span class="orientator">José Luís Brinquete Borbinha (ist13085)</span>
</small>
</h5>
</div>
<hr>
</div>
</li>
</ul>
</body>
</html>
		`))
	})

	return httptest.NewServer(mux)
}

func TestTwoTheses(t *testing.T) {
	server := newTestServer()
	defer server.Close()

	theses := []thesis{}
	domain := server.URL[7:]
	baseURL := server.URL
	visitURL := server.URL + "/dissertacoes"

	c := createCollyCollector(&theses, domain, baseURL)
	c.Visit(visitURL)

	printTheses(theses)

	if len(theses) != 2 {
		t.Errorf("Expected 1 thesis, got %d", len(theses))
	}

	if theses[0].Title != "Syntactic REAP.PT: Exercises on Word Formation" {
		t.Errorf("Expected title to be 'Syntactic REAP.PT: Exercises on Word Formation', got '%s'", theses[0].Title)
	}

	if theses[1].Title != "(Projeto) Análise de dados do CERT.PT" {
		t.Errorf("Expected title to be '(Projeto) Análise de dados do CERT.PT', got '%s'", theses[1].Title)
	}

	if theses[0].Link != (baseURL + "dissertacao/2353642462449") {
		t.Errorf("Expected link to be 'dissertacao/2353642462449', got '%s'", theses[0].Link)
	}

	if theses[1].Link != (baseURL + "dissertacao/1128253548923306") {
		t.Errorf("Expected link to be 'dissertacao/1128253548923306', got '%s'", theses[1].Link)
	}

	if theses[0].Author != "Pedro Henrique Santos Figueirinha (ist158123)" {
		t.Errorf("Expected author to be 'Pedro Henrique Santos Figueirinha (ist158123)', got '%s'", theses[0].Author)
	}

	if theses[1].Author != "Clara Gil da Silva Pereira (ist197070)" {
		t.Errorf("Expected author to be 'Clara Gil da Silva Pereira (ist197070)', got '%s'", theses[1].Author)
	}

	if len(theses[0].Advisors) != 2 {
		t.Errorf("Expected 2 advisors, got %d", len(theses[0].Advisors))
	}

	if len(theses[1].Advisors) != 1 {
		t.Errorf("Expected 1 advisor, got %d", len(theses[1].Advisors))
	}

	if theses[0].Advisors[0] != "Jorge Manuel Evangelista Baptista (ist406262)" {
		t.Errorf("Expected advisor to be 'Jorge Manuel Evangelista Baptista (ist406262)', got '%s'", theses[0].Advisors[0])
	}

	if theses[0].Advisors[1] != "Nuno João Neves Mamede (ist12099)" {
		t.Errorf("Expected advisor to be 'Nuno João Neves Mamede (ist12099)', got '%s'", theses[0].Advisors[1])
	}

	if theses[1].Advisors[0] != "José Luís Brinquete Borbinha (ist13085)" {
		t.Errorf("Expected advisor to be 'José Luís Brinquete Borbinha (ist13085)', got '%s'", theses[1].Advisors[0])
	}
}
