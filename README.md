# gosnaptron SDK

leverage Go to handle mRNA data

current version: 1.0



## Features

1) Basic Level. Build Basic Queries. Leverage language parallelism when handling mRNA data from Snaptron server.

2) Intermediate Level. Assemble and Create Intermediate Queries. Put together different dplyr-like functions: Union, Intersect, Bind, Summarize, Filter, etc.

3) High Level. Summary statistics like Shared Sample Count, Junction Inclusion Ratio, etc.



## How to run gosnaptron

1) Install Go 1.9+. Set up one's GOPATH.
2) Clone repo
3) Navigate to src/
4) Run "go run playground.go". See the default Shared Sample Count run.
5) Go to func main() in src/playground.go and choose another example to run by uncommenting. 
6) Profit :)
7) Create your own by following the examples in src/playground.go.



## etc

gosnaptron SDK: ©2017, Charlie Wang

Snaptron is Christopher Wilks, Phani Gaddipati, Abhinav Nellore, Ben Langmead; Snaptron: querying splicing patterns across tens of thousands of RNA-seq samples, Bioinformatics, (2017), btx547, https://doi.org/10.1093/bioinformatics/btx547