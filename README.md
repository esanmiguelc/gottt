# Go TTT

## Demo
Visit: [52.0.207.158:8080](http://52.0.207.158:8080)
    
## Setting Up
*Note: (follow all of the steps instead of using `git clone` )*
### Make sure you have go installed

If you are using a mac:

    brew install go 
    
If you are using linux:

    yum install go
    
### Preparing your environment

Have a folder where you save all your go projects ( if you don't already )

    mkdir gocode
    
Export your go path to this directory

    export GOPATH=$HOME/gocode

For convenience also add the bin directory to your path

    export PATH=$PATH:$GOPATH/bin
    
Add the following to run the project

    go get github.com/stretchr/testify/assert
    go get github.com/julienschmidt/httprouter
    go get github.com/esanmiguelc/gottt
    
*Note: You will probably be prompted for your login information this should be fine as long as you have access to this repository*
    
    
## To run the tests
    
    cd src/github.com/esanmiguelc/gottt
    go test ./...
    
## To run this project

    cd src/github.com/esanmiguelc/gottt
    PORT=8080 go run web/main.go
    
