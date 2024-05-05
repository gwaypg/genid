# Instruction
A distributed id node 

# Build
```
git clone https://github.com/gwaypg/genid

cd genid
. env.sh
sup build all

Or

. env.sh
cd cmd/snowflake
go build # or sup build
```

# Run
```
cd cmd/snowflake
./snowflake
```

# Deployment

## origin
```
PRJ_ROOT=<genid root path> ./$PRJ_ROOT/cmd/snowflake/snowflake
```

## supd
```
git clone https://github.com/gwaypg/supd
cd supd
. env.sh
cd cmd/supd
./publish.sh
./setup.sh install # uninstall: ./setup.sh clean 

cd genid
. env.sh
cd cmd/snowflake
sup install
sup status
```

TODO: more deployments

