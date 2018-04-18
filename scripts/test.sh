cd src/main/
for i in *; do 
  echo "Testing $i" 
  cd $i
  for f in *_test.go; do
    go test $f
  done
  cd ..
done
echo "Done."