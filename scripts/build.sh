#!/usr/bin/env bash
dep ensure 
echo "Compiling functions to bin/handlers/ ..."

rm -rf bin/

cd src/main/
for i in *; do 
  cd $i
  for f in *.go; do
    
    if ! [[ "${f}" == *test* ]]; then   
      if GOOS=linux go build -o "../../../bin/handlers/$i" ${f}; then
        echo "✓ Compiled $i"
      else
        echo "✕ Failed to compile $filename!"
        exit 1
      fi  
    fi   
  done
  cd ..
done
echo "Done."
