### revor
###### Rest service for asterisk cdrs

###### Please install git mercurial package into your system

For exampel for debian based : apt-get install git mercurial

######
Add GOROOT and GOPATH enverenement variabls to .bachrc as follow :

export GOROOT=$HOME/Projects/golang/go

export GOPATH=$HOME/Projects/golang/mygo

export PATH=$PATH:$GOROOT/bin:$GOPATH/bin

=
Please install 
go get github.com/robfig/config

go get -u github.com/howeyc/fsnotify

go get github.com/revel/cmd/revel



=
Execute revel run revor into your host system and see the result on the port 9001

