# Go Snaptron API (gosnaptron)

leverage go to handle snaptron

current version: 1.0



## Features

1) Build Basic Queries. Leverage language parallelism when handling Snaptron.

2) Assembly and Create Intermediate Queries. Put together different ops funcs (like LEGO pieces), use the prebuilt summary statistics such as shared sample count, junction inclusion ratio, etc., or build your own.



## How to run API

1) Install Go 1.9+
2) Clone Repo
2) Create your own basic or intermediate queries. Refer to src/simulator.go to see examples.



## etc

Go Snaptron API: ©2017, Charlie Wang

Snaptron is Christopher Wilks, Phani Gaddipati, Abhinav Nellore, Ben Langmead; Snaptron: querying splicing patterns across tens of thousands of RNA-seq samples, Bioinformatics, (2017), btx547, https://doi.org/10.1093/bioinformatics/btx547