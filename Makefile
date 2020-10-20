all:
	packr build -ldflags="-w -s -H windowsgui" -o google-tran.exe
	#upx -9 google-tran.exe