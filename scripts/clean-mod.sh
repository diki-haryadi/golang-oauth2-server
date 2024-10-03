go clean -modcache

# 2. Update go.mod to use the newer version
go get google.golang.org/genproto/googleapis/rpc@latest

# 3. Update all dependencies and cleanup unused ones
go mod tidy

# 4. If you're still having issues, try explicitly removing the old version
go get google.golang.org/genproto@none

# 5. Run tidy again to ensure everything is clean
go mod tidy
