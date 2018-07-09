

solc IdentityManager.sol --abi --bin -o build

cd build

abigen --abi IdentityManager.abi --pkg bindigs --type IdentityManager --out IdentityManager.go --bin IdentityManager.bin

rm -Rf ../bindings/IdentityManager.go

mv IdentityManager.go ../bindings/IdentityManager.go

cd ..
rm -Rf build