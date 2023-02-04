TARGET_EXEC := regen_abi
ABIGEN_PATH := ./abigen

$(TARGET_EXEC):
	solc SapoBridge.sol --bin --abi -o temp
	./abigen --abi temp/SapoBridge.abi --bin temp/SapoBridge.bin --pkg contracts --out SapoBridge.go
	./abigen --abi temp/SapoJob.abi --bin temp/SapoJob.bin --pkg jobs --out SapoJob.go
	rm -rf temp/
	sed -i '4s/.*/package contracts/' SapoJob.go